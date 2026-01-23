package yokoy

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"github.com/dylanj/yokoy/api"
	"github.com/shopspring/decimal"
	"github.com/aarondl/null/v9"
)

func fetchCostCenters(ctx context.Context, legalEntityId string, c api.ClientWithResponsesInterface) (*[]api.CostCenter, error) {
	p := &api.GetLegalEntitiesLegalEntityIdCostCentersParams{}

	h, err := c.GetLegalEntitiesLegalEntityIdCostCentersWithResponse(ctx, legalEntityId, p)

	if err != nil {
		return nil, err
	}

	if h.StatusCode() != 200 {
		return nil, errors.New("got non 200 response code")
	}

	return h.JSON200.CostCenters, nil
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
		approvalLimitInt, err := getInt(cc.ApprovalLimit)
		if err == nil {
			approvalLimit = NullDecimalFromInt64(int64(approvalLimitInt))
		}

		var autoApprovalLimit decimal.Decimal
		if cc.AutoApprovalLimit != nil {
			autoApprovalLimit = NullDecimalFromFloat32(*cc.AutoApprovalLimit)
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
			NullStringValue(null.StringFrom(cc.Description)),
			NullStringValue(null.StringFromPtr(cc.ParentId)),
			NullBoolValue(null.BoolFrom(cc.StatusActive)),
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
