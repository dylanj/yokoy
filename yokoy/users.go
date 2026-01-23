package yokoy

import (
	"context"
	"database/sql"
	"errors"

	"github.com/dylanj/yokoy/api"
	"github.com/aarondl/null/v9"
)

func fetchUsers(ctx context.Context, c api.ClientWithResponsesInterface) (*[]api.User, error) {
	p := &api.GetUsersParams{}
	h, err := c.GetUsersWithResponse(ctx, p)

	if err != nil {
		return nil, err
	}

	if h.StatusCode() != 200 {
		return nil, errors.New("got non 200 response code")
	}

	return h.JSON200.Users, nil
}

func truncateUsers(ctx context.Context, db *sql.DB) {
	db.ExecContext(ctx, "truncate users;")
}

func insertUsers(ctx context.Context, db *sql.DB, users *[]api.User) error {
	if len(*users) == 0 {
		return nil
	}

	columns := []string{"id", "cost_center_id", "email", "language", "first_name", "last_name", "legal_entity_id", "line_manager_id", "line_manager_threshold", "policy_id", "status_active", "submission_delegate_id"}
	values := make([][]interface{}, 0, len(*users))

	for _, u := range *users {
		var language interface{}
		if u.Language != nil {
			language = string(*u.Language)
		}

		var lineManagerThreshold interface{}
		if u.LineManagerThreshold != nil {
			lineManagerThreshold = int(*u.LineManagerThreshold)
		}

		row := []interface{}{
			*u.Id,
			NullStringValue(null.StringFromPtr(u.CostCenterId)),
			NullStringValue(null.StringFrom(u.Email)),
			language,
			NullStringValue(null.StringFrom(u.FirstName)),
			NullStringValue(null.StringFrom(u.LastName)),
			NullStringValue(null.StringFrom(u.LegalEntityId)),
			NullStringValue(null.StringFromPtr(u.LineManagerId)),
			lineManagerThreshold,
			NullStringValue(null.StringFromPtr(u.PolicyId)),
			NullBoolValue(null.BoolFrom(u.StatusActive)),
			NullStringValue(null.StringFromPtr(u.SubmissionDelegateId)),
		}
		values = append(values, row)
	}

	return BulkInsert(ctx, db, "users", columns, values)
}
