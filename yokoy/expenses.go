package yokoy

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/dylanj/yokoy/api"
	"github.com/dylanj/yokoy/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func fetchExpenses(ctx context.Context, c api.ClientWithResponsesInterface) (*[]api.Expense, error) {
	p := &api.GetExpensesParams{}
	h, err := c.GetExpensesWithResponse(ctx, p)

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
}

func insertExpenses(ctx context.Context, db *sql.DB, expenses *[]api.Expense) error {
	for _, e := range *expenses {
		r := models.Expense{}
		r.ID = *e.Id
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
			r.TotalAmount = null.IntFrom(int(*e.TotalAmount))
		}
		if e.TotalClaim != nil {
			r.TotalClaim = null.IntFrom(int(*e.TotalClaim))
		}
		r.TripID = null.StringFromPtr(e.TripId)
		r.UserID = null.StringFromPtr(e.UserId)

		err := r.Insert(ctx, db, boil.Infer())
		if err != nil {
			return err
		}
	}

	return nil
}
