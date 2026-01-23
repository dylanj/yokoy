package yokoy

import (
	"context"
	"database/sql"
	"errors"

	"github.com/dylanj/yokoy/api"
	"github.com/aarondl/null/v9"
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
	if len(*taxRates) == 0 {
		return nil
	}

	columns := []string{"id", "account_reference", "country", "rate", "name", "code", "status_active", "legal_entity_id"}
	values := make([][]interface{}, 0, len(*taxRates))

	for _, tr := range *taxRates {
		row := []interface{}{
			*tr.Id,
			NullStringValue(null.StringFrom(tr.AccountReference)),
			NullStringValue(null.StringFrom(tr.Country)),
			NullDecimalFromFloat32(tr.Rate),
			NullStringValue(null.StringFrom(tr.Name)),
			NullStringValue(null.StringFrom(tr.Code)),
			NullBoolValue(null.BoolFrom(tr.StatusActive)),
			legalEntityID,
		}
		values = append(values, row)
	}

	return BulkInsert(ctx, db, "tax_rates", columns, values)
}
