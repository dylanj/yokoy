package yokoy

import (
	"context"
	"database/sql"
	"errors"

	"github.com/dylanj/yokoy/api"
	"github.com/aarondl/null/v9"
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
	if len(*suppliers) == 0 {
		return nil
	}

	columns := []string{"id", "legal_entity_id", "supplier_id", "name", "city", "country_code", "external_id", "secondary_name", "short_name", "status_active", "street", "tax_number", "url", "zip_code", "default_approver_id", "default_category_id", "default_cost_center", "default_payment_term_id"}
	values := make([][]interface{}, 0, len(*suppliers))

	for _, s := range *suppliers {
		row := []interface{}{
			*s.Id,
			legalEntityID,
			NullStringValue(null.StringFrom(s.SupplierId)),
			NullStringValue(null.StringFrom(s.Name)),
			NullStringValue(null.StringFromPtr(s.City)),
			NullStringValue(null.StringFromPtr(s.CountryCode)),
			NullStringValue(null.StringFrom(s.ExternalId)),
			NullStringValue(null.StringFromPtr(s.SecondaryName)),
			NullStringValue(null.StringFromPtr(s.ShortName)),
			NullBoolValue(null.BoolFrom(s.StatusActive)),
			NullStringValue(null.StringFromPtr(s.Street)),
			NullStringValue(null.StringFromPtr(s.TaxNumber)),
			NullStringValue(null.StringFromPtr(s.Url)),
			NullStringValue(null.StringFromPtr(s.ZipCode)),
			NullStringValue(null.StringFromPtr(s.DefaultApproverId)),
			NullStringValue(null.StringFromPtr(s.DefaultCategoryId)),
			NullStringValue(null.StringFromPtr(s.DefaultCostCenterId)),
			NullStringValue(null.StringFromPtr(s.DefaultPaymentTermId)),
		}
		values = append(values, row)
	}

	return BulkInsert(ctx, db, "suppliers", columns, values)
}
