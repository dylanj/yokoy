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

func fetchTags(ctx context.Context, legalEntityId string, c api.ClientWithResponsesInterface) (*[]api.Tag, error) {
	p := &api.GetLegalEntitiesLegalEntityIdTagsParams{}

	h, err := c.GetLegalEntitiesLegalEntityIdTagsWithResponse(ctx, legalEntityId, p)

	if err != nil {
		return nil, err
	}

	if h.StatusCode() != 200 {
		return nil, errors.New("got non 200 response code")
	}

	return h.JSON200.Tags, nil

}

func truncateTags(ctx context.Context, db *sql.DB) {
	db.ExecContext(ctx, "truncate tax_rates;")
}

func insertTags(ctx context.Context, db *sql.DB, legalEntityID string, tags *[]api.Tag) error {
	for _, cc := range *tags {
		r := models.Tag{}
		r.ID = *cc.Id
		r.DimensionCode = null.StringFrom(cc.DimensionCode)
		r.Name = null.StringFrom(cc.Name)
		r.Code = null.StringFrom(cc.Code)
		r.StatusActive = null.BoolFrom(cc.StatusActive)
		r.LegalEntityID = null.StringFrom(legalEntityID)

		err := r.Insert(ctx, db, boil.Infer())
		if err != nil {
			return err
		}
	}

	return nil
}
