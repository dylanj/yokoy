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

func fetchInvoices(ctx context.Context, legalEntityId string, c api.ClientWithResponsesInterface) (*[]api.Invoice, error) {
	p := &api.GetLegalEntitiesLegalEntityIdInvoicesParams{}
	h, err := c.GetLegalEntitiesLegalEntityIdInvoicesWithResponse(ctx, legalEntityId, p)

	if err != nil {
		return nil, err
	}

	if h.StatusCode() != 200 {
		return nil, errors.New("got non 200 response code")
	}

	return h.JSON200.Invoices, nil
}

func truncateInvoices(ctx context.Context, db *sql.DB) {
	db.ExecContext(ctx, "truncate invoices;")
	db.ExecContext(ctx, "truncate invoice_line_items;")
}

func insertInvoices(ctx context.Context, db *sql.DB, legalEntityID string, invoices *[]api.Invoice) error {
	for _, e := range *invoices {
		r := models.Invoice{}
		r.ID = *e.Id
		r.Country = null.StringFromPtr(e.Country)
		r.Currency = null.StringFromPtr(e.Currency)
		if e.Date != nil {
			r.Date = null.TimeFrom(e.Date.Time)
		}
		if e.GrossAmount != nil {

			r.GrossAmount = NullDecimalFromFloat32(*e.GrossAmount)
		}
		r.InvoiceNumber = null.StringFrom(e.InvoiceNumber)
		if e.IsCreditNote != nil {
			r.IsCreditNode = null.BoolFrom(*e.IsCreditNote)
		}
		if e.NetAmount != nil {

			r.NetAmount = NullDecimalFromFloat32(*e.NetAmount)
		}
		r.PaymentTermID = null.StringFromPtr(e.PaymentTermId)
		if e.PostingDate != nil {
			r.PostingDate = null.TimeFrom(e.PostingDate.Time)
		}
		if e.PurchaseOrderIds != nil {
			r.PurchaseOrderIds = *e.PurchaseOrderIds
		}
		/*
			if e.Status != nil {
				r.Status = null.StringFrom(string(*e.Status))
			}
		*/
		if e.Submitters != nil {
			r.Submitters = *e.Submitters
		}
		r.SupplierID = null.StringFromPtr(e.SupplierId)
		// service date can be a date, a range of dates or null
		//r.ServiceDate = null.TimeFromPtr(parseDate(e.ServiceDate.AsDateString()))

		if e.TaxableAmount != nil {
			r.TaxableAmount = NullDecimalFromFloat32(*e.TaxableAmount)
		}
		r.LegalEntityID = null.StringFrom(legalEntityID)
		if e.PostingDate != nil {
			r.PostingDate = null.TimeFrom(e.PostingDate.Time)
		}

		if e.BankAccount != nil {
			ba := e.BankAccount
			r.BankAccount = null.StringFromPtr(ba.BankAccount)
			r.BankCountry = null.StringFromPtr(ba.BankCountry)
			r.BankKey = null.StringFromPtr(ba.BankKey)
			r.BankNumber = null.StringFromPtr(ba.BankNumber)
			r.ExternalID = null.StringFromPtr(ba.ExternalId)
			r.Iban = null.StringFromPtr(ba.Iban)
			r.SwiftCode = null.StringFromPtr(ba.SwiftCode)
		}

		err := r.Insert(ctx, db, boil.Infer())
		if err != nil {
			return err
		}

		for _, li := range e.LineItems {
			r := models.InvoiceLineItem{}
			r.InvoiceID = null.StringFromPtr(e.Id)
			r.CategoryID = null.StringFromPtr(li.CategoryId)
			r.CostObjectID = null.StringFromPtr(li.CostObjectId)
			r.Description = null.StringFromPtr(li.Description)
			r.Gross = NullDecimalFromFloat32(li.Gross)
			if li.ItemPrice != nil {
				r.ItemPrice = NullDecimalFromFloat32(*li.ItemPrice)
			}

			r.Net = NullDecimalFromFloat32(li.Net)
			r.PurchaseOrderID = null.StringFromPtr(li.PurchaseOrderId)
			r.PurchaseOrderItemID = null.StringFromPtr(li.PurchaseOrderItemId)
			if li.Quantity != nil {
				r.Quantity = NullDecimalFromFloat32(*li.Quantity)
			}

			// tags is being deserialized into a map
			//r.Tags = null.StringFrom(*li.Tags[])
			r.TaxRateID = null.StringFromPtr(li.TaxRateId)
			r.Unit = null.StringFromPtr(li.Unit)
			err := r.Insert(ctx, db, boil.Infer())
			if err != nil {
				return err
			}
		}
	}

	return nil
}
