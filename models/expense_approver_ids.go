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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// ExpenseApproverID is an object representing the database table.
type ExpenseApproverID struct {
	ExpenseID  string `boil:"expense_id" json:"expense_id" toml:"expense_id" yaml:"expense_id"`
	ApproverID string `boil:"approver_id" json:"approver_id" toml:"approver_id" yaml:"approver_id"`

	R *expenseApproverIDR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L expenseApproverIDL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ExpenseApproverIDColumns = struct {
	ExpenseID  string
	ApproverID string
}{
	ExpenseID:  "expense_id",
	ApproverID: "approver_id",
}

var ExpenseApproverIDTableColumns = struct {
	ExpenseID  string
	ApproverID string
}{
	ExpenseID:  "expense_approver_ids.expense_id",
	ApproverID: "expense_approver_ids.approver_id",
}

// Generated where

var ExpenseApproverIDWhere = struct {
	ExpenseID  whereHelperstring
	ApproverID whereHelperstring
}{
	ExpenseID:  whereHelperstring{field: "\"expense_approver_ids\".\"expense_id\""},
	ApproverID: whereHelperstring{field: "\"expense_approver_ids\".\"approver_id\""},
}

// ExpenseApproverIDRels is where relationship names are stored.
var ExpenseApproverIDRels = struct {
}{}

// expenseApproverIDR is where relationships are stored.
type expenseApproverIDR struct {
}

// NewStruct creates a new relationship struct
func (*expenseApproverIDR) NewStruct() *expenseApproverIDR {
	return &expenseApproverIDR{}
}

// expenseApproverIDL is where Load methods for each relationship are stored.
type expenseApproverIDL struct{}

var (
	expenseApproverIDAllColumns            = []string{"expense_id", "approver_id"}
	expenseApproverIDColumnsWithoutDefault = []string{"expense_id", "approver_id"}
	expenseApproverIDColumnsWithDefault    = []string{}
	expenseApproverIDPrimaryKeyColumns     = []string{"expense_id", "approver_id"}
	expenseApproverIDGeneratedColumns      = []string{}
)

type (
	// ExpenseApproverIDSlice is an alias for a slice of pointers to ExpenseApproverID.
	// This should almost always be used instead of []ExpenseApproverID.
	ExpenseApproverIDSlice []*ExpenseApproverID
	// ExpenseApproverIDHook is the signature for custom ExpenseApproverID hook methods
	ExpenseApproverIDHook func(context.Context, boil.ContextExecutor, *ExpenseApproverID) error

	expenseApproverIDQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	expenseApproverIDType                 = reflect.TypeOf(&ExpenseApproverID{})
	expenseApproverIDMapping              = queries.MakeStructMapping(expenseApproverIDType)
	expenseApproverIDPrimaryKeyMapping, _ = queries.BindMapping(expenseApproverIDType, expenseApproverIDMapping, expenseApproverIDPrimaryKeyColumns)
	expenseApproverIDInsertCacheMut       sync.RWMutex
	expenseApproverIDInsertCache          = make(map[string]insertCache)
	expenseApproverIDUpdateCacheMut       sync.RWMutex
	expenseApproverIDUpdateCache          = make(map[string]updateCache)
	expenseApproverIDUpsertCacheMut       sync.RWMutex
	expenseApproverIDUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var expenseApproverIDAfterSelectMu sync.Mutex
var expenseApproverIDAfterSelectHooks []ExpenseApproverIDHook

var expenseApproverIDBeforeInsertMu sync.Mutex
var expenseApproverIDBeforeInsertHooks []ExpenseApproverIDHook
var expenseApproverIDAfterInsertMu sync.Mutex
var expenseApproverIDAfterInsertHooks []ExpenseApproverIDHook

var expenseApproverIDBeforeUpdateMu sync.Mutex
var expenseApproverIDBeforeUpdateHooks []ExpenseApproverIDHook
var expenseApproverIDAfterUpdateMu sync.Mutex
var expenseApproverIDAfterUpdateHooks []ExpenseApproverIDHook

var expenseApproverIDBeforeDeleteMu sync.Mutex
var expenseApproverIDBeforeDeleteHooks []ExpenseApproverIDHook
var expenseApproverIDAfterDeleteMu sync.Mutex
var expenseApproverIDAfterDeleteHooks []ExpenseApproverIDHook

var expenseApproverIDBeforeUpsertMu sync.Mutex
var expenseApproverIDBeforeUpsertHooks []ExpenseApproverIDHook
var expenseApproverIDAfterUpsertMu sync.Mutex
var expenseApproverIDAfterUpsertHooks []ExpenseApproverIDHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ExpenseApproverID) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range expenseApproverIDAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ExpenseApproverID) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range expenseApproverIDBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ExpenseApproverID) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range expenseApproverIDAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ExpenseApproverID) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range expenseApproverIDBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ExpenseApproverID) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range expenseApproverIDAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ExpenseApproverID) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range expenseApproverIDBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ExpenseApproverID) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range expenseApproverIDAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ExpenseApproverID) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range expenseApproverIDBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ExpenseApproverID) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range expenseApproverIDAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddExpenseApproverIDHook registers your hook function for all future operations.
func AddExpenseApproverIDHook(hookPoint boil.HookPoint, expenseApproverIDHook ExpenseApproverIDHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		expenseApproverIDAfterSelectMu.Lock()
		expenseApproverIDAfterSelectHooks = append(expenseApproverIDAfterSelectHooks, expenseApproverIDHook)
		expenseApproverIDAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		expenseApproverIDBeforeInsertMu.Lock()
		expenseApproverIDBeforeInsertHooks = append(expenseApproverIDBeforeInsertHooks, expenseApproverIDHook)
		expenseApproverIDBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		expenseApproverIDAfterInsertMu.Lock()
		expenseApproverIDAfterInsertHooks = append(expenseApproverIDAfterInsertHooks, expenseApproverIDHook)
		expenseApproverIDAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		expenseApproverIDBeforeUpdateMu.Lock()
		expenseApproverIDBeforeUpdateHooks = append(expenseApproverIDBeforeUpdateHooks, expenseApproverIDHook)
		expenseApproverIDBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		expenseApproverIDAfterUpdateMu.Lock()
		expenseApproverIDAfterUpdateHooks = append(expenseApproverIDAfterUpdateHooks, expenseApproverIDHook)
		expenseApproverIDAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		expenseApproverIDBeforeDeleteMu.Lock()
		expenseApproverIDBeforeDeleteHooks = append(expenseApproverIDBeforeDeleteHooks, expenseApproverIDHook)
		expenseApproverIDBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		expenseApproverIDAfterDeleteMu.Lock()
		expenseApproverIDAfterDeleteHooks = append(expenseApproverIDAfterDeleteHooks, expenseApproverIDHook)
		expenseApproverIDAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		expenseApproverIDBeforeUpsertMu.Lock()
		expenseApproverIDBeforeUpsertHooks = append(expenseApproverIDBeforeUpsertHooks, expenseApproverIDHook)
		expenseApproverIDBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		expenseApproverIDAfterUpsertMu.Lock()
		expenseApproverIDAfterUpsertHooks = append(expenseApproverIDAfterUpsertHooks, expenseApproverIDHook)
		expenseApproverIDAfterUpsertMu.Unlock()
	}
}

// One returns a single expenseApproverID record from the query.
func (q expenseApproverIDQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ExpenseApproverID, error) {
	o := &ExpenseApproverID{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for expense_approver_ids")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ExpenseApproverID records from the query.
func (q expenseApproverIDQuery) All(ctx context.Context, exec boil.ContextExecutor) (ExpenseApproverIDSlice, error) {
	var o []*ExpenseApproverID

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ExpenseApproverID slice")
	}

	if len(expenseApproverIDAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ExpenseApproverID records in the query.
func (q expenseApproverIDQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count expense_approver_ids rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q expenseApproverIDQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if expense_approver_ids exists")
	}

	return count > 0, nil
}

// ExpenseApproverIds retrieves all the records using an executor.
func ExpenseApproverIds(mods ...qm.QueryMod) expenseApproverIDQuery {
	mods = append(mods, qm.From("\"expense_approver_ids\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"expense_approver_ids\".*"})
	}

	return expenseApproverIDQuery{q}
}

// FindExpenseApproverID retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindExpenseApproverID(ctx context.Context, exec boil.ContextExecutor, expenseID string, approverID string, selectCols ...string) (*ExpenseApproverID, error) {
	expenseApproverIDObj := &ExpenseApproverID{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"expense_approver_ids\" where \"expense_id\"=$1 AND \"approver_id\"=$2", sel,
	)

	q := queries.Raw(query, expenseID, approverID)

	err := q.Bind(ctx, exec, expenseApproverIDObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from expense_approver_ids")
	}

	if err = expenseApproverIDObj.doAfterSelectHooks(ctx, exec); err != nil {
		return expenseApproverIDObj, err
	}

	return expenseApproverIDObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ExpenseApproverID) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no expense_approver_ids provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(expenseApproverIDColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	expenseApproverIDInsertCacheMut.RLock()
	cache, cached := expenseApproverIDInsertCache[key]
	expenseApproverIDInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			expenseApproverIDAllColumns,
			expenseApproverIDColumnsWithDefault,
			expenseApproverIDColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(expenseApproverIDType, expenseApproverIDMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(expenseApproverIDType, expenseApproverIDMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"expense_approver_ids\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"expense_approver_ids\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into expense_approver_ids")
	}

	if !cached {
		expenseApproverIDInsertCacheMut.Lock()
		expenseApproverIDInsertCache[key] = cache
		expenseApproverIDInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ExpenseApproverID.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ExpenseApproverID) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	expenseApproverIDUpdateCacheMut.RLock()
	cache, cached := expenseApproverIDUpdateCache[key]
	expenseApproverIDUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			expenseApproverIDAllColumns,
			expenseApproverIDPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update expense_approver_ids, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"expense_approver_ids\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, expenseApproverIDPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(expenseApproverIDType, expenseApproverIDMapping, append(wl, expenseApproverIDPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update expense_approver_ids row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for expense_approver_ids")
	}

	if !cached {
		expenseApproverIDUpdateCacheMut.Lock()
		expenseApproverIDUpdateCache[key] = cache
		expenseApproverIDUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q expenseApproverIDQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for expense_approver_ids")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for expense_approver_ids")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ExpenseApproverIDSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), expenseApproverIDPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"expense_approver_ids\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, expenseApproverIDPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in expenseApproverID slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all expenseApproverID")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ExpenseApproverID) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no expense_approver_ids provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(expenseApproverIDColumnsWithDefault, o)

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

	expenseApproverIDUpsertCacheMut.RLock()
	cache, cached := expenseApproverIDUpsertCache[key]
	expenseApproverIDUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			expenseApproverIDAllColumns,
			expenseApproverIDColumnsWithDefault,
			expenseApproverIDColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			expenseApproverIDAllColumns,
			expenseApproverIDPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert expense_approver_ids, could not build update column list")
		}

		ret := strmangle.SetComplement(expenseApproverIDAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(expenseApproverIDPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert expense_approver_ids, could not build conflict column list")
			}

			conflict = make([]string, len(expenseApproverIDPrimaryKeyColumns))
			copy(conflict, expenseApproverIDPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"expense_approver_ids\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(expenseApproverIDType, expenseApproverIDMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(expenseApproverIDType, expenseApproverIDMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert expense_approver_ids")
	}

	if !cached {
		expenseApproverIDUpsertCacheMut.Lock()
		expenseApproverIDUpsertCache[key] = cache
		expenseApproverIDUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ExpenseApproverID record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ExpenseApproverID) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ExpenseApproverID provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), expenseApproverIDPrimaryKeyMapping)
	sql := "DELETE FROM \"expense_approver_ids\" WHERE \"expense_id\"=$1 AND \"approver_id\"=$2"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from expense_approver_ids")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for expense_approver_ids")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q expenseApproverIDQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no expenseApproverIDQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from expense_approver_ids")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for expense_approver_ids")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ExpenseApproverIDSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(expenseApproverIDBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), expenseApproverIDPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"expense_approver_ids\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, expenseApproverIDPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from expenseApproverID slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for expense_approver_ids")
	}

	if len(expenseApproverIDAfterDeleteHooks) != 0 {
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
func (o *ExpenseApproverID) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindExpenseApproverID(ctx, exec, o.ExpenseID, o.ApproverID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ExpenseApproverIDSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ExpenseApproverIDSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), expenseApproverIDPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"expense_approver_ids\".* FROM \"expense_approver_ids\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, expenseApproverIDPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ExpenseApproverIDSlice")
	}

	*o = slice

	return nil
}

// ExpenseApproverIDExists checks if the ExpenseApproverID row exists.
func ExpenseApproverIDExists(ctx context.Context, exec boil.ContextExecutor, expenseID string, approverID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"expense_approver_ids\" where \"expense_id\"=$1 AND \"approver_id\"=$2 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, expenseID, approverID)
	}
	row := exec.QueryRowContext(ctx, sql, expenseID, approverID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if expense_approver_ids exists")
	}

	return exists, nil
}

// Exists checks if the ExpenseApproverID row exists.
func (o *ExpenseApproverID) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ExpenseApproverIDExists(ctx, exec, o.ExpenseID, o.ApproverID)
}
