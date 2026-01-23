package yokoy

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/dylanj/yokoy/api"
	"github.com/shopspring/decimal"
	"github.com/aarondl/null/v9"
	"github.com/aarondl/sqlboiler/v4/types"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

func fetchInvoices(ctx context.Context, legalEntityId string, c api.ClientWithResponsesInterface) (*[]api.Invoice, error) {
	// Use underlying client to get raw response and parse manually
	// This handles date fields that can be "NAN"
	clientWithResponses, ok := c.(*api.ClientWithResponses)
	if !ok {
		return nil, errors.New("cannot access underlying client for raw request")
	}

	p := &api.ListInvoicesParams{}
	httpResp, err := clientWithResponses.ListInvoices(ctx, api.LegalEntityIdInPath(legalEntityId), p)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != 200 {
		return nil, errors.New("got non 200 response code")
	}

	bodyBytes, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}

	// Parse response manually to handle "NAN" date values
	var responseData struct {
		Invoices []struct {
			Country             string                 `json:"country"`
			Created             *string                `json:"created"`
			Currency            *string                `json:"currency"`
			CustomInformation   *map[string]string     `json:"customInformation"`
			Date                *string                `json:"date"` // Parse as string to handle "NAN"
			DocumentData        *string                `json:"documentData"`
			DueDate             *string                `json:"dueDate"` // Parse as string to handle "NAN"
			EventLog            *[]json.RawMessage     `json:"eventLog"`
			ExportJobId         *string                `json:"exportJobId"`
			FxRates             *map[string]interface{} `json:"fxRates"`
			GrossAmount         *float32               `json:"grossAmount"`
			Id                  *string                `json:"id"`
			InvoiceNumber       string                 `json:"invoiceNumber"`
			IsCreditNote        *bool                  `json:"isCreditNote"`
			LastModified        *string                `json:"lastModified"`
			LineItems           []json.RawMessage      `json:"lineItems"`
			NetAmount           *float32               `json:"netAmount"`
			PaymentBlocked      *bool                  `json:"paymentBlocked"`
			PaymentInformation  *json.RawMessage        `json:"paymentInformation"`
			PaymentTermId       *string                `json:"paymentTermId"`
			PostingDate         *string                `json:"postingDate"` // Parse as string to handle "NAN"
			PurchaseOrderIds    *[]string              `json:"purchaseOrderIds"`
			QrInfo              *json.RawMessage       `json:"qrInfo"`
			ServiceDate         interface{}             `json:"serviceDate"`
			Status              *int                   `json:"status"`
			Submitters          *[]string              `json:"submitters"`
			SupplierId          *string                `json:"supplierId"`
			TaxableAmount       *float32               `json:"taxableAmount"`
		} `json:"invoices"`
	}

	if err := json.Unmarshal(bodyBytes, &responseData); err != nil {
		return nil, err
	}

	// Convert to api.Invoice, handling "NAN" dates
	invoices := make([]api.Invoice, 0, len(responseData.Invoices))
	for _, inv := range responseData.Invoices {
		// Parse dates, handling "NAN" values
		var date *openapi_types.Date
		if inv.Date != nil && *inv.Date != "" && *inv.Date != "NAN" && *inv.Date != "NaN" {
			d := openapi_types.Date{}
			if err := d.UnmarshalText([]byte(*inv.Date)); err == nil {
				date = &d
			}
		}

		var dueDate *openapi_types.Date
		if inv.DueDate != nil && *inv.DueDate != "" && *inv.DueDate != "NAN" && *inv.DueDate != "NaN" {
			d := openapi_types.Date{}
			if err := d.UnmarshalText([]byte(*inv.DueDate)); err == nil {
				dueDate = &d
			}
		}

		var postingDate *openapi_types.Date
		if inv.PostingDate != nil && *inv.PostingDate != "" && *inv.PostingDate != "NAN" && *inv.PostingDate != "NaN" {
			d := openapi_types.Date{}
			if err := d.UnmarshalText([]byte(*inv.PostingDate)); err == nil {
				postingDate = &d
			}
		}

		// Parse timestamps
		var created, lastModified *time.Time
		if inv.Created != nil {
			if t, err := time.Parse(time.RFC3339Nano, *inv.Created); err == nil {
				created = &t
			} else if t, err := time.Parse(time.RFC3339, *inv.Created); err == nil {
				created = &t
			}
		}
		if inv.LastModified != nil {
			if t, err := time.Parse(time.RFC3339Nano, *inv.LastModified); err == nil {
				lastModified = &t
			} else if t, err := time.Parse(time.RFC3339, *inv.LastModified); err == nil {
				lastModified = &t
			}
		}

		// Parse payment information
		var paymentInfo *api.InvoiceBankAccount
		if inv.PaymentInformation != nil {
			var pi api.InvoiceBankAccount
			if err := json.Unmarshal(*inv.PaymentInformation, &pi); err == nil {
				paymentInfo = &pi
			}
		}

		// Parse line items
		lineItems := make([]api.InvoiceLineItem, 0, len(inv.LineItems))
		for _, liBytes := range inv.LineItems {
			var li api.InvoiceLineItem
			if err := json.Unmarshal(liBytes, &li); err == nil {
				lineItems = append(lineItems, li)
			}
		}

		invoice := api.Invoice{
			Country:            inv.Country,
			Created:            created,
			Currency:           inv.Currency,
			CustomInformation:  inv.CustomInformation,
			Date:               date,
			DocumentData:       inv.DocumentData,
			DueDate:            dueDate,
			ExportJobId:        inv.ExportJobId,
			FxRates:            inv.FxRates,
			GrossAmount:        inv.GrossAmount,
			Id:                 inv.Id,
			InvoiceNumber:      inv.InvoiceNumber,
			IsCreditNote:       inv.IsCreditNote,
			LastModified:       lastModified,
			LineItems:          lineItems,
			NetAmount:          inv.NetAmount,
			PaymentBlocked:     inv.PaymentBlocked,
			PaymentInformation: paymentInfo,
			PaymentTermId:      inv.PaymentTermId,
			PostingDate:        postingDate,
			PurchaseOrderIds:   inv.PurchaseOrderIds,
			Submitters:         inv.Submitters,
			SupplierId:         inv.SupplierId,
			TaxableAmount:      inv.TaxableAmount,
		}

		// Parse status if present (API returns int but type is string)
		if inv.Status != nil {
			statusStr := fmt.Sprintf("%d", *inv.Status)
			status := api.InvoiceStatus(statusStr)
			invoice.Status = &status
		}

		invoices = append(invoices, invoice)
	}

	return &invoices, nil
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
	invoiceColumns := []string{"id", "country", "currency", "date", "gross_amount", "invoice_number", "is_credit_node", "net_amount", "payment_term_id", "posting_date", "purchase_order_ids", "submitters", "supplier_id", "taxable_amount", "legal_entity_id", "bank_account", "bank_country", "bank_key", "bank_number", "external_id", "iban", "swift_code", "status"}
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
		if inv.PaymentInformation != nil {
			bankAccount = inv.PaymentInformation.BankAccount
			bankCountry = inv.PaymentInformation.BankCountry
			bankKey = inv.PaymentInformation.BankKey
			bankNumber = inv.PaymentInformation.BankNumber
			externalID = inv.PaymentInformation.ExternalId
			iban = inv.PaymentInformation.Iban
			swiftCode = inv.PaymentInformation.SwiftCode
		}

		var statusStr interface{}
		if inv.Status != nil {
			statusStr = string(*inv.Status)
		}

		row := []interface{}{
			*inv.Id,
			NullStringValue(null.StringFrom(inv.Country)),
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
			statusStr,
		}
		invoiceValues = append(invoiceValues, row)

		// Collect line items
		for _, li := range inv.LineItems {
			// ItemPrice is required (NOT NULL constraint), default to 0 if nil
			var itemPrice decimal.Decimal
			if li.ItemPrice != nil {
				itemPrice = NullDecimalFromFloat32(*li.ItemPrice)
			} else {
				itemPrice = decimal.Zero
			}

			// Quantity is required (NOT NULL constraint), default to 0 if nil
			var quantity decimal.Decimal
			if li.Quantity != nil {
				quantity = NullDecimalFromFloat32(*li.Quantity)
			} else {
				quantity = decimal.Zero
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
