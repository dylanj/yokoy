package yokoy

import (
	"context"
	"database/sql"
	"errors"

	"github.com/dylanj/yokoy/api"
	"github.com/aarondl/null/v9"
)

func fetchTrips(ctx context.Context, c api.ClientWithResponsesInterface) (*[]api.Trip, error) {
	p := &api.ListTripsParams{}
	h, err := c.ListTripsWithResponse(ctx, p)

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
	if len(*trips) == 0 {
		return nil
	}

	columns := []string{"id", "currency", "start_date_time", "end_date_time", "legal_entity_id", "name", "status", "total_claim", "user_id"}
	values := make([][]interface{}, 0, len(*trips))

	for _, t := range *trips {
		row := []interface{}{
			*t.Id,
			NullStringValue(null.StringFromPtr(t.Currency)),
			NullTimeValue(null.TimeFromPtr(parseTime(t.StartDatetime))),
			NullTimeValue(null.TimeFromPtr(parseTime(t.EndDatetime))),
			NullStringValue(null.StringFromPtr(t.LegalEntityId)),
			NullStringValue(null.StringFromPtr(t.Name)),
			string(t.Status),
			int(*t.TotalClaim),
			NullStringValue(null.StringFromPtr(t.UserId)),
		}
		values = append(values, row)
	}

	return BulkInsert(ctx, db, "trips", columns, values)
}
