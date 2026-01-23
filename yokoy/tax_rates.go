package yokoy

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/dylanj/yokoy/api"
	"github.com/aarondl/null/v9"
)

func fetchTaxRates(ctx context.Context, legalEntityId string, c api.ClientWithResponsesInterface) (*[]api.TaxRate, error) {
	p := &api.ListTaxRatesParams{}

	h, err := c.ListTaxRatesWithResponse(ctx, api.LegalEntityIdInPath(legalEntityId), p)
	if err != nil {
		return nil, err
	}

	if h.StatusCode() != 200 {
		return nil, errors.New("got non 200 response code")
	}

	// The API returns a union type. We need to unmarshal from the response body directly
	if h.JSON200 == nil || h.JSON200.TaxRates == nil {
		return &[]api.TaxRate{}, nil
	}
	// Unmarshal the response body to get tax rates
	var responseData struct {
		TaxRates []api.TaxRate `json:"taxRates"`
	}
	if err := json.Unmarshal(h.Body, &responseData); err != nil {
		return nil, err
	}
	return &responseData.TaxRates, nil
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
