package yokoy

import (
	"context"
	"database/sql"
	"errors"

	"github.com/dylanj/yokoy/api"
	"github.com/aarondl/null/v9"
)

func fetchTags(ctx context.Context, legalEntityId string, c api.ClientWithResponsesInterface) (*[]api.Tag, error) {
	p := &api.ListTagsParams{}

	h, err := c.ListTagsWithResponse(ctx, api.LegalEntityIdInPath(legalEntityId), p)

	if err != nil {
		return nil, err
	}

	if h.StatusCode() != 200 {
		return nil, errors.New("got non 200 response code")
	}

	return h.JSON200.Tags, nil

}

func truncateTags(ctx context.Context, db *sql.DB) {
	db.ExecContext(ctx, "truncate tags;")
}

func insertTags(ctx context.Context, db *sql.DB, legalEntityID string, tags *[]api.Tag) error {
	if len(*tags) == 0 {
		return nil
	}

	columns := []string{"id", "dimension_code", "name", "code", "status_active", "legal_entity_id"}
	values := make([][]interface{}, 0, len(*tags))

	for _, tag := range *tags {
		// Skip tags with nil IDs - they cannot be inserted
		if tag.Id == nil {
			continue
		}

		row := []interface{}{
			*tag.Id,
			NullStringValue(null.StringFrom(tag.DimensionCode)),
			NullStringValue(null.StringFrom(tag.Name)),
			NullStringValue(null.StringFrom(tag.Code)),
			NullBoolValue(null.BoolFrom(tag.StatusActive)),
			legalEntityID,
		}
		values = append(values, row)
	}

	// Use ON CONFLICT DO NOTHING to handle duplicate tags across legal entities
	opts := BulkInsertOptions{
		OnConflict:    "DO NOTHING",
		ConflictTarget: "id",
	}
	return BulkInsert(ctx, db, "tags", columns, values, opts)
}
