// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testInvoiceCategories(t *testing.T) {
	t.Parallel()

	query := InvoiceCategories()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testInvoiceCategoriesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &InvoiceCategory{}
	if err = randomize.Struct(seed, o, invoiceCategoryDBTypes, true, invoiceCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InvoiceCategory struct: %s", err)
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

	count, err := InvoiceCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testInvoiceCategoriesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &InvoiceCategory{}
	if err = randomize.Struct(seed, o, invoiceCategoryDBTypes, true, invoiceCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InvoiceCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := InvoiceCategories().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := InvoiceCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testInvoiceCategoriesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &InvoiceCategory{}
	if err = randomize.Struct(seed, o, invoiceCategoryDBTypes, true, invoiceCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InvoiceCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := InvoiceCategorySlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := InvoiceCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testInvoiceCategoriesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &InvoiceCategory{}
	if err = randomize.Struct(seed, o, invoiceCategoryDBTypes, true, invoiceCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InvoiceCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := InvoiceCategoryExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if InvoiceCategory exists: %s", err)
	}
	if !e {
		t.Errorf("Expected InvoiceCategoryExists to return true, but got false.")
	}
}

func testInvoiceCategoriesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &InvoiceCategory{}
	if err = randomize.Struct(seed, o, invoiceCategoryDBTypes, true, invoiceCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InvoiceCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	invoiceCategoryFound, err := FindInvoiceCategory(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if invoiceCategoryFound == nil {
		t.Error("want a record, got nil")
	}
}

func testInvoiceCategoriesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &InvoiceCategory{}
	if err = randomize.Struct(seed, o, invoiceCategoryDBTypes, true, invoiceCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InvoiceCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = InvoiceCategories().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testInvoiceCategoriesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &InvoiceCategory{}
	if err = randomize.Struct(seed, o, invoiceCategoryDBTypes, true, invoiceCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InvoiceCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := InvoiceCategories().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testInvoiceCategoriesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	invoiceCategoryOne := &InvoiceCategory{}
	invoiceCategoryTwo := &InvoiceCategory{}
	if err = randomize.Struct(seed, invoiceCategoryOne, invoiceCategoryDBTypes, false, invoiceCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InvoiceCategory struct: %s", err)
	}
	if err = randomize.Struct(seed, invoiceCategoryTwo, invoiceCategoryDBTypes, false, invoiceCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InvoiceCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = invoiceCategoryOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = invoiceCategoryTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := InvoiceCategories().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testInvoiceCategoriesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	invoiceCategoryOne := &InvoiceCategory{}
	invoiceCategoryTwo := &InvoiceCategory{}
	if err = randomize.Struct(seed, invoiceCategoryOne, invoiceCategoryDBTypes, false, invoiceCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InvoiceCategory struct: %s", err)
	}
	if err = randomize.Struct(seed, invoiceCategoryTwo, invoiceCategoryDBTypes, false, invoiceCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InvoiceCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = invoiceCategoryOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = invoiceCategoryTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := InvoiceCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func invoiceCategoryBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *InvoiceCategory) error {
	*o = InvoiceCategory{}
	return nil
}

func invoiceCategoryAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *InvoiceCategory) error {
	*o = InvoiceCategory{}
	return nil
}

func invoiceCategoryAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *InvoiceCategory) error {
	*o = InvoiceCategory{}
	return nil
}

func invoiceCategoryBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *InvoiceCategory) error {
	*o = InvoiceCategory{}
	return nil
}

func invoiceCategoryAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *InvoiceCategory) error {
	*o = InvoiceCategory{}
	return nil
}

func invoiceCategoryBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *InvoiceCategory) error {
	*o = InvoiceCategory{}
	return nil
}

func invoiceCategoryAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *InvoiceCategory) error {
	*o = InvoiceCategory{}
	return nil
}

func invoiceCategoryBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *InvoiceCategory) error {
	*o = InvoiceCategory{}
	return nil
}

func invoiceCategoryAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *InvoiceCategory) error {
	*o = InvoiceCategory{}
	return nil
}

func testInvoiceCategoriesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &InvoiceCategory{}
	o := &InvoiceCategory{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, invoiceCategoryDBTypes, false); err != nil {
		t.Errorf("Unable to randomize InvoiceCategory object: %s", err)
	}

	AddInvoiceCategoryHook(boil.BeforeInsertHook, invoiceCategoryBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	invoiceCategoryBeforeInsertHooks = []InvoiceCategoryHook{}

	AddInvoiceCategoryHook(boil.AfterInsertHook, invoiceCategoryAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	invoiceCategoryAfterInsertHooks = []InvoiceCategoryHook{}

	AddInvoiceCategoryHook(boil.AfterSelectHook, invoiceCategoryAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	invoiceCategoryAfterSelectHooks = []InvoiceCategoryHook{}

	AddInvoiceCategoryHook(boil.BeforeUpdateHook, invoiceCategoryBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	invoiceCategoryBeforeUpdateHooks = []InvoiceCategoryHook{}

	AddInvoiceCategoryHook(boil.AfterUpdateHook, invoiceCategoryAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	invoiceCategoryAfterUpdateHooks = []InvoiceCategoryHook{}

	AddInvoiceCategoryHook(boil.BeforeDeleteHook, invoiceCategoryBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	invoiceCategoryBeforeDeleteHooks = []InvoiceCategoryHook{}

	AddInvoiceCategoryHook(boil.AfterDeleteHook, invoiceCategoryAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	invoiceCategoryAfterDeleteHooks = []InvoiceCategoryHook{}

	AddInvoiceCategoryHook(boil.BeforeUpsertHook, invoiceCategoryBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	invoiceCategoryBeforeUpsertHooks = []InvoiceCategoryHook{}

	AddInvoiceCategoryHook(boil.AfterUpsertHook, invoiceCategoryAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	invoiceCategoryAfterUpsertHooks = []InvoiceCategoryHook{}
}

func testInvoiceCategoriesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &InvoiceCategory{}
	if err = randomize.Struct(seed, o, invoiceCategoryDBTypes, true, invoiceCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InvoiceCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := InvoiceCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testInvoiceCategoriesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &InvoiceCategory{}
	if err = randomize.Struct(seed, o, invoiceCategoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize InvoiceCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(invoiceCategoryColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := InvoiceCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testInvoiceCategoriesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &InvoiceCategory{}
	if err = randomize.Struct(seed, o, invoiceCategoryDBTypes, true, invoiceCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InvoiceCategory struct: %s", err)
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

func testInvoiceCategoriesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &InvoiceCategory{}
	if err = randomize.Struct(seed, o, invoiceCategoryDBTypes, true, invoiceCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InvoiceCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := InvoiceCategorySlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testInvoiceCategoriesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &InvoiceCategory{}
	if err = randomize.Struct(seed, o, invoiceCategoryDBTypes, true, invoiceCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InvoiceCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := InvoiceCategories().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	invoiceCategoryDBTypes = map[string]string{`ID`: `text`, `AccountReference`: `text`, `Description`: `text`, `Name`: `text`, `StatusActive`: `boolean`, `LegalEntityID`: `text`}
	_                      = bytes.MinRead
)

func testInvoiceCategoriesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(invoiceCategoryPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(invoiceCategoryAllColumns) == len(invoiceCategoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &InvoiceCategory{}
	if err = randomize.Struct(seed, o, invoiceCategoryDBTypes, true, invoiceCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InvoiceCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := InvoiceCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, invoiceCategoryDBTypes, true, invoiceCategoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize InvoiceCategory struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testInvoiceCategoriesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(invoiceCategoryAllColumns) == len(invoiceCategoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &InvoiceCategory{}
	if err = randomize.Struct(seed, o, invoiceCategoryDBTypes, true, invoiceCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InvoiceCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := InvoiceCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, invoiceCategoryDBTypes, true, invoiceCategoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize InvoiceCategory struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(invoiceCategoryAllColumns, invoiceCategoryPrimaryKeyColumns) {
		fields = invoiceCategoryAllColumns
	} else {
		fields = strmangle.SetComplement(
			invoiceCategoryAllColumns,
			invoiceCategoryPrimaryKeyColumns,
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

	slice := InvoiceCategorySlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testInvoiceCategoriesUpsert(t *testing.T) {
	t.Parallel()

	if len(invoiceCategoryAllColumns) == len(invoiceCategoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := InvoiceCategory{}
	if err = randomize.Struct(seed, &o, invoiceCategoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize InvoiceCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert InvoiceCategory: %s", err)
	}

	count, err := InvoiceCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, invoiceCategoryDBTypes, false, invoiceCategoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize InvoiceCategory struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert InvoiceCategory: %s", err)
	}

	count, err = InvoiceCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
