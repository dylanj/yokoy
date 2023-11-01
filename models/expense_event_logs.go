// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// ExpenseEventLog is an object representing the database table.
type ExpenseEventLog struct {
	ID         int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	ExpenseID  null.String `boil:"expense_id" json:"expense_id,omitempty" toml:"expense_id" yaml:"expense_id,omitempty"`
	ActionType null.String `boil:"action_type" json:"action_type,omitempty" toml:"action_type" yaml:"action_type,omitempty"`
	Comment    null.String `boil:"comment" json:"comment,omitempty" toml:"comment" yaml:"comment,omitempty"`
	Name       null.String `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	Timestamp  null.Time   `boil:"timestamp" json:"timestamp,omitempty" toml:"timestamp" yaml:"timestamp,omitempty"`
	UserID     null.String `boil:"user_id" json:"user_id,omitempty" toml:"user_id" yaml:"user_id,omitempty"`

	R *expenseEventLogR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L expenseEventLogL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ExpenseEventLogColumns = struct {
	ID         string
	ExpenseID  string
	ActionType string
	Comment    string
	Name       string
	Timestamp  string
	UserID     string
}{
	ID:         "id",
	ExpenseID:  "expense_id",
	ActionType: "action_type",
	Comment:    "comment",
	Name:       "name",
	Timestamp:  "timestamp",
	UserID:     "user_id",
}

var ExpenseEventLogTableColumns = struct {
	ID         string
	ExpenseID  string
	ActionType string
	Comment    string
	Name       string
	Timestamp  string
	UserID     string
}{
	ID:         "expense_event_logs.id",
	ExpenseID:  "expense_event_logs.expense_id",
	ActionType: "expense_event_logs.action_type",
	Comment:    "expense_event_logs.comment",
	Name:       "expense_event_logs.name",
	Timestamp:  "expense_event_logs.timestamp",
	UserID:     "expense_event_logs.user_id",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var ExpenseEventLogWhere = struct {
	ID         whereHelperint
	ExpenseID  whereHelpernull_String
	ActionType whereHelpernull_String
	Comment    whereHelpernull_String
	Name       whereHelpernull_String
	Timestamp  whereHelpernull_Time
	UserID     whereHelpernull_String
}{
	ID:         whereHelperint{field: "\"expense_event_logs\".\"id\""},
	ExpenseID:  whereHelpernull_String{field: "\"expense_event_logs\".\"expense_id\""},
	ActionType: whereHelpernull_String{field: "\"expense_event_logs\".\"action_type\""},
	Comment:    whereHelpernull_String{field: "\"expense_event_logs\".\"comment\""},
	Name:       whereHelpernull_String{field: "\"expense_event_logs\".\"name\""},
	Timestamp:  whereHelpernull_Time{field: "\"expense_event_logs\".\"timestamp\""},
	UserID:     whereHelpernull_String{field: "\"expense_event_logs\".\"user_id\""},
}

// ExpenseEventLogRels is where relationship names are stored.
var ExpenseEventLogRels = struct {
}{}

// expenseEventLogR is where relationships are stored.
type expenseEventLogR struct {
}

// NewStruct creates a new relationship struct
func (*expenseEventLogR) NewStruct() *expenseEventLogR {
	return &expenseEventLogR{}
}

// expenseEventLogL is where Load methods for each relationship are stored.
type expenseEventLogL struct{}

var (
	expenseEventLogAllColumns            = []string{"id", "expense_id", "action_type", "comment", "name", "timestamp", "user_id"}
	expenseEventLogColumnsWithoutDefault = []string{}
	expenseEventLogColumnsWithDefault    = []string{"id", "expense_id", "action_type", "comment", "name", "timestamp", "user_id"}
	expenseEventLogPrimaryKeyColumns     = []string{"id"}
	expenseEventLogGeneratedColumns      = []string{}
)

type (
	// ExpenseEventLogSlice is an alias for a slice of pointers to ExpenseEventLog.
	// This should almost always be used instead of []ExpenseEventLog.
	ExpenseEventLogSlice []*ExpenseEventLog
	// ExpenseEventLogHook is the signature for custom ExpenseEventLog hook methods
	ExpenseEventLogHook func(context.Context, boil.ContextExecutor, *ExpenseEventLog) error

	expenseEventLogQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	expenseEventLogType                 = reflect.TypeOf(&ExpenseEventLog{})
	expenseEventLogMapping              = queries.MakeStructMapping(expenseEventLogType)
	expenseEventLogPrimaryKeyMapping, _ = queries.BindMapping(expenseEventLogType, expenseEventLogMapping, expenseEventLogPrimaryKeyColumns)
	expenseEventLogInsertCacheMut       sync.RWMutex
	expenseEventLogInsertCache          = make(map[string]insertCache)
	expenseEventLogUpdateCacheMut       sync.RWMutex
	expenseEventLogUpdateCache          = make(map[string]updateCache)
	expenseEventLogUpsertCacheMut       sync.RWMutex
	expenseEventLogUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var expenseEventLogAfterSelectHooks []ExpenseEventLogHook

var expenseEventLogBeforeInsertHooks []ExpenseEventLogHook
var expenseEventLogAfterInsertHooks []ExpenseEventLogHook

var expenseEventLogBeforeUpdateHooks []ExpenseEventLogHook
var expenseEventLogAfterUpdateHooks []ExpenseEventLogHook

var expenseEventLogBeforeDeleteHooks []ExpenseEventLogHook
var expenseEventLogAfterDeleteHooks []ExpenseEventLogHook

var expenseEventLogBeforeUpsertHooks []ExpenseEventLogHook
var expenseEventLogAfterUpsertHooks []ExpenseEventLogHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ExpenseEventLog) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range expenseEventLogAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ExpenseEventLog) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range expenseEventLogBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ExpenseEventLog) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range expenseEventLogAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ExpenseEventLog) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range expenseEventLogBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ExpenseEventLog) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range expenseEventLogAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ExpenseEventLog) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range expenseEventLogBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ExpenseEventLog) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range expenseEventLogAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ExpenseEventLog) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range expenseEventLogBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ExpenseEventLog) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range expenseEventLogAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddExpenseEventLogHook registers your hook function for all future operations.
func AddExpenseEventLogHook(hookPoint boil.HookPoint, expenseEventLogHook ExpenseEventLogHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		expenseEventLogAfterSelectHooks = append(expenseEventLogAfterSelectHooks, expenseEventLogHook)
	case boil.BeforeInsertHook:
		expenseEventLogBeforeInsertHooks = append(expenseEventLogBeforeInsertHooks, expenseEventLogHook)
	case boil.AfterInsertHook:
		expenseEventLogAfterInsertHooks = append(expenseEventLogAfterInsertHooks, expenseEventLogHook)
	case boil.BeforeUpdateHook:
		expenseEventLogBeforeUpdateHooks = append(expenseEventLogBeforeUpdateHooks, expenseEventLogHook)
	case boil.AfterUpdateHook:
		expenseEventLogAfterUpdateHooks = append(expenseEventLogAfterUpdateHooks, expenseEventLogHook)
	case boil.BeforeDeleteHook:
		expenseEventLogBeforeDeleteHooks = append(expenseEventLogBeforeDeleteHooks, expenseEventLogHook)
	case boil.AfterDeleteHook:
		expenseEventLogAfterDeleteHooks = append(expenseEventLogAfterDeleteHooks, expenseEventLogHook)
	case boil.BeforeUpsertHook:
		expenseEventLogBeforeUpsertHooks = append(expenseEventLogBeforeUpsertHooks, expenseEventLogHook)
	case boil.AfterUpsertHook:
		expenseEventLogAfterUpsertHooks = append(expenseEventLogAfterUpsertHooks, expenseEventLogHook)
	}
}

// One returns a single expenseEventLog record from the query.
func (q expenseEventLogQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ExpenseEventLog, error) {
	o := &ExpenseEventLog{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for expense_event_logs")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ExpenseEventLog records from the query.
func (q expenseEventLogQuery) All(ctx context.Context, exec boil.ContextExecutor) (ExpenseEventLogSlice, error) {
	var o []*ExpenseEventLog

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ExpenseEventLog slice")
	}

	if len(expenseEventLogAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ExpenseEventLog records in the query.
func (q expenseEventLogQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count expense_event_logs rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q expenseEventLogQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if expense_event_logs exists")
	}

	return count > 0, nil
}

// ExpenseEventLogs retrieves all the records using an executor.
func ExpenseEventLogs(mods ...qm.QueryMod) expenseEventLogQuery {
	mods = append(mods, qm.From("\"expense_event_logs\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"expense_event_logs\".*"})
	}

	return expenseEventLogQuery{q}
}

// FindExpenseEventLog retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindExpenseEventLog(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*ExpenseEventLog, error) {
	expenseEventLogObj := &ExpenseEventLog{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"expense_event_logs\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, expenseEventLogObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from expense_event_logs")
	}

	if err = expenseEventLogObj.doAfterSelectHooks(ctx, exec); err != nil {
		return expenseEventLogObj, err
	}

	return expenseEventLogObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ExpenseEventLog) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no expense_event_logs provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(expenseEventLogColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	expenseEventLogInsertCacheMut.RLock()
	cache, cached := expenseEventLogInsertCache[key]
	expenseEventLogInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			expenseEventLogAllColumns,
			expenseEventLogColumnsWithDefault,
			expenseEventLogColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(expenseEventLogType, expenseEventLogMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(expenseEventLogType, expenseEventLogMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"expense_event_logs\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"expense_event_logs\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into expense_event_logs")
	}

	if !cached {
		expenseEventLogInsertCacheMut.Lock()
		expenseEventLogInsertCache[key] = cache
		expenseEventLogInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ExpenseEventLog.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ExpenseEventLog) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	expenseEventLogUpdateCacheMut.RLock()
	cache, cached := expenseEventLogUpdateCache[key]
	expenseEventLogUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			expenseEventLogAllColumns,
			expenseEventLogPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update expense_event_logs, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"expense_event_logs\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, expenseEventLogPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(expenseEventLogType, expenseEventLogMapping, append(wl, expenseEventLogPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update expense_event_logs row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for expense_event_logs")
	}

	if !cached {
		expenseEventLogUpdateCacheMut.Lock()
		expenseEventLogUpdateCache[key] = cache
		expenseEventLogUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q expenseEventLogQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for expense_event_logs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for expense_event_logs")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ExpenseEventLogSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), expenseEventLogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"expense_event_logs\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, expenseEventLogPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in expenseEventLog slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all expenseEventLog")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ExpenseEventLog) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no expense_event_logs provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(expenseEventLogColumnsWithDefault, o)

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

	expenseEventLogUpsertCacheMut.RLock()
	cache, cached := expenseEventLogUpsertCache[key]
	expenseEventLogUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			expenseEventLogAllColumns,
			expenseEventLogColumnsWithDefault,
			expenseEventLogColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			expenseEventLogAllColumns,
			expenseEventLogPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert expense_event_logs, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(expenseEventLogPrimaryKeyColumns))
			copy(conflict, expenseEventLogPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"expense_event_logs\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(expenseEventLogType, expenseEventLogMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(expenseEventLogType, expenseEventLogMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert expense_event_logs")
	}

	if !cached {
		expenseEventLogUpsertCacheMut.Lock()
		expenseEventLogUpsertCache[key] = cache
		expenseEventLogUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ExpenseEventLog record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ExpenseEventLog) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ExpenseEventLog provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), expenseEventLogPrimaryKeyMapping)
	sql := "DELETE FROM \"expense_event_logs\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from expense_event_logs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for expense_event_logs")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q expenseEventLogQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no expenseEventLogQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from expense_event_logs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for expense_event_logs")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ExpenseEventLogSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(expenseEventLogBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), expenseEventLogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"expense_event_logs\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, expenseEventLogPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from expenseEventLog slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for expense_event_logs")
	}

	if len(expenseEventLogAfterDeleteHooks) != 0 {
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
func (o *ExpenseEventLog) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindExpenseEventLog(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ExpenseEventLogSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ExpenseEventLogSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), expenseEventLogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"expense_event_logs\".* FROM \"expense_event_logs\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, expenseEventLogPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ExpenseEventLogSlice")
	}

	*o = slice

	return nil
}

// ExpenseEventLogExists checks if the ExpenseEventLog row exists.
func ExpenseEventLogExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"expense_event_logs\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if expense_event_logs exists")
	}

	return exists, nil
}

// Exists checks if the ExpenseEventLog row exists.
func (o *ExpenseEventLog) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ExpenseEventLogExists(ctx, exec, o.ID)
}