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
	for _, p := range *invoiceCategories {
		r := models.InvoiceCategory{}
		r.ID = *p.Id
		r.Name = null.StringFrom(p.Name)
		r.Description = null.StringFromPtr(p.Description)
		r.AccountReference = null.StringFrom(p.AccountReference)
		r.StatusActive = null.BoolFrom(p.StatusActive)
		r.LegalEntityID = null.StringFrom(legalEntityID)

		err := r.Insert(ctx, db, boil.Infer())
		if err != nil {
			return err
		}
	}

	return nil
}
