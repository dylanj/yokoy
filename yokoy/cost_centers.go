package yokoy

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"strconv"

	"github.com/dylanj/yokoy/api"
	"github.com/shopspring/decimal"
	"github.com/aarondl/null/v9"
)

func fetchCostCenters(ctx context.Context, legalEntityId string, c api.ClientWithResponsesInterface) (*[]api.CostCenter, error) {
	// Use underlying client to get raw response and parse manually
	// This handles approvalLimit which can be string or number
	clientWithResponses, ok := c.(*api.ClientWithResponses)
	if !ok {
		return nil, errors.New("cannot access underlying client for raw request")
	}

	p := &api.ListCostObjectsParams{}
	httpResp, err := clientWithResponses.ListCostObjects(ctx, api.LegalEntityIdInPath(legalEntityId), p)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != 200 {
		return nil, errors.New("got non 200 response code")
	}

	bodyBytes, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}

	// Parse response manually to handle approvalLimit as string or number
	var responseData struct {
		CostCenters []struct {
			AccessType         *string                `json:"accessType"`
			ApprovalLimit      interface{}            `json:"approvalLimit"` // Can be string or float32
			ApproverId         *string                `json:"approverId"`
			AutoApprovalLimit  *float32               `json:"autoApprovalLimit"`
			Code               string                 `json:"code"`
			CustomInformation  *map[string]string     `json:"customInformation"`
			DelegateExpiry     *string                `json:"delegateExpiry"`
			DelegateId         *string                `json:"delegateId"`
			Description        *string                `json:"description"`
			Id                 *string                `json:"id"`
			Name               string                 `json:"name"`
			ParentId           *string                `json:"parentId"`
			Scopes             *[]string              `json:"scopes"`
			StatusActive       *bool                  `json:"statusActive"`
		} `json:"costCenters"`
	}

	if err := json.Unmarshal(bodyBytes, &responseData); err != nil {
		return nil, err
	}

	// Convert to api.CostCenter, handling approvalLimit type conversion
	costCenters := make([]api.CostCenter, 0, len(responseData.CostCenters))
	for _, cc := range responseData.CostCenters {
		var approvalLimit *float32
		if cc.ApprovalLimit != nil {
			switch v := cc.ApprovalLimit.(type) {
			case float32:
				if !math.IsNaN(float64(v)) {
					approvalLimit = &v
				}
			case float64:
				if !math.IsNaN(v) {
					f32 := float32(v)
					approvalLimit = &f32
				}
			case string:
				// Handle "NaN" string
				if v != "NaN" && v != "nan" {
					if f, err := strconv.ParseFloat(v, 32); err == nil && !math.IsNaN(f) {
						f32 := float32(f)
						approvalLimit = &f32
					}
				}
			case int:
				f32 := float32(v)
				approvalLimit = &f32
			case int64:
				f32 := float32(v)
				approvalLimit = &f32
			}
		}

		// Handle Scopes conversion
		var scopes *[]api.CostCenterScopes
		if cc.Scopes != nil {
			scopesList := make([]api.CostCenterScopes, 0, len(*cc.Scopes))
			for _, s := range *cc.Scopes {
				scopesList = append(scopesList, api.CostCenterScopes(s))
			}
			scopes = &scopesList
		}

		costCenter := api.CostCenter{
			AccessType:        (*api.CostCenterAccessType)(cc.AccessType),
			ApprovalLimit:      approvalLimit,
			ApproverId:         cc.ApproverId,
			AutoApprovalLimit:  cc.AutoApprovalLimit,
			Code:               cc.Code,
			CustomInformation: cc.CustomInformation,
			DelegateExpiry:     cc.DelegateExpiry,
			DelegateId:         cc.DelegateId,
			Description:        cc.Description,
			Id:                 cc.Id,
			Name:               cc.Name,
			ParentId:           cc.ParentId,
			Scopes:             scopes,
			StatusActive:       cc.StatusActive,
		}

		costCenters = append(costCenters, costCenter)
	}

	return &costCenters, nil
}

func truncateCostCenters(ctx context.Context, db *sql.DB) {
	db.ExecContext(ctx, "truncate cost_centers;")
}

func getInt(v interface{}) (int, error) {
	switch v := v.(type) {
	case float64:
		return int(v), nil
	case string:
		c, err := strconv.Atoi(v)
		if err != nil {
			return 0, err
		}
		return c, nil
	default:
		return 0, fmt.Errorf("conversion to int from %T not supported", v)
	}
}

func insertCostCenters(ctx context.Context, db *sql.DB, legalEntityID string, costCenters *[]api.CostCenter) error {
	if len(*costCenters) == 0 {
		return nil
	}

	columns := []string{"id", "name", "approval_limit", "approver_id", "auto_approval_limit", "code", "delegate_expiry", "delegate_id", "description", "parent_id", "status_active", "legal_entity_id"}
	values := make([][]interface{}, 0, len(*costCenters))

	for _, cc := range *costCenters {
		var approvalLimit decimal.Decimal
		if cc.ApprovalLimit != nil {
			approvalLimit = NullDecimalFromFloat32(*cc.ApprovalLimit)
		}

		var autoApprovalLimit decimal.Decimal
		if cc.AutoApprovalLimit != nil {
			// Handle NaN values (as mentioned in README)
			if !math.IsNaN(float64(*cc.AutoApprovalLimit)) {
				autoApprovalLimit = NullDecimalFromFloat32(*cc.AutoApprovalLimit)
			}
			// If NaN, autoApprovalLimit remains zero value
		}

		row := []interface{}{
			*cc.Id,
			NullStringValue(null.StringFrom(cc.Name)),
			approvalLimit,
			NullStringValue(null.StringFromPtr(cc.ApproverId)),
			autoApprovalLimit, // Will be zero value if nil
			NullStringValue(null.StringFrom(cc.Code)),
			NullTimeValue(null.TimeFromPtr(parseTime(cc.DelegateExpiry))),
			NullStringValue(null.StringFromPtr(cc.DelegateId)),
			NullStringValue(null.StringFromPtr(cc.Description)),
			NullStringValue(null.StringFromPtr(cc.ParentId)),
			NullBoolValue(null.BoolFromPtr(cc.StatusActive)),
			legalEntityID,
		}
		values = append(values, row)
	}

	return BulkInsert(ctx, db, "cost_centers", columns, values)
}

func NullDecimalFromFloat32(v float32) decimal.Decimal {
	return NullDecimalFromFloat64(float64(v))
}

func NullDecimalFromFloat64(v float64) decimal.Decimal {
	b := decimal.NewFromFloat(v)
	return b
}

func NullDecimalFromInt64(v int64) decimal.Decimal {
	b := decimal.NewFromInt(v)
	return b
}

func NullDecimalFromInt32(v int32) decimal.Decimal {
	return NullDecimalFromInt64(int64(v))
}
