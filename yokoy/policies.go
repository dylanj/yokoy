package yokoy

import (
	"context"
	"database/sql"
	"errors"

	"github.com/dylanj/yokoy/api"
	"github.com/dylanj/yokoy/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
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
	for _, p := range *policies {
		r := models.Policy{}
		r.ID = *p.Id
		r.Name = null.StringFrom(p.Name)
		r.Code = null.StringFrom(p.Code)
		r.StatusActive = null.BoolFrom(p.StatusActive)
		r.LegalEntityID = null.StringFrom(legalEntityID)

		err := r.Insert(ctx, db, boil.Infer())
		if err != nil {
			return err
		}
	}

	return nil
}
