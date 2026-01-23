package yokoy

import (
	"context"
	"database/sql"
	"errors"

	"github.com/dylanj/yokoy/api"
	"github.com/aarondl/null/v9"
)

func fetchPolicies(ctx context.Context, legalEntityId string, c api.ClientWithResponsesInterface) (*[]api.Policy, error) {
	p := &api.GetLegalEntitiesLegalEntityIdPoliciesParams{}

	h, err := c.GetLegalEntitiesLegalEntityIdPoliciesWithResponse(ctx, legalEntityId, p)

	if err != nil {
		return nil, err
	}

	if h.StatusCode() != 200 {
		return nil, errors.New("got non 200 response code")
	}

	return h.JSON200.Policies, nil

}

func truncatePolicies(ctx context.Context, db *sql.DB) {
	db.ExecContext(ctx, "truncate policies;")
}

func insertPolicies(ctx context.Context, db *sql.DB, legalEntityID string, policies *[]api.Policy) error {
	if len(*policies) == 0 {
		return nil
	}

	columns := []string{"id", "name", "code", "status_active", "legal_entity_id"}
	values := make([][]interface{}, 0, len(*policies))

	for _, p := range *policies {
		row := []interface{}{
			*p.Id,
			NullStringValue(null.StringFrom(p.Name)),
			NullStringValue(null.StringFrom(p.Code)),
			NullBoolValue(null.BoolFrom(p.StatusActive)),
			legalEntityID,
		}
		values = append(values, row)
	}

	return BulkInsert(ctx, db, "policies", columns, values)
}
