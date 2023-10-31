package yokoy

import (
	"context"
	"database/sql"
	"errors"

	"github.com/salesfive/yokoy/api"
	"github.com/salesfive/yokoy/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
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
	for _, le := range *legalEntities {
		dbLE := models.LegalEntity{}
		dbLE.Name = null.StringFrom(*le.Name)
		dbLE.ID = *le.Id
		dbLE.Code = null.StringFrom(le.Code)
		dbLE.Language = null.StringFrom(string(le.Language))

		err := dbLE.Insert(ctx, db, boil.Infer())
		if err != nil {
			return err
		}
	}

	return nil
}
