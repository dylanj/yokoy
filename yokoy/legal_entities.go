package yokoy

import (
	"context"
	"database/sql"
	"errors"

	"github.com/dylanj/yokoy/api"
	"github.com/volatiletech/null/v8"
)

func fetchLegalEntities(ctx context.Context, c api.ClientWithResponsesInterface) (*[]api.LegalEntity, error) {
	p := &api.GetLegalEntitiesParams{}
	h, err := c.GetLegalEntitiesWithResponse(ctx, p)

	if err != nil {
		return nil, err
	}

	if h.StatusCode() != 200 {
		return nil, errors.New("got non 200 response code")
	}

	return h.JSON200.LegalEntities, nil
}

func truncateLegalEntities(ctx context.Context, db *sql.DB) {
	db.ExecContext(ctx, "truncate legal_entities;")
}

func insertLegalEntities(ctx context.Context, db *sql.DB, legalEntities *[]api.LegalEntity) error {
	if len(*legalEntities) == 0 {
		return nil
	}

	columns := []string{"id", "name", "code", "language"}
	values := make([][]interface{}, 0, len(*legalEntities))

	for _, le := range *legalEntities {
		row := []interface{}{
			*le.Id,
			NullStringValue(null.StringFromPtr(le.Name)),
			NullStringValue(null.StringFrom(le.Code)),
			string(le.Language),
		}
		values = append(values, row)
	}

	return BulkInsert(ctx, db, "legal_entities", columns, values)
}
