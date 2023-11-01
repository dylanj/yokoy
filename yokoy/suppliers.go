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

func fetchSuppliers(ctx context.Context, legalEntityId string, c api.ClientWithResponsesInterface) (*[]api.Supplier, error) {
	p := &api.GetLegalEntitiesLegalEntityIdSuppliersParams{}

	h, err := c.GetLegalEntitiesLegalEntityIdSuppliersWithResponse(ctx, legalEntityId, p)

	if err != nil {
		return nil, err
	}

	if h.StatusCode() != 200 {
		return nil, errors.New("got non 200 response code")
	}

	return h.JSON200.Suppliers, nil
}

func truncateSuppliers(ctx context.Context, db *sql.DB) {
	db.ExecContext(ctx, "truncate suppliers;")
}

func insertSuppliers(ctx context.Context, db *sql.DB, legalEntityID string, suppliers *[]api.Supplier) error {
	for _, p := range *suppliers {
		r := models.Supplier{}
		r.ID = *p.Id
		r.LegalEntityID = legalEntityID
		r.SupplierID = null.StringFrom(p.SupplierId)
		r.Name = null.StringFrom(p.Name)
		r.City = null.StringFromPtr(p.City)
		r.CountryCode = null.StringFromPtr(p.CountryCode)
		r.ExternalID = null.StringFrom(p.ExternalId)
		r.SecondaryName = null.StringFromPtr(p.SecondaryName)
		r.ShortName = null.StringFromPtr(p.ShortName)
		r.StatusActive = null.BoolFrom(p.StatusActive)
		r.Street = null.StringFromPtr(p.Street)
		r.TaxNumber = null.StringFromPtr(p.TaxNumber)
		r.URL = null.StringFromPtr(p.Url)
		r.ZipCode = null.StringFromPtr(p.ZipCode)
		r.DefaultApproverID = null.StringFromPtr(p.DefaultApproverId)
		r.DefaultCategoryID = null.StringFromPtr(p.DefaultCategoryId)
		r.DefaultCostCenter = null.StringFromPtr(p.DefaultCostCenterId)
		r.DefaultPaymentTermID = null.StringFromPtr(p.DefaultPaymentTermId)

		err := r.Insert(ctx, db, boil.Infer())
		if err != nil {
			return err
		}
	}

	return nil
}
