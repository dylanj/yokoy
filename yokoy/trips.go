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

func fetchTrips(ctx context.Context, c api.ClientWithResponsesInterface) (*[]api.Trip, error) {
	p := &api.GetTripsParams{}
	h, err := c.GetTripsWithResponse(ctx, p)

	if err != nil {
		return nil, err
	}

	if h.StatusCode() != 200 {
		return nil, errors.New("got non 200 response code")
	}

	return h.JSON200.Trips, nil
}

func truncateTrips(ctx context.Context, db *sql.DB) {
	db.ExecContext(ctx, "truncate trips;")
}

func insertTrips(ctx context.Context, db *sql.DB, trips *[]api.Trip) error {
	for _, t := range *trips {
		r := models.Trip{}
		r.ID = *t.Id
		r.Currency = null.StringFromPtr(t.Currency)
		r.StartDateTime = null.TimeFromPtr(parseTime(t.StartDatetime))
		r.EndDateTime = null.TimeFromPtr(parseTime(t.EndDatetime))
		r.LegalEntityID = null.StringFromPtr(t.LegalEntityId)
		r.Name = null.StringFromPtr(t.Name)
		r.Status = null.StringFrom(string(t.Status))
		r.TotalClaim = null.IntFrom(int(*t.TotalClaim))
		r.UserID = null.StringFromPtr(t.UserId)

		err := r.Insert(ctx, db, boil.Infer())
		if err != nil {
			return err
		}
	}

	return nil
}
