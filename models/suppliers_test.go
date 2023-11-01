// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testSuppliers(t *testing.T) {
	t.Parallel()

	query := Suppliers()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testSuppliersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Supplier{}
	if err = randomize.Struct(seed, o, supplierDBTypes, true, supplierColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Supplier struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Suppliers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSuppliersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Supplier{}
	if err = randomize.Struct(seed, o, supplierDBTypes, true, supplierColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Supplier struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Suppliers().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Suppliers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSuppliersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Supplier{}
	if err = randomize.Struct(seed, o, supplierDBTypes, true, supplierColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Supplier struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SupplierSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Suppliers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSuppliersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Supplier{}
	if err = randomize.Struct(seed, o, supplierDBTypes, true, supplierColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Supplier struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := SupplierExists(ctx, tx, o.LegalEntityID, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Supplier exists: %s", err)
	}
	if !e {
		t.Errorf("Expected SupplierExists to return true, but got false.")
	}
}

func testSuppliersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Supplier{}
	if err = randomize.Struct(seed, o, supplierDBTypes, true, supplierColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Supplier struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	supplierFound, err := FindSupplier(ctx, tx, o.LegalEntityID, o.ID)
	if err != nil {
		t.Error(err)
	}

	if supplierFound == nil {
		t.Error("want a record, got nil")
	}
}

func testSuppliersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Supplier{}
	if err = randomize.Struct(seed, o, supplierDBTypes, true, supplierColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Supplier struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Suppliers().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testSuppliersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Supplier{}
	if err = randomize.Struct(seed, o, supplierDBTypes, true, supplierColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Supplier struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Suppliers().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testSuppliersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	supplierOne := &Supplier{}
	supplierTwo := &Supplier{}
	if err = randomize.Struct(seed, supplierOne, supplierDBTypes, false, supplierColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Supplier struct: %s", err)
	}
	if err = randomize.Struct(seed, supplierTwo, supplierDBTypes, false, supplierColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Supplier struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = supplierOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = supplierTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Suppliers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testSuppliersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	supplierOne := &Supplier{}
	supplierTwo := &Supplier{}
	if err = randomize.Struct(seed, supplierOne, supplierDBTypes, false, supplierColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Supplier struct: %s", err)
	}
	if err = randomize.Struct(seed, supplierTwo, supplierDBTypes, false, supplierColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Supplier struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = supplierOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = supplierTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Suppliers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func supplierBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Supplier) error {
	*o = Supplier{}
	return nil
}

func supplierAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Supplier) error {
	*o = Supplier{}
	return nil
}

func supplierAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Supplier) error {
	*o = Supplier{}
	return nil
}

func supplierBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Supplier) error {
	*o = Supplier{}
	return nil
}

func supplierAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Supplier) error {
	*o = Supplier{}
	return nil
}

func supplierBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Supplier) error {
	*o = Supplier{}
	return nil
}

func supplierAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Supplier) error {
	*o = Supplier{}
	return nil
}

func supplierBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Supplier) error {
	*o = Supplier{}
	return nil
}

func supplierAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Supplier) error {
	*o = Supplier{}
	return nil
}

func testSuppliersHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Supplier{}
	o := &Supplier{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, supplierDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Supplier object: %s", err)
	}

	AddSupplierHook(boil.BeforeInsertHook, supplierBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	supplierBeforeInsertHooks = []SupplierHook{}

	AddSupplierHook(boil.AfterInsertHook, supplierAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	supplierAfterInsertHooks = []SupplierHook{}

	AddSupplierHook(boil.AfterSelectHook, supplierAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	supplierAfterSelectHooks = []SupplierHook{}

	AddSupplierHook(boil.BeforeUpdateHook, supplierBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	supplierBeforeUpdateHooks = []SupplierHook{}

	AddSupplierHook(boil.AfterUpdateHook, supplierAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	supplierAfterUpdateHooks = []SupplierHook{}

	AddSupplierHook(boil.BeforeDeleteHook, supplierBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	supplierBeforeDeleteHooks = []SupplierHook{}

	AddSupplierHook(boil.AfterDeleteHook, supplierAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	supplierAfterDeleteHooks = []SupplierHook{}

	AddSupplierHook(boil.BeforeUpsertHook, supplierBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	supplierBeforeUpsertHooks = []SupplierHook{}

	AddSupplierHook(boil.AfterUpsertHook, supplierAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	supplierAfterUpsertHooks = []SupplierHook{}
}

func testSuppliersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Supplier{}
	if err = randomize.Struct(seed, o, supplierDBTypes, true, supplierColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Supplier struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Suppliers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSuppliersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Supplier{}
	if err = randomize.Struct(seed, o, supplierDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Supplier struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(supplierColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Suppliers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSuppliersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Supplier{}
	if err = randomize.Struct(seed, o, supplierDBTypes, true, supplierColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Supplier struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testSuppliersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Supplier{}
	if err = randomize.Struct(seed, o, supplierDBTypes, true, supplierColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Supplier struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SupplierSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testSuppliersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Supplier{}
	if err = randomize.Struct(seed, o, supplierDBTypes, true, supplierColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Supplier struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Suppliers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	supplierDBTypes = map[string]string{`ID`: `text`, `LegalEntityID`: `text`, `City`: `text`, `CountryCode`: `text`, `ExternalID`: `text`, `Name`: `text`, `SecondaryName`: `text`, `ShortName`: `text`, `StatusActive`: `boolean`, `Street`: `text`, `TaxNumber`: `text`, `URL`: `text`, `ZipCode`: `text`, `DefaultApproverID`: `text`, `DefaultCategoryID`: `text`, `DefaultCostCenter`: `text`, `DefaultPaymentTermID`: `text`, `SupplierID`: `text`}
	_               = bytes.MinRead
)

func testSuppliersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(supplierPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(supplierAllColumns) == len(supplierPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Supplier{}
	if err = randomize.Struct(seed, o, supplierDBTypes, true, supplierColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Supplier struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Suppliers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, supplierDBTypes, true, supplierPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Supplier struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testSuppliersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(supplierAllColumns) == len(supplierPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Supplier{}
	if err = randomize.Struct(seed, o, supplierDBTypes, true, supplierColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Supplier struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Suppliers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, supplierDBTypes, true, supplierPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Supplier struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(supplierAllColumns, supplierPrimaryKeyColumns) {
		fields = supplierAllColumns
	} else {
		fields = strmangle.SetComplement(
			supplierAllColumns,
			supplierPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := SupplierSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testSuppliersUpsert(t *testing.T) {
	t.Parallel()

	if len(supplierAllColumns) == len(supplierPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Supplier{}
	if err = randomize.Struct(seed, &o, supplierDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Supplier struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Supplier: %s", err)
	}

	count, err := Suppliers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, supplierDBTypes, false, supplierPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Supplier struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Supplier: %s", err)
	}

	count, err = Suppliers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
