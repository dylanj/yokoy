// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// InvoiceCategory is an object representing the database table.
type InvoiceCategory struct {
	ID               string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	AccountReference null.String `boil:"account_reference" json:"account_reference,omitempty" toml:"account_reference" yaml:"account_reference,omitempty"`
	Description      null.String `boil:"description" json:"description,omitempty" toml:"description" yaml:"description,omitempty"`
	Name             null.String `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	StatusActive     null.Bool   `boil:"status_active" json:"status_active,omitempty" toml:"status_active" yaml:"status_active,omitempty"`
	LegalEntityID    null.String `boil:"legal_entity_id" json:"legal_entity_id,omitempty" toml:"legal_entity_id" yaml:"legal_entity_id,omitempty"`

	R *invoiceCategoryR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L invoiceCategoryL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var InvoiceCategoryColumns = struct {
	ID               string
	AccountReference string
	Description      string
	Name             string
	StatusActive     string
	LegalEntityID    string
}{
	ID:               "id",
	AccountReference: "account_reference",
	Description:      "description",
	Name:             "name",
	StatusActive:     "status_active",
	LegalEntityID:    "legal_entity_id",
}

var InvoiceCategoryTableColumns = struct {
	ID               string
	AccountReference string
	Description      string
	Name             string
	StatusActive     string
	LegalEntityID    string
}{
	ID:               "invoice_categories.id",
	AccountReference: "invoice_categories.account_reference",
	Description:      "invoice_categories.description",
	Name:             "invoice_categories.name",
	StatusActive:     "invoice_categories.status_active",
	LegalEntityID:    "invoice_categories.legal_entity_id",
}

// Generated where

var InvoiceCategoryWhere = struct {
	ID               whereHelperstring
	AccountReference whereHelpernull_String
	Description      whereHelpernull_String
	Name             whereHelpernull_String
	StatusActive     whereHelpernull_Bool
	LegalEntityID    whereHelpernull_String
}{
	ID:               whereHelperstring{field: "\"invoice_categories\".\"id\""},
	AccountReference: whereHelpernull_String{field: "\"invoice_categories\".\"account_reference\""},
	Description:      whereHelpernull_String{field: "\"invoice_categories\".\"description\""},
	Name:             whereHelpernull_String{field: "\"invoice_categories\".\"name\""},
	StatusActive:     whereHelpernull_Bool{field: "\"invoice_categories\".\"status_active\""},
	LegalEntityID:    whereHelpernull_String{field: "\"invoice_categories\".\"legal_entity_id\""},
}

// InvoiceCategoryRels is where relationship names are stored.
var InvoiceCategoryRels = struct {
}{}

// invoiceCategoryR is where relationships are stored.
type invoiceCategoryR struct {
}

// NewStruct creates a new relationship struct
func (*invoiceCategoryR) NewStruct() *invoiceCategoryR {
	return &invoiceCategoryR{}
}

// invoiceCategoryL is where Load methods for each relationship are stored.
type invoiceCategoryL struct{}

var (
	invoiceCategoryAllColumns            = []string{"id", "account_reference", "description", "name", "status_active", "legal_entity_id"}
	invoiceCategoryColumnsWithoutDefault = []string{"id"}
	invoiceCategoryColumnsWithDefault    = []string{"account_reference", "description", "name", "status_active", "legal_entity_id"}
	invoiceCategoryPrimaryKeyColumns     = []string{"id"}
	invoiceCategoryGeneratedColumns      = []string{}
)

type (
	// InvoiceCategorySlice is an alias for a slice of pointers to InvoiceCategory.
	// This should almost always be used instead of []InvoiceCategory.
	InvoiceCategorySlice []*InvoiceCategory
	// InvoiceCategoryHook is the signature for custom InvoiceCategory hook methods
	InvoiceCategoryHook func(context.Context, boil.ContextExecutor, *InvoiceCategory) error

	invoiceCategoryQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	invoiceCategoryType                 = reflect.TypeOf(&InvoiceCategory{})
	invoiceCategoryMapping              = queries.MakeStructMapping(invoiceCategoryType)
	invoiceCategoryPrimaryKeyMapping, _ = queries.BindMapping(invoiceCategoryType, invoiceCategoryMapping, invoiceCategoryPrimaryKeyColumns)
	invoiceCategoryInsertCacheMut       sync.RWMutex
	invoiceCategoryInsertCache          = make(map[string]insertCache)
	invoiceCategoryUpdateCacheMut       sync.RWMutex
	invoiceCategoryUpdateCache          = make(map[string]updateCache)
	invoiceCategoryUpsertCacheMut       sync.RWMutex
	invoiceCategoryUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var invoiceCategoryAfterSelectMu sync.Mutex
var invoiceCategoryAfterSelectHooks []InvoiceCategoryHook

var invoiceCategoryBeforeInsertMu sync.Mutex
var invoiceCategoryBeforeInsertHooks []InvoiceCategoryHook
var invoiceCategoryAfterInsertMu sync.Mutex
var invoiceCategoryAfterInsertHooks []InvoiceCategoryHook

var invoiceCategoryBeforeUpdateMu sync.Mutex
var invoiceCategoryBeforeUpdateHooks []InvoiceCategoryHook
var invoiceCategoryAfterUpdateMu sync.Mutex
var invoiceCategoryAfterUpdateHooks []InvoiceCategoryHook

var invoiceCategoryBeforeDeleteMu sync.Mutex
var invoiceCategoryBeforeDeleteHooks []InvoiceCategoryHook
var invoiceCategoryAfterDeleteMu sync.Mutex
var invoiceCategoryAfterDeleteHooks []InvoiceCategoryHook

var invoiceCategoryBeforeUpsertMu sync.Mutex
var invoiceCategoryBeforeUpsertHooks []InvoiceCategoryHook
var invoiceCategoryAfterUpsertMu sync.Mutex
var invoiceCategoryAfterUpsertHooks []InvoiceCategoryHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *InvoiceCategory) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range invoiceCategoryAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *InvoiceCategory) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range invoiceCategoryBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *InvoiceCategory) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range invoiceCategoryAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *InvoiceCategory) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range invoiceCategoryBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *InvoiceCategory) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range invoiceCategoryAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *InvoiceCategory) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range invoiceCategoryBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *InvoiceCategory) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range invoiceCategoryAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *InvoiceCategory) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range invoiceCategoryBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *InvoiceCategory) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range invoiceCategoryAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddInvoiceCategoryHook registers your hook function for all future operations.
func AddInvoiceCategoryHook(hookPoint boil.HookPoint, invoiceCategoryHook InvoiceCategoryHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		invoiceCategoryAfterSelectMu.Lock()
		invoiceCategoryAfterSelectHooks = append(invoiceCategoryAfterSelectHooks, invoiceCategoryHook)
		invoiceCategoryAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		invoiceCategoryBeforeInsertMu.Lock()
		invoiceCategoryBeforeInsertHooks = append(invoiceCategoryBeforeInsertHooks, invoiceCategoryHook)
		invoiceCategoryBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		invoiceCategoryAfterInsertMu.Lock()
		invoiceCategoryAfterInsertHooks = append(invoiceCategoryAfterInsertHooks, invoiceCategoryHook)
		invoiceCategoryAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		invoiceCategoryBeforeUpdateMu.Lock()
		invoiceCategoryBeforeUpdateHooks = append(invoiceCategoryBeforeUpdateHooks, invoiceCategoryHook)
		invoiceCategoryBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		invoiceCategoryAfterUpdateMu.Lock()
		invoiceCategoryAfterUpdateHooks = append(invoiceCategoryAfterUpdateHooks, invoiceCategoryHook)
		invoiceCategoryAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		invoiceCategoryBeforeDeleteMu.Lock()
		invoiceCategoryBeforeDeleteHooks = append(invoiceCategoryBeforeDeleteHooks, invoiceCategoryHook)
		invoiceCategoryBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		invoiceCategoryAfterDeleteMu.Lock()
		invoiceCategoryAfterDeleteHooks = append(invoiceCategoryAfterDeleteHooks, invoiceCategoryHook)
		invoiceCategoryAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		invoiceCategoryBeforeUpsertMu.Lock()
		invoiceCategoryBeforeUpsertHooks = append(invoiceCategoryBeforeUpsertHooks, invoiceCategoryHook)
		invoiceCategoryBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		invoiceCategoryAfterUpsertMu.Lock()
		invoiceCategoryAfterUpsertHooks = append(invoiceCategoryAfterUpsertHooks, invoiceCategoryHook)
		invoiceCategoryAfterUpsertMu.Unlock()
	}
}

// One returns a single invoiceCategory record from the query.
func (q invoiceCategoryQuery) One(ctx context.Context, exec boil.ContextExecutor) (*InvoiceCategory, error) {
	o := &InvoiceCategory{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for invoice_categories")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all InvoiceCategory records from the query.
func (q invoiceCategoryQuery) All(ctx context.Context, exec boil.ContextExecutor) (InvoiceCategorySlice, error) {
	var o []*InvoiceCategory

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to InvoiceCategory slice")
	}

	if len(invoiceCategoryAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all InvoiceCategory records in the query.
func (q invoiceCategoryQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count invoice_categories rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q invoiceCategoryQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if invoice_categories exists")
	}

	return count > 0, nil
}

// InvoiceCategories retrieves all the records using an executor.
func InvoiceCategories(mods ...qm.QueryMod) invoiceCategoryQuery {
	mods = append(mods, qm.From("\"invoice_categories\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"invoice_categories\".*"})
	}

	return invoiceCategoryQuery{q}
}

// FindInvoiceCategory retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindInvoiceCategory(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*InvoiceCategory, error) {
	invoiceCategoryObj := &InvoiceCategory{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"invoice_categories\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, invoiceCategoryObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from invoice_categories")
	}

	if err = invoiceCategoryObj.doAfterSelectHooks(ctx, exec); err != nil {
		return invoiceCategoryObj, err
	}

	return invoiceCategoryObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *InvoiceCategory) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no invoice_categories provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(invoiceCategoryColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	invoiceCategoryInsertCacheMut.RLock()
	cache, cached := invoiceCategoryInsertCache[key]
	invoiceCategoryInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			invoiceCategoryAllColumns,
			invoiceCategoryColumnsWithDefault,
			invoiceCategoryColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(invoiceCategoryType, invoiceCategoryMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(invoiceCategoryType, invoiceCategoryMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"invoice_categories\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"invoice_categories\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into invoice_categories")
	}

	if !cached {
		invoiceCategoryInsertCacheMut.Lock()
		invoiceCategoryInsertCache[key] = cache
		invoiceCategoryInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the InvoiceCategory.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *InvoiceCategory) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	invoiceCategoryUpdateCacheMut.RLock()
	cache, cached := invoiceCategoryUpdateCache[key]
	invoiceCategoryUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			invoiceCategoryAllColumns,
			invoiceCategoryPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update invoice_categories, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"invoice_categories\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, invoiceCategoryPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(invoiceCategoryType, invoiceCategoryMapping, append(wl, invoiceCategoryPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update invoice_categories row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for invoice_categories")
	}

	if !cached {
		invoiceCategoryUpdateCacheMut.Lock()
		invoiceCategoryUpdateCache[key] = cache
		invoiceCategoryUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q invoiceCategoryQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for invoice_categories")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for invoice_categories")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o InvoiceCategorySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), invoiceCategoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"invoice_categories\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, invoiceCategoryPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in invoiceCategory slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all invoiceCategory")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *InvoiceCategory) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no invoice_categories provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(invoiceCategoryColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	invoiceCategoryUpsertCacheMut.RLock()
	cache, cached := invoiceCategoryUpsertCache[key]
	invoiceCategoryUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			invoiceCategoryAllColumns,
			invoiceCategoryColumnsWithDefault,
			invoiceCategoryColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			invoiceCategoryAllColumns,
			invoiceCategoryPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert invoice_categories, could not build update column list")
		}

		ret := strmangle.SetComplement(invoiceCategoryAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(invoiceCategoryPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert invoice_categories, could not build conflict column list")
			}

			conflict = make([]string, len(invoiceCategoryPrimaryKeyColumns))
			copy(conflict, invoiceCategoryPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"invoice_categories\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(invoiceCategoryType, invoiceCategoryMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(invoiceCategoryType, invoiceCategoryMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert invoice_categories")
	}

	if !cached {
		invoiceCategoryUpsertCacheMut.Lock()
		invoiceCategoryUpsertCache[key] = cache
		invoiceCategoryUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single InvoiceCategory record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *InvoiceCategory) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no InvoiceCategory provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), invoiceCategoryPrimaryKeyMapping)
	sql := "DELETE FROM \"invoice_categories\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from invoice_categories")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for invoice_categories")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q invoiceCategoryQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no invoiceCategoryQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from invoice_categories")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for invoice_categories")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o InvoiceCategorySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(invoiceCategoryBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), invoiceCategoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"invoice_categories\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, invoiceCategoryPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from invoiceCategory slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for invoice_categories")
	}

	if len(invoiceCategoryAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *InvoiceCategory) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindInvoiceCategory(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *InvoiceCategorySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := InvoiceCategorySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), invoiceCategoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"invoice_categories\".* FROM \"invoice_categories\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, invoiceCategoryPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in InvoiceCategorySlice")
	}

	*o = slice

	return nil
}

// InvoiceCategoryExists checks if the InvoiceCategory row exists.
func InvoiceCategoryExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"invoice_categories\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if invoice_categories exists")
	}

	return exists, nil
}

// Exists checks if the InvoiceCategory row exists.
func (o *InvoiceCategory) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return InvoiceCategoryExists(ctx, exec, o.ID)
}
