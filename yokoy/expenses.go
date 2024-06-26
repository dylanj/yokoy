package yokoy

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/dylanj/yokoy/api"
	"github.com/dylanj/yokoy/models"
	"github.com/ericlagergren/decimal"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/types"
)

func fetchExpenses(ctx context.Context, c api.ClientWithResponsesInterface) (*[]api.Expense, error) {
	filter := "created ge 2024-01-01 and created lt 2024-01-02"
	// filter := "created ge \"Mon, 12 Feb 2023 10:08:45 GMT\""
	//filter := "legalEntityId eq BzTHfPWrnm"
	p := &api.GetExpensesParams{
		Filter: &filter,
	}
	h, err := c.GetExpensesWithResponse(ctx, p)
	fmt.Println("finished?")

	if err != nil {
		return nil, err
	}

	if h.StatusCode() != 200 {
		return nil, errors.New("got non 200 response code")
	}
	fmt.Println(string(h.Body))

	return h.JSON200.Expenses, nil
}

func truncateExpenses(ctx context.Context, db *sql.DB) {
	db.ExecContext(ctx, "truncate expenses;")
	db.ExecContext(ctx, "truncate expense_cost_center;")
	db.ExecContext(ctx, "truncate expense_tax_items;")
	db.ExecContext(ctx, "truncate expense_approver_ids;")
	db.ExecContext(ctx, "truncate expense_event_logs;")
}

func insertExpenses(ctx context.Context, db *sql.DB, expenses *[]api.Expense) error {
	for _, e := range *expenses {
		r := models.Expense{}
		r.ID = *e.Id
		if e.AdditionalCharges != nil {
			r.AdditionalCharges = null.IntFrom(int(*e.AdditionalCharges))
		}
		r.CategoryID = null.StringFromPtr(e.CategoryId)
		r.Country = null.StringFromPtr(e.Country)
		r.Created = null.TimeFromPtr(parseTime(e.Created))
		r.Currency = null.StringFromPtr(e.Currency)
		r.Description = null.StringFromPtr(e.Description)
		r.ExpenseDate = null.TimeFromPtr(parseDate(e.ExpenseDate))
		r.ExpenseEndDate = null.TimeFromPtr(parseTime(e.ExpenseEndDate))
		if e.ExpenseType != nil {
			r.ExpenseType = null.StringFrom(string(*e.ExpenseType))
		}
		r.IsCreditNote = null.BoolFromPtr(e.IsCreditNote)
		r.LastModified = null.TimeFromPtr(parseTime(e.LastModified))
		r.LegalEntityID = null.StringFromPtr(e.LegalEntityId)
		if e.PaymentMethod != nil {
			r.PaymentMethod = null.StringFrom((string(*e.PaymentMethod)))
		}
		r.PostingDate = null.TimeFromPtr(parseDate(e.PostingDate))
		r.Status = null.StringFrom(string(e.Status))
		r.TaxNumber = null.StringFromPtr(e.TaxNumber)
		if e.TotalAmount != nil {
			r.TotalAmount = NullDecimalFromFloat32(*e.TotalAmount)
		}
		if e.TotalClaim != nil {
			r.TotalClaim = NullDecimalFromFloat32(*e.TotalClaim)
		}
		r.TripID = null.StringFromPtr(e.TripId)
		r.UserID = null.StringFromPtr(e.UserId)

		err := r.Insert(ctx, db, boil.Infer())
		if err != nil {
			return err
		}

		for _, cci := range *e.CostCenterItems {
			r := models.ExpenseCostCenter{}
			r.ExpenseID = *e.Id
			r.CostCenterID = *cci.Id
			if cci.PctWeight != nil {
				weight := decimal.Big{}
				weight.SetFloat64(*cci.PctWeight)
				r.PercentWeight = types.NewNullDecimal(&weight)

				null.IntFrom(int(*cci.PctWeight))
			}

			fmt.Println("inserting", r.ExpenseID, r.CostCenterID)
			err := r.Insert(ctx, db, boil.Infer())
			if err != nil {
				return err
			}
		}

		for _, ti := range *e.TaxItems {
			r := models.ExpenseTaxItem{}
			r.ExpenseID = *e.Id
			r.RateID = *ti.RateId

			r.Gross = NullDecimalFromFloat32(*ti.Gross)
			r.Tax = NullDecimalFromFloat32(*ti.Tax)

			err := r.Insert(ctx, db, boil.Infer())
			if err != nil {
				return err
			}
		}

		for _, a := range *e.ApproverIds {
			r := models.ExpenseApproverID{}
			r.ExpenseID = *e.Id
			r.ApproverID = a

			err := r.Insert(ctx, db, boil.Infer())
			if err != nil {
				return err
			}
		}

		for _, l := range *e.EventLog {
			r := models.ExpenseEventLog{}
			r.ExpenseID = null.StringFromPtr(e.Id)
			if l.ActionType != nil {
				r.ActionType = null.StringFrom(string(*l.ActionType))
			}
			r.Comment = null.StringFromPtr(l.Comment)
			r.Name = null.StringFromPtr(l.Name)
			r.Timestamp = null.TimeFromPtr(parseLongTime(l.Timestamp))
			r.UserID = null.StringFromPtr(l.UserId)

			err := r.Insert(ctx, db, boil.Infer())
			if err != nil {
				return err
			}
		}
	}

	return nil
}
