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

func testCostCenters(t *testing.T) {
	t.Parallel()

	query := CostCenters()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testCostCentersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CostCenter{}
	if err = randomize.Struct(seed, o, costCenterDBTypes, true, costCenterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CostCenter struct: %s", err)
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

	count, err := CostCenters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCostCentersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CostCenter{}
	if err = randomize.Struct(seed, o, costCenterDBTypes, true, costCenterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CostCenter struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := CostCenters().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := CostCenters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCostCentersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CostCenter{}
	if err = randomize.Struct(seed, o, costCenterDBTypes, true, costCenterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CostCenter struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CostCenterSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := CostCenters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCostCentersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CostCenter{}
	if err = randomize.Struct(seed, o, costCenterDBTypes, true, costCenterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CostCenter struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := CostCenterExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if CostCenter exists: %s", err)
	}
	if !e {
		t.Errorf("Expected CostCenterExists to return true, but got false.")
	}
}

func testCostCentersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CostCenter{}
	if err = randomize.Struct(seed, o, costCenterDBTypes, true, costCenterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CostCenter struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	costCenterFound, err := FindCostCenter(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if costCenterFound == nil {
		t.Error("want a record, got nil")
	}
}

func testCostCentersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CostCenter{}
	if err = randomize.Struct(seed, o, costCenterDBTypes, true, costCenterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CostCenter struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = CostCenters().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testCostCentersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CostCenter{}
	if err = randomize.Struct(seed, o, costCenterDBTypes, true, costCenterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CostCenter struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := CostCenters().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testCostCentersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	costCenterOne := &CostCenter{}
	costCenterTwo := &CostCenter{}
	if err = randomize.Struct(seed, costCenterOne, costCenterDBTypes, false, costCenterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CostCenter struct: %s", err)
	}
	if err = randomize.Struct(seed, costCenterTwo, costCenterDBTypes, false, costCenterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CostCenter struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = costCenterOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = costCenterTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := CostCenters().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testCostCentersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	costCenterOne := &CostCenter{}
	costCenterTwo := &CostCenter{}
	if err = randomize.Struct(seed, costCenterOne, costCenterDBTypes, false, costCenterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CostCenter struct: %s", err)
	}
	if err = randomize.Struct(seed, costCenterTwo, costCenterDBTypes, false, costCenterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CostCenter struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = costCenterOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = costCenterTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CostCenters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func costCenterBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *CostCenter) error {
	*o = CostCenter{}
	return nil
}

func costCenterAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *CostCenter) error {
	*o = CostCenter{}
	return nil
}

func costCenterAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *CostCenter) error {
	*o = CostCenter{}
	return nil
}

func costCenterBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *CostCenter) error {
	*o = CostCenter{}
	return nil
}

func costCenterAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *CostCenter) error {
	*o = CostCenter{}
	return nil
}

func costCenterBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *CostCenter) error {
	*o = CostCenter{}
	return nil
}

func costCenterAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *CostCenter) error {
	*o = CostCenter{}
	return nil
}

func costCenterBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *CostCenter) error {
	*o = CostCenter{}
	return nil
}

func costCenterAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *CostCenter) error {
	*o = CostCenter{}
	return nil
}

func testCostCentersHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &CostCenter{}
	o := &CostCenter{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, costCenterDBTypes, false); err != nil {
		t.Errorf("Unable to randomize CostCenter object: %s", err)
	}

	AddCostCenterHook(boil.BeforeInsertHook, costCenterBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	costCenterBeforeInsertHooks = []CostCenterHook{}

	AddCostCenterHook(boil.AfterInsertHook, costCenterAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	costCenterAfterInsertHooks = []CostCenterHook{}

	AddCostCenterHook(boil.AfterSelectHook, costCenterAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	costCenterAfterSelectHooks = []CostCenterHook{}

	AddCostCenterHook(boil.BeforeUpdateHook, costCenterBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	costCenterBeforeUpdateHooks = []CostCenterHook{}

	AddCostCenterHook(boil.AfterUpdateHook, costCenterAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	costCenterAfterUpdateHooks = []CostCenterHook{}

	AddCostCenterHook(boil.BeforeDeleteHook, costCenterBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	costCenterBeforeDeleteHooks = []CostCenterHook{}

	AddCostCenterHook(boil.AfterDeleteHook, costCenterAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	costCenterAfterDeleteHooks = []CostCenterHook{}

	AddCostCenterHook(boil.BeforeUpsertHook, costCenterBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	costCenterBeforeUpsertHooks = []CostCenterHook{}

	AddCostCenterHook(boil.AfterUpsertHook, costCenterAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	costCenterAfterUpsertHooks = []CostCenterHook{}
}

func testCostCentersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CostCenter{}
	if err = randomize.Struct(seed, o, costCenterDBTypes, true, costCenterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CostCenter struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CostCenters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCostCentersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CostCenter{}
	if err = randomize.Struct(seed, o, costCenterDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CostCenter struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(costCenterColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := CostCenters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCostCentersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CostCenter{}
	if err = randomize.Struct(seed, o, costCenterDBTypes, true, costCenterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CostCenter struct: %s", err)
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

func testCostCentersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CostCenter{}
	if err = randomize.Struct(seed, o, costCenterDBTypes, true, costCenterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CostCenter struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CostCenterSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testCostCentersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CostCenter{}
	if err = randomize.Struct(seed, o, costCenterDBTypes, true, costCenterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CostCenter struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := CostCenters().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	costCenterDBTypes = map[string]string{`ID`: `text`, `ApprovalLimit`: `numeric`, `ApproverID`: `text`, `AutoApprovalLimit`: `numeric`, `Code`: `text`, `DelegateExpiry`: `timestamp without time zone`, `DelegateID`: `text`, `Description`: `text`, `Name`: `text`, `ParentID`: `text`, `LegalEntityID`: `text`, `StatusActive`: `boolean`}
	_                 = bytes.MinRead
)

func testCostCentersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(costCenterPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(costCenterAllColumns) == len(costCenterPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &CostCenter{}
	if err = randomize.Struct(seed, o, costCenterDBTypes, true, costCenterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CostCenter struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CostCenters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, costCenterDBTypes, true, costCenterPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CostCenter struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testCostCentersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(costCenterAllColumns) == len(costCenterPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &CostCenter{}
	if err = randomize.Struct(seed, o, costCenterDBTypes, true, costCenterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CostCenter struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CostCenters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, costCenterDBTypes, true, costCenterPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CostCenter struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(costCenterAllColumns, costCenterPrimaryKeyColumns) {
		fields = costCenterAllColumns
	} else {
		fields = strmangle.SetComplement(
			costCenterAllColumns,
			costCenterPrimaryKeyColumns,
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

	slice := CostCenterSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testCostCentersUpsert(t *testing.T) {
	t.Parallel()

	if len(costCenterAllColumns) == len(costCenterPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := CostCenter{}
	if err = randomize.Struct(seed, &o, costCenterDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CostCenter struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert CostCenter: %s", err)
	}

	count, err := CostCenters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, costCenterDBTypes, false, costCenterPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CostCenter struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert CostCenter: %s", err)
	}

	count, err = CostCenters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
