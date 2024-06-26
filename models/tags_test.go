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

func testTags(t *testing.T) {
	t.Parallel()

	query := Tags()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testTagsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Tag{}
	if err = randomize.Struct(seed, o, tagDBTypes, true, tagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tag struct: %s", err)
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

	count, err := Tags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTagsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Tag{}
	if err = randomize.Struct(seed, o, tagDBTypes, true, tagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Tags().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Tags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTagsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Tag{}
	if err = randomize.Struct(seed, o, tagDBTypes, true, tagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TagSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Tags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTagsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Tag{}
	if err = randomize.Struct(seed, o, tagDBTypes, true, tagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := TagExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Tag exists: %s", err)
	}
	if !e {
		t.Errorf("Expected TagExists to return true, but got false.")
	}
}

func testTagsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Tag{}
	if err = randomize.Struct(seed, o, tagDBTypes, true, tagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	tagFound, err := FindTag(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if tagFound == nil {
		t.Error("want a record, got nil")
	}
}

func testTagsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Tag{}
	if err = randomize.Struct(seed, o, tagDBTypes, true, tagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Tags().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testTagsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Tag{}
	if err = randomize.Struct(seed, o, tagDBTypes, true, tagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Tags().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testTagsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	tagOne := &Tag{}
	tagTwo := &Tag{}
	if err = randomize.Struct(seed, tagOne, tagDBTypes, false, tagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tag struct: %s", err)
	}
	if err = randomize.Struct(seed, tagTwo, tagDBTypes, false, tagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = tagOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = tagTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Tags().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testTagsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	tagOne := &Tag{}
	tagTwo := &Tag{}
	if err = randomize.Struct(seed, tagOne, tagDBTypes, false, tagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tag struct: %s", err)
	}
	if err = randomize.Struct(seed, tagTwo, tagDBTypes, false, tagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = tagOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = tagTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Tags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func tagBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Tag) error {
	*o = Tag{}
	return nil
}

func tagAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Tag) error {
	*o = Tag{}
	return nil
}

func tagAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Tag) error {
	*o = Tag{}
	return nil
}

func tagBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Tag) error {
	*o = Tag{}
	return nil
}

func tagAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Tag) error {
	*o = Tag{}
	return nil
}

func tagBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Tag) error {
	*o = Tag{}
	return nil
}

func tagAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Tag) error {
	*o = Tag{}
	return nil
}

func tagBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Tag) error {
	*o = Tag{}
	return nil
}

func tagAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Tag) error {
	*o = Tag{}
	return nil
}

func testTagsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Tag{}
	o := &Tag{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, tagDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Tag object: %s", err)
	}

	AddTagHook(boil.BeforeInsertHook, tagBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	tagBeforeInsertHooks = []TagHook{}

	AddTagHook(boil.AfterInsertHook, tagAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	tagAfterInsertHooks = []TagHook{}

	AddTagHook(boil.AfterSelectHook, tagAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	tagAfterSelectHooks = []TagHook{}

	AddTagHook(boil.BeforeUpdateHook, tagBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	tagBeforeUpdateHooks = []TagHook{}

	AddTagHook(boil.AfterUpdateHook, tagAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	tagAfterUpdateHooks = []TagHook{}

	AddTagHook(boil.BeforeDeleteHook, tagBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	tagBeforeDeleteHooks = []TagHook{}

	AddTagHook(boil.AfterDeleteHook, tagAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	tagAfterDeleteHooks = []TagHook{}

	AddTagHook(boil.BeforeUpsertHook, tagBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	tagBeforeUpsertHooks = []TagHook{}

	AddTagHook(boil.AfterUpsertHook, tagAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	tagAfterUpsertHooks = []TagHook{}
}

func testTagsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Tag{}
	if err = randomize.Struct(seed, o, tagDBTypes, true, tagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Tags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTagsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Tag{}
	if err = randomize.Struct(seed, o, tagDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Tag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(tagColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Tags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTagsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Tag{}
	if err = randomize.Struct(seed, o, tagDBTypes, true, tagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tag struct: %s", err)
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

func testTagsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Tag{}
	if err = randomize.Struct(seed, o, tagDBTypes, true, tagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TagSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTagsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Tag{}
	if err = randomize.Struct(seed, o, tagDBTypes, true, tagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Tags().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	tagDBTypes = map[string]string{`ID`: `text`, `DimensionCode`: `text`, `Code`: `text`, `Name`: `text`, `LegalEntityID`: `text`, `StatusActive`: `boolean`}
	_          = bytes.MinRead
)

func testTagsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(tagPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(tagAllColumns) == len(tagPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Tag{}
	if err = randomize.Struct(seed, o, tagDBTypes, true, tagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Tags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, tagDBTypes, true, tagPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Tag struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testTagsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(tagAllColumns) == len(tagPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Tag{}
	if err = randomize.Struct(seed, o, tagDBTypes, true, tagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Tags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, tagDBTypes, true, tagPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Tag struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(tagAllColumns, tagPrimaryKeyColumns) {
		fields = tagAllColumns
	} else {
		fields = strmangle.SetComplement(
			tagAllColumns,
			tagPrimaryKeyColumns,
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

	slice := TagSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testTagsUpsert(t *testing.T) {
	t.Parallel()

	if len(tagAllColumns) == len(tagPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Tag{}
	if err = randomize.Struct(seed, &o, tagDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Tag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Tag: %s", err)
	}

	count, err := Tags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, tagDBTypes, false, tagPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Tag struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Tag: %s", err)
	}

	count, err = Tags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
