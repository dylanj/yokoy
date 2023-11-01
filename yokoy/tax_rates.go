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

func fetchTaxRates(ctx context.Context, legalEntityId string, c api.ClientWithResponsesInterface) (*[]api.TaxRate, error) {
	p := &api.GetLegalEntitiesLegalEntityIdTaxRatesParams{}

	h, err := c.GetLegalEntitiesLegalEntityIdTaxRatesWithResponse(ctx, legalEntityId, p)

	if err != nil {
		return nil, err
	}

	if h.StatusCode() != 200 {
		return nil, errors.New("got non 200 response code")
	}

	return h.JSON200.TaxRates, nil

}

func truncateTaxRates(ctx context.Context, db *sql.DB) {
	db.ExecContext(ctx, "truncate tax_rates;")
}

func insertTaxRates(ctx context.Context, db *sql.DB, legalEntityID string, taxRates *[]api.TaxRate) error {
	for _, cc := range *taxRates {
		r := models.TaxRate{}
		r.ID = *cc.Id
		r.AccountReference = null.StringFrom(cc.AccountReference)
		r.Country = null.StringFrom(cc.Country)
		r.Rate = null.IntFrom(int(cc.Rate))
		r.Name = null.StringFrom(cc.Name)
		r.Code = null.StringFrom(cc.Code)
		r.StatusActive = null.BoolFrom(cc.StatusActive)
		r.LegalEntityID = null.StringFrom(legalEntityID)

		err := r.Insert(ctx, db, boil.Infer())
		if err != nil {
			return err
		}
	}

	return nil
}
