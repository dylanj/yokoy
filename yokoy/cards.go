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

func fetchCompanyCards(ctx context.Context, legalEntityId string, c api.ClientWithResponsesInterface) (*[]api.CompanyCard, error) {
	p := &api.GetLegalEntitiesLegalEntityIdCompanyCardsParams{}

	h, err := c.GetLegalEntitiesLegalEntityIdCompanyCardsWithResponse(ctx, legalEntityId, p)

	if err != nil {
		return nil, err
	}

	if h.StatusCode() != 200 {
		return nil, errors.New("got non 200 response code")
	}

	return h.JSON200.CompanyCards, nil

}

func truncateCompanyCards(ctx context.Context, db *sql.DB) {
	db.ExecContext(ctx, "truncate company_cards;")
}

func insertCompanyCards(ctx context.Context, db *sql.DB, legalEntityID string, companyCards *[]api.CompanyCard) error {
	for _, cc := range *companyCards {
		r := models.CompanyCard{}
		r.ID = *cc.Id
		r.AccountReference = null.StringFrom(cc.AccountReference)
		r.Currency = null.StringFrom(cc.Currency)
		r.Name = null.StringFrom(cc.Name)
		r.Description = null.StringFromPtr(cc.Description)
		r.Number = null.StringFrom(cc.Number)
		r.OwnerID = null.StringFrom(cc.OwnerId)
		r.StatusActive = null.BoolFrom(cc.StatusActive)
		r.LegalEntityID = null.StringFrom(legalEntityID)

		err := r.Insert(ctx, db, boil.Infer())
		if err != nil {
			return err
		}
	}

	return nil
}
