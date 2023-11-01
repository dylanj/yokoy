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
	for _, u := range *users {
		dbU := models.User{}
		dbU.ID = *u.Id
		dbU.CostCenterID = null.StringFromPtr(u.CostCenterId)
		dbU.Email = null.StringFrom(u.Email)
		//dbU.EmployeeID = u.EmployeeId
		if u.Language != nil {
			dbU.Language = null.StringFrom(string(*u.Language))
		}
		dbU.FirstName = null.StringFrom(u.FirstName)
		dbU.LastName = null.StringFrom(u.LastName)
		dbU.LegalEntityID = null.StringFrom(u.LegalEntityId)
		dbU.LineManagerID = null.StringFromPtr(u.LineManagerId)
		if u.LineManagerThreshold != nil {
			dbU.LineManagerThreshold = null.IntFrom(int(*u.LineManagerThreshold))
		}
		dbU.PolicyID = null.StringFromPtr(u.PolicyId)
		dbU.StatusActive = null.BoolFrom(u.StatusActive)
		dbU.SubmissionDelegateID = null.StringFromPtr(u.SubmissionDelegateId)

		err := dbU.Insert(ctx, db, boil.Infer())
		if err != nil {
			return err
		}
	}

	return nil
}
