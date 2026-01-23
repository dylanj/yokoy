package yokoy

import (
	"context"
	"database/sql"
	"errors"

	"github.com/aarondl/null/v9"
	"github.com/dylanj/yokoy/api"
)

func fetchTransactions(ctx context.Context, c api.ClientWithResponsesInterface, filter string) (*[]api.Transaction, error) {
	p := &api.ListTransactionsParams{
		Filter: &filter,
	}
	h, err := c.ListTransactionsWithResponse(ctx, p)

	if err != nil {
		return nil, err
	}

	if h.StatusCode() != 200 {
		return nil, errors.New("got non 200 response code")
	}

	return h.JSON200.Transactions, nil
}

func truncateTransactions(ctx context.Context, db *sql.DB) {
	db.ExecContext(ctx, "truncate transactions;")
}

func insertTransactions(ctx context.Context, db *sql.DB, transactions *[]api.Transaction) error {
	if len(*transactions) == 0 {
		return nil
	}

	columns := []string{"id", "billing_amount", "billing_currency", "booking_date", "card_id", "country", "created", "expense_id", "expense_status", "is_credit_note", "legal_entity_id", "matching_status", "mcc_sic_code", "merchant", "source_amount", "source_currency", "statement_reference", "transaction_date"}
	values := make([][]interface{}, 0, len(*transactions))

	for _, t := range *transactions {
		billingAmount := NullDecimalFromFloat32(t.BillingAmount)

		var sourceAmount interface{}
		if t.SourceAmount != nil {
			amount := NullDecimalFromFloat32(*t.SourceAmount)
			sourceAmount = &amount
		}

		var expenseStatus interface{}
		if t.ExpenseStatus != nil {
			expenseStatus = float32(*t.ExpenseStatus)
		}

		var matchingStatus interface{}
		if t.MatchingStatus != nil {
			matchingStatus = float32(*t.MatchingStatus)
		}

		row := []interface{}{
			*t.Id,
			billingAmount,
			NullStringValue(null.StringFrom(t.BillingCurrency)),
			NullTimeValue(null.TimeFromPtr(t.BookingDate)),
			NullStringValue(null.StringFrom(t.CardId)),
			NullStringValue(null.StringFromPtr(t.Country)),
			NullTimeValue(null.TimeFromPtr(t.Created)),
			NullStringValue(null.StringFromPtr(t.ExpenseId)),
			expenseStatus,
			NullBoolValue(null.BoolFromPtr(t.IsCreditNote)),
			NullStringValue(null.StringFrom(t.LegalEntityId)),
			matchingStatus,
			NullStringValue(null.StringFromPtr(t.MccSicCode)),
			NullStringValue(null.StringFrom(t.Merchant)),
			sourceAmount,
			NullStringValue(null.StringFromPtr(t.SourceCurrency)),
			NullStringValue(null.StringFromPtr(t.StatementReference)),
			NullTimeValue(null.TimeFromPtr(parseDate(&t.TransactionDate))),
		}
		values = append(values, row)
	}

	return BulkInsert(ctx, db, "transactions", columns, values)
}
