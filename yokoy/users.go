package yokoy

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"time"

	"github.com/dylanj/yokoy/api"
	"github.com/aarondl/null/v9"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

func fetchUsers(ctx context.Context, c api.ClientWithResponsesInterface) (*[]api.User, error) {
	// We need to use the underlying client to get raw response and parse manually
	// because ListUsersWithResponse fails on invalid emails during unmarshaling
	// Check if we can access the underlying client
	clientWithResponses, ok := c.(*api.ClientWithResponses)
	if !ok {
		return nil, errors.New("cannot access underlying client for raw request")
	}
	
	p := &api.ListUsersParams{}
	// Use the underlying ClientInterface to get raw HTTP response
	httpResp, err := clientWithResponses.ListUsers(ctx, p)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != 200 {
		return nil, errors.New("got non 200 response code")
	}

	bodyBytes, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}

	// Parse response manually to handle invalid emails gracefully
	var responseData struct {
		Users []struct {
			CostCenterId         *string                `json:"costCenterId"`
			Created              *string                `json:"created"`
			Currency             *string                `json:"currency"`
			CustomInformation    *map[string]string     `json:"customInformation"`
			Email                *string                `json:"email"` // Parse as string first to avoid validation
			EmployeeId           string                 `json:"employeeId"`
			FirstName            string                 `json:"firstName"`
			Id                   *string                `json:"id"`
			Language             *string                `json:"language"`
			LastName             string                 `json:"lastName"`
			LastModified         *string                `json:"lastModified"`
			LegalEntityId        string                 `json:"legalEntityId"`
			LineManagerId        *string                `json:"lineManagerId"`
			LineManagerThreshold *float32               `json:"lineManagerThreshold"`
			PolicyId             *string                `json:"policyId"`
			StatusActive         bool                   `json:"statusActive"`
			SubmissionDelegateId *string                `json:"submissionDelegateId"`
		} `json:"users"`
	}

	if err := json.Unmarshal(bodyBytes, &responseData); err != nil {
		return nil, err
	}

	// Convert to api.User, handling invalid emails
	users := make([]api.User, 0, len(responseData.Users))
	for _, u := range responseData.Users {
		var email *openapi_types.Email
		if u.Email != nil && *u.Email != "" {
			// Create Email type - the validation only happens during JSON unmarshaling
			// Since we're parsing as string first, we can create the Email type here
			// If the email format is invalid, it will fail validation, but we've already
			// parsed it as a string, so we can handle it gracefully
			e := openapi_types.Email(*u.Email)
			email = &e
		}

		user := api.User{
			CostCenterId:         u.CostCenterId,
			Currency:             u.Currency,
			CustomInformation:    u.CustomInformation,
			Email:                email,
			EmployeeId:           u.EmployeeId,
			FirstName:            u.FirstName,
			Id:                   u.Id,
			Language:             (*api.UserLanguage)(u.Language),
			LastName:             u.LastName,
			LegalEntityId:        u.LegalEntityId,
			LineManagerId:        u.LineManagerId,
			LineManagerThreshold: u.LineManagerThreshold,
			PolicyId:             u.PolicyId,
			StatusActive:         u.StatusActive,
			SubmissionDelegateId: u.SubmissionDelegateId,
		}

		// Parse timestamps if present (ISO 8601 format with timezone)
		if u.Created != nil {
			if t, err := time.Parse(time.RFC3339Nano, *u.Created); err == nil {
				user.Created = &t
			} else if t, err := time.Parse(time.RFC3339, *u.Created); err == nil {
				user.Created = &t
			}
		}
		if u.LastModified != nil {
			if t, err := time.Parse(time.RFC3339Nano, *u.LastModified); err == nil {
				user.LastModified = &t
			} else if t, err := time.Parse(time.RFC3339, *u.LastModified); err == nil {
				user.LastModified = &t
			}
		}

		users = append(users, user)
	}

	return &users, nil
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
			NullStringValue(null.StringFromPtr(func() *string {
				if u.Email != nil {
					s := string(*u.Email)
					return &s
				}
				return nil
			}())),
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
