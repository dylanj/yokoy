package yokoy

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/dylanj/yokoy/api"
	"github.com/shopspring/decimal"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/types"
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
	if len(*invoices) == 0 {
		return nil
	}

	// Collect invoices for bulk insert
	invoiceColumns := []string{"id", "country", "currency", "date", "gross_amount", "invoice_number", "is_credit_node", "net_amount", "payment_term_id", "posting_date", "purchase_order_ids", "submitters", "supplier_id", "taxable_amount", "legal_entity_id", "bank_account", "bank_country", "bank_key", "bank_number", "external_id", "iban", "swift_code"}
	invoiceValues := make([][]interface{}, 0, len(*invoices))

	// Collect line items for bulk insert
	lineItemColumns := []string{"invoice_id", "category_id", "cost_object_id", "description", "gross", "item_price", "net", "purchase_order_id", "purchase_order_item_id", "quantity", "tax_rate_id", "unit"}
	lineItemValues := make([][]interface{}, 0)

	for _, inv := range *invoices {
		var grossAmount, netAmount, taxableAmount decimal.Decimal
		if inv.GrossAmount != nil {
			grossAmount = NullDecimalFromFloat32(*inv.GrossAmount)
		}
		if inv.NetAmount != nil {
			netAmount = NullDecimalFromFloat32(*inv.NetAmount)
		}
		if inv.TaxableAmount != nil {
			taxableAmount = NullDecimalFromFloat32(*inv.TaxableAmount)
		}

		var purchaseOrderIds types.StringArray
		if inv.PurchaseOrderIds != nil {
			purchaseOrderIds = *inv.PurchaseOrderIds
		}

		var submitters types.StringArray
		if inv.Submitters != nil {
			submitters = *inv.Submitters
		}

		var date, postingDate *time.Time
		if inv.Date != nil {
			date = &inv.Date.Time
		}
		if inv.PostingDate != nil {
			postingDate = &inv.PostingDate.Time
		}

		var bankAccount, bankCountry, bankKey, bankNumber, externalID, iban, swiftCode *string
		if inv.BankAccount != nil {
			bankAccount = inv.BankAccount.BankAccount
			bankCountry = inv.BankAccount.BankCountry
			bankKey = inv.BankAccount.BankKey
			bankNumber = inv.BankAccount.BankNumber
			externalID = inv.BankAccount.ExternalId
			iban = inv.BankAccount.Iban
			swiftCode = inv.BankAccount.SwiftCode
		}

		row := []interface{}{
			*inv.Id,
			NullStringValue(null.StringFromPtr(inv.Country)),
			NullStringValue(null.StringFromPtr(inv.Currency)),
			NullTimeValue(null.TimeFromPtr(date)),
			grossAmount,
			NullStringValue(null.StringFrom(inv.InvoiceNumber)),
			NullBoolValue(null.BoolFromPtr(inv.IsCreditNote)),
			netAmount,
			NullStringValue(null.StringFromPtr(inv.PaymentTermId)),
			NullTimeValue(null.TimeFromPtr(postingDate)),
			purchaseOrderIds,
			submitters,
			NullStringValue(null.StringFromPtr(inv.SupplierId)),
			taxableAmount,
			legalEntityID,
			NullStringValue(null.StringFromPtr(bankAccount)),
			NullStringValue(null.StringFromPtr(bankCountry)),
			NullStringValue(null.StringFromPtr(bankKey)),
			NullStringValue(null.StringFromPtr(bankNumber)),
			NullStringValue(null.StringFromPtr(externalID)),
			NullStringValue(null.StringFromPtr(iban)),
			NullStringValue(null.StringFromPtr(swiftCode)),
		}
		invoiceValues = append(invoiceValues, row)

		// Collect line items
		for _, li := range inv.LineItems {
			var itemPrice, quantity *decimal.Decimal
			if li.ItemPrice != nil {
				price := NullDecimalFromFloat32(*li.ItemPrice)
				itemPrice = &price
			}
			if li.Quantity != nil {
				qty := NullDecimalFromFloat32(*li.Quantity)
				quantity = &qty
			}

			lineRow := []interface{}{
				*inv.Id,
				NullStringValue(null.StringFromPtr(li.CategoryId)),
				NullStringValue(null.StringFromPtr(li.CostObjectId)),
				NullStringValue(null.StringFromPtr(li.Description)),
				NullDecimalFromFloat32(li.Gross),
				itemPrice,
				NullDecimalFromFloat32(li.Net),
				NullStringValue(null.StringFromPtr(li.PurchaseOrderId)),
				NullStringValue(null.StringFromPtr(li.PurchaseOrderItemId)),
				quantity,
				NullStringValue(null.StringFromPtr(li.TaxRateId)),
				NullStringValue(null.StringFromPtr(li.Unit)),
			}
			lineItemValues = append(lineItemValues, lineRow)
		}
	}

	// Bulk insert invoices
	if err := BulkInsert(ctx, db, "invoices", invoiceColumns, invoiceValues); err != nil {
		return err
	}

	// Bulk insert line items
	if len(lineItemValues) > 0 {
		if err := BulkInsert(ctx, db, "invoice_line_items", lineItemColumns, lineItemValues); err != nil {
			return err
		}
	}

	return nil
}
