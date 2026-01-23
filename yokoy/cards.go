package yokoy

import (
	"context"
	"database/sql"
	"errors"

	"github.com/dylanj/yokoy/api"
	"github.com/aarondl/null/v9"
)

func fetchCompanyCards(ctx context.Context, legalEntityId string, c api.ClientWithResponsesInterface) (*[]api.CompanyCard, error) {
	p := &api.ListCompanyCardsParams{}

	h, err := c.ListCompanyCardsWithResponse(ctx, api.LegalEntityIdInPath(legalEntityId), p)

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
	if len(*companyCards) == 0 {
		return nil
	}

	columns := []string{"id", "account_reference", "currency", "name", "description", "number", "owner_id", "status_active", "legal_entity_id"}
	values := make([][]interface{}, 0, len(*companyCards))

	for _, cc := range *companyCards {
		row := []interface{}{
			*cc.Id,
			NullStringValue(null.StringFrom(cc.AccountReference)),
			NullStringValue(null.StringFrom(cc.Currency)),
			NullStringValue(null.StringFrom(cc.Name)),
			NullStringValue(null.StringFromPtr(cc.Description)),
			NullStringValue(null.StringFrom(cc.Number)),
			NullStringValue(null.StringFrom(cc.OwnerId)),
			NullBoolValue(null.BoolFrom(cc.StatusActive)),
			legalEntityID,
		}
		values = append(values, row)
	}

	return BulkInsert(ctx, db, "company_cards", columns, values)
}
