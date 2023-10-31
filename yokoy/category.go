package yokoy

import (
	"context"
	"database/sql"
	"errors"

	"github.com/salesfive/yokoy/api"
	"github.com/salesfive/yokoy/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func fetchCategories(ctx context.Context, legalEntityId string, c api.ClientWithResponsesInterface) (*[]api.Category, error) {
	p := &api.GetLegalEntitiesLegalEntityIdCategoriesParams{}

	h, err := c.GetLegalEntitiesLegalEntityIdCategoriesWithResponse(ctx, legalEntityId, p)

	if err != nil {
		return nil, err
	}

	if h.StatusCode() != 200 {
		return nil, errors.New("got non 200 response code")
	}

	return h.JSON200.Categories, nil

}

func truncateCategories(ctx context.Context, db *sql.DB) {
	db.ExecContext(ctx, "truncate categories;")
}

func insertCategories(ctx context.Context, db *sql.DB, legalEntityID string, categories *[]api.Category) error {
	for _, c := range *categories {
		r := models.Category{}
		r.ID = *c.Id
		r.AccountReference = null.StringFrom(c.AccountReference)
		r.ChargeToEmployee = null.BoolFromPtr(c.ChargeToEmployee)
		r.Name = null.StringFrom(c.Name)
		r.StatusActive = null.BoolFrom(c.StatusActive)
		r.LegalEntityID = null.StringFrom(legalEntityID)

		err := r.Insert(ctx, db, boil.Infer())
		if err != nil {
			return err
		}
	}

	return nil
}
