package yokoy

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"github.com/dylanj/yokoy/api"
	"github.com/dylanj/yokoy/models"
	"github.com/shopspring/decimal"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
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
	for _, cc := range *costCenters {
		r := models.CostCenter{}
		r.ID = *cc.Id
		r.Name = null.StringFrom(cc.Name)

		approvalLimit, err := getInt(cc.ApprovalLimit)
		if err == nil {
			r.ApprovalLimit = NullDecimalFromInt64(int64(approvalLimit))
		}

		r.ApproverID = null.StringFromPtr(cc.ApproverId)
		if cc.AutoApprovalLimit != nil {
			r.AutoApprovalLimit = NullDecimalFromFloat32(*cc.AutoApprovalLimit)
		}

		r.Code = null.StringFrom(cc.Code)
		r.DelegateExpiry = null.TimeFromPtr(parseTime(cc.DelegateExpiry))
		r.DelegateID = null.StringFromPtr(cc.DelegateId)
		r.Description = null.StringFrom(cc.Description)
		r.ParentID = null.StringFromPtr(cc.ParentId)
		r.StatusActive = null.BoolFrom(cc.StatusActive)
		r.LegalEntityID = null.StringFrom(legalEntityID)

		err = r.Insert(ctx, db, boil.Infer())
		if err != nil {
			return err
		}
	}

	return nil
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
