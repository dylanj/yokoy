package yokoy

import (
	"context"
	"database/sql"
	"errors"

	"github.com/dylanj/yokoy/api"
	"github.com/volatiletech/null/v8"
)

func fetchInvoiceCategories(ctx context.Context, legalEntityId string, c api.ClientWithResponsesInterface) (*[]api.InvoiceCategory, error) {
	p := &api.GetLegalEntitiesLegalEntityIdInvoiceCategoriesParams{}

	h, err := c.GetLegalEntitiesLegalEntityIdInvoiceCategoriesWithResponse(ctx, legalEntityId, p)

	if err != nil {
		return nil, err
	}

	if h.StatusCode() != 200 {
		return nil, errors.New("got non 200 response code")
	}

	return h.JSON200.InvoiceCategories, nil

}

func truncateInvoiceCategories(ctx context.Context, db *sql.DB) {
	db.ExecContext(ctx, "truncate invoice_categories;")
}

func insertInvoiceCategories(ctx context.Context, db *sql.DB, legalEntityID string, invoiceCategories *[]api.InvoiceCategory) error {
	if len(*invoiceCategories) == 0 {
		return nil
	}

	columns := []string{"id", "name", "description", "account_reference", "status_active", "legal_entity_id"}
	values := make([][]interface{}, 0, len(*invoiceCategories))

	for _, ic := range *invoiceCategories {
		row := []interface{}{
			*ic.Id,
			NullStringValue(null.StringFrom(ic.Name)),
			NullStringValue(null.StringFromPtr(ic.Description)),
			NullStringValue(null.StringFrom(ic.AccountReference)),
			NullBoolValue(null.BoolFrom(ic.StatusActive)),
			legalEntityID,
		}
		values = append(values, row)
	}

	return BulkInsert(ctx, db, "invoice_categories", columns, values)
}
