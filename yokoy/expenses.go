package yokoy

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/dylanj/yokoy/api"
	"github.com/shopspring/decimal"
	"github.com/aarondl/null/v9"
)

func isoString(t time.Time) string {
	return t.Format("2006-01-02")
}

func filterString(field string, startDate, endDate time.Time) string {
	start := isoString(startDate)
	end := isoString(endDate)

	return field + " ge " + start + " and " + field + " lt " + end
}

func forEachMonthSince(field string, startDate time.Time) []string {
	filters := make([]string, 0)

	now := time.Now()
	endDate := startDate.AddDate(0, 1, 0)
	cnt := 0

	for {
		cnt = cnt + 1
		if cnt > 500 {
			// Log limit hit - this should not normally happen
			break
		}

		f := filterString(field, startDate, endDate)
		filters = append(filters, f)

		if endDate.After(now) {
			break
		}

		startDate = startDate.AddDate(0, 1, 0)
		endDate = startDate.AddDate(0, 1, 0)
	}

	return filters
}

func fetchExpenses(ctx context.Context, c api.ClientWithResponsesInterface, filter string) (*[]api.Expense, error) {
	//filter := "created ge 2024-01-01 and created lt 2024-01-02"
	//filter := "created ge \"Mon, 12 Feb 2023 10:08:45 GMT\""
	//filter := "legalEntityId eq BzTHfPWrnm"
	p := &api.ListExpensesParams{
		Filter: &filter,
	}
	h, err := c.ListExpensesWithResponse(ctx, p)

	if err != nil {
		return nil, err
	}

	if h.StatusCode() != 200 {
		return nil, errors.New("got non 200 response code")
	}
	//fmt.Println(string(h.Body))

	return h.JSON200.Expenses, nil
}

func truncateExpenses(ctx context.Context, db *sql.DB) {
	db.ExecContext(ctx, "truncate expenses;")
	db.ExecContext(ctx, "truncate expense_cost_centers;")
	db.ExecContext(ctx, "truncate expense_tax_items;")
	db.ExecContext(ctx, "truncate expense_approver_ids;")
	db.ExecContext(ctx, "truncate expense_event_logs;")
}

func insertExpenses(ctx context.Context, db *sql.DB, expenses *[]api.Expense) error {
	if len(*expenses) == 0 {
		return nil
	}

	// Collect expenses for bulk insert
	expenseColumns := []string{"id", "additional_charges", "category_id", "country", "created", "currency", "description", "expense_date", "expense_end_date", "expense_type", "is_credit_note", "last_modified", "legal_entity_id", "payment_method", "posting_date", "status", "tax_number", "total_amount", "total_claim", "trip_id", "user_id"}
	expenseValues := make([][]interface{}, 0, len(*expenses))

	// Collect related records for bulk insert
	costCenterColumns := []string{"expense_id", "cost_center_id", "percent_weight"}
	costCenterValues := make([][]interface{}, 0)

	taxItemColumns := []string{"expense_id", "rate_id", "gross", "tax"}
	taxItemValues := make([][]interface{}, 0)

	approverColumns := []string{"expense_id", "approver_id"}
	approverValues := make([][]interface{}, 0)

	eventLogColumns := []string{"expense_id", "action_type", "comment", "name", "timestamp", "user_id"}
	eventLogValues := make([][]interface{}, 0)

	for _, exp := range *expenses {
		var additionalCharges interface{}
		if exp.AdditionalCharges != nil {
			additionalCharges = int(*exp.AdditionalCharges)
		}

		var expenseType interface{}
		if exp.ExpenseType != nil {
			expenseType = string(*exp.ExpenseType)
		}

		var paymentMethod interface{}
		if exp.PaymentMethod != nil {
			paymentMethod = string(*exp.PaymentMethod)
		}

		var totalAmount, totalClaim *decimal.Decimal
		if exp.TotalAmount != nil {
			amount := NullDecimalFromFloat32(*exp.TotalAmount)
			totalAmount = &amount
		}
		if exp.TotalClaim != nil {
			claim := NullDecimalFromFloat32(*exp.TotalClaim)
			totalClaim = &claim
		}

		row := []interface{}{
			*exp.Id,
			additionalCharges,
			NullStringValue(null.StringFromPtr(exp.CategoryId)),
			NullStringValue(null.StringFromPtr(exp.Country)),
			NullTimeValue(null.TimeFromPtr(exp.Created)),
			NullStringValue(null.StringFromPtr(exp.Currency)),
			NullStringValue(null.StringFromPtr(exp.Description)),
			NullTimeValue(null.TimeFromPtr(parseDate(exp.ExpenseDate))),
			NullTimeValue(null.TimeFromPtr(parseTime(exp.ExpenseEndDate))),
			expenseType,
			NullBoolValue(null.BoolFromPtr(exp.IsCreditNote)),
			NullTimeValue(null.TimeFromPtr(exp.LastModified)),
			NullStringValue(null.StringFromPtr(exp.LegalEntityId)),
			paymentMethod,
			NullTimeValue(null.TimeFromPtr(parseDate(exp.PostingDate))),
			string(exp.Status),
			NullStringValue(null.StringFromPtr(exp.TaxNumber)),
			totalAmount,
			totalClaim,
			NullStringValue(null.StringFromPtr(exp.TripId)),
			NullStringValue(null.StringFromPtr(exp.UserId)),
		}
		expenseValues = append(expenseValues, row)

		// Collect cost center items
		for _, cci := range *exp.CostCenterItems {
			var percentWeight *decimal.Decimal
			if cci.PctWeight != nil {
				weight := NullDecimalFromFloat32(*cci.PctWeight)
				percentWeight = &weight
			}
			costCenterValues = append(costCenterValues, []interface{}{
				*exp.Id,
				NullStringValue(null.StringFromPtr(cci.Id)),
				percentWeight,
			})
		}

		// Collect tax items
		for _, ti := range *exp.TaxItems {
			var gross, tax *decimal.Decimal
			if ti.Gross != nil {
				g := NullDecimalFromFloat32(*ti.Gross)
				gross = &g
			}
			if ti.Tax != nil {
				t := NullDecimalFromFloat32(*ti.Tax)
				tax = &t
			}
			taxItemValues = append(taxItemValues, []interface{}{
				*exp.Id,
				NullStringValue(null.StringFromPtr(ti.RateId)),
				gross,
				tax,
			})
		}

		// Collect approver IDs
		for _, a := range *exp.ApproverIds {
			approverValues = append(approverValues, []interface{}{
				*exp.Id,
				a,
			})
		}

		// Collect event logs
		for _, l := range *exp.EventLog {
			var actionType interface{}
			if l.ActionType != nil {
				actionType = string(*l.ActionType)
			}
			eventLogValues = append(eventLogValues, []interface{}{
				*exp.Id,
				actionType,
				NullStringValue(null.StringFromPtr(l.Comment)),
				NullStringValue(null.StringFromPtr(l.Name)),
				NullTimeValue(null.TimeFromPtr(l.Timestamp)),
				NullStringValue(null.StringFromPtr(l.UserId)),
			})
		}
	}

	// Bulk insert expenses
	if err := BulkInsert(ctx, db, "expenses", expenseColumns, expenseValues); err != nil {
		return err
	}

	// Bulk insert related records
	if len(costCenterValues) > 0 {
		if err := BulkInsert(ctx, db, "expense_cost_centers", costCenterColumns, costCenterValues); err != nil {
			return err
		}
	}

	if len(taxItemValues) > 0 {
		if err := BulkInsert(ctx, db, "expense_tax_items", taxItemColumns, taxItemValues); err != nil {
			return err
		}
	}

	if len(approverValues) > 0 {
		if err := BulkInsert(ctx, db, "expense_approver_ids", approverColumns, approverValues); err != nil {
			return err
		}
	}

	if len(eventLogValues) > 0 {
		if err := BulkInsert(ctx, db, "expense_event_logs", eventLogColumns, eventLogValues); err != nil {
			return err
		}
	}

	return nil
}
