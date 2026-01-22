package yokoy

import (
	"context"
	"database/sql"
	"errors"

	"github.com/dylanj/yokoy/api"
	"github.com/volatiletech/null/v8"
)

func fetchCategories(ctx context.Context, legalEntityId string, c api.ClientWithResponsesInterface) (*[]api.Category, error) {
	p := &api.GetLegalEntitiesLegalEntityIdCategoriesParams{}

	h, err := c.GetLegalEntitiesLegalEntityIdCategoriesWithResponse(ctx, legalEntityId, p)

	if err != nil {
		return nil, err
	}

	if h.StatusCode() != 200 {
		return nil, errors.New("got non 200 response code")
	}

	return h.JSON200.Categories, nil

}

func truncateCategories(ctx context.Context, db *sql.DB) {
	db.ExecContext(ctx, "truncate categories;")
}

func insertCategories(ctx context.Context, db *sql.DB, legalEntityID string, categories *[]api.Category) error {
	if len(*categories) == 0 {
		return nil
	}

	columns := []string{"id", "account_reference", "charge_to_employee", "name", "status_active", "legal_entity_id"}
	values := make([][]interface{}, 0, len(*categories))

	for _, c := range *categories {
		row := []interface{}{
			*c.Id,
			NullStringValue(null.StringFrom(c.AccountReference)),
			NullBoolValue(null.BoolFromPtr(c.ChargeToEmployee)),
			NullStringValue(null.StringFrom(c.Name)),
			NullBoolValue(null.BoolFrom(c.StatusActive)),
			legalEntityID,
		}
		values = append(values, row)
	}

	return BulkInsert(ctx, db, "categories", columns, values)
}
