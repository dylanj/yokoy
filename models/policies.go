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

// Policy is an object representing the database table.
type Policy struct {
	ID            string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	Code          null.String `boil:"code" json:"code,omitempty" toml:"code" yaml:"code,omitempty"`
	Name          null.String `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	LegalEntityID null.String `boil:"legal_entity_id" json:"legal_entity_id,omitempty" toml:"legal_entity_id" yaml:"legal_entity_id,omitempty"`
	StatusActive  null.Bool   `boil:"status_active" json:"status_active,omitempty" toml:"status_active" yaml:"status_active,omitempty"`

	R *policyR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L policyL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PolicyColumns = struct {
	ID            string
	Code          string
	Name          string
	LegalEntityID string
	StatusActive  string
}{
	ID:            "id",
	Code:          "code",
	Name:          "name",
	LegalEntityID: "legal_entity_id",
	StatusActive:  "status_active",
}

var PolicyTableColumns = struct {
	ID            string
	Code          string
	Name          string
	LegalEntityID string
	StatusActive  string
}{
	ID:            "policies.id",
	Code:          "policies.code",
	Name:          "policies.name",
	LegalEntityID: "policies.legal_entity_id",
	StatusActive:  "policies.status_active",
}

// Generated where

var PolicyWhere = struct {
	ID            whereHelperstring
	Code          whereHelpernull_String
	Name          whereHelpernull_String
	LegalEntityID whereHelpernull_String
	StatusActive  whereHelpernull_Bool
}{
	ID:            whereHelperstring{field: "\"policies\".\"id\""},
	Code:          whereHelpernull_String{field: "\"policies\".\"code\""},
	Name:          whereHelpernull_String{field: "\"policies\".\"name\""},
	LegalEntityID: whereHelpernull_String{field: "\"policies\".\"legal_entity_id\""},
	StatusActive:  whereHelpernull_Bool{field: "\"policies\".\"status_active\""},
}

// PolicyRels is where relationship names are stored.
var PolicyRels = struct {
}{}

// policyR is where relationships are stored.
type policyR struct {
}

// NewStruct creates a new relationship struct
func (*policyR) NewStruct() *policyR {
	return &policyR{}
}

// policyL is where Load methods for each relationship are stored.
type policyL struct{}

var (
	policyAllColumns            = []string{"id", "code", "name", "legal_entity_id", "status_active"}
	policyColumnsWithoutDefault = []string{"id"}
	policyColumnsWithDefault    = []string{"code", "name", "legal_entity_id", "status_active"}
	policyPrimaryKeyColumns     = []string{"id"}
	policyGeneratedColumns      = []string{}
)

type (
	// PolicySlice is an alias for a slice of pointers to Policy.
	// This should almost always be used instead of []Policy.
	PolicySlice []*Policy
	// PolicyHook is the signature for custom Policy hook methods
	PolicyHook func(context.Context, boil.ContextExecutor, *Policy) error

	policyQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	policyType                 = reflect.TypeOf(&Policy{})
	policyMapping              = queries.MakeStructMapping(policyType)
	policyPrimaryKeyMapping, _ = queries.BindMapping(policyType, policyMapping, policyPrimaryKeyColumns)
	policyInsertCacheMut       sync.RWMutex
	policyInsertCache          = make(map[string]insertCache)
	policyUpdateCacheMut       sync.RWMutex
	policyUpdateCache          = make(map[string]updateCache)
	policyUpsertCacheMut       sync.RWMutex
	policyUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var policyAfterSelectMu sync.Mutex
var policyAfterSelectHooks []PolicyHook

var policyBeforeInsertMu sync.Mutex
var policyBeforeInsertHooks []PolicyHook
var policyAfterInsertMu sync.Mutex
var policyAfterInsertHooks []PolicyHook

var policyBeforeUpdateMu sync.Mutex
var policyBeforeUpdateHooks []PolicyHook
var policyAfterUpdateMu sync.Mutex
var policyAfterUpdateHooks []PolicyHook

var policyBeforeDeleteMu sync.Mutex
var policyBeforeDeleteHooks []PolicyHook
var policyAfterDeleteMu sync.Mutex
var policyAfterDeleteHooks []PolicyHook

var policyBeforeUpsertMu sync.Mutex
var policyBeforeUpsertHooks []PolicyHook
var policyAfterUpsertMu sync.Mutex
var policyAfterUpsertHooks []PolicyHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Policy) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range policyAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Policy) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range policyBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Policy) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range policyAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Policy) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range policyBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Policy) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range policyAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Policy) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range policyBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Policy) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range policyAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Policy) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range policyBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Policy) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range policyAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPolicyHook registers your hook function for all future operations.
func AddPolicyHook(hookPoint boil.HookPoint, policyHook PolicyHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		policyAfterSelectMu.Lock()
		policyAfterSelectHooks = append(policyAfterSelectHooks, policyHook)
		policyAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		policyBeforeInsertMu.Lock()
		policyBeforeInsertHooks = append(policyBeforeInsertHooks, policyHook)
		policyBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		policyAfterInsertMu.Lock()
		policyAfterInsertHooks = append(policyAfterInsertHooks, policyHook)
		policyAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		policyBeforeUpdateMu.Lock()
		policyBeforeUpdateHooks = append(policyBeforeUpdateHooks, policyHook)
		policyBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		policyAfterUpdateMu.Lock()
		policyAfterUpdateHooks = append(policyAfterUpdateHooks, policyHook)
		policyAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		policyBeforeDeleteMu.Lock()
		policyBeforeDeleteHooks = append(policyBeforeDeleteHooks, policyHook)
		policyBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		policyAfterDeleteMu.Lock()
		policyAfterDeleteHooks = append(policyAfterDeleteHooks, policyHook)
		policyAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		policyBeforeUpsertMu.Lock()
		policyBeforeUpsertHooks = append(policyBeforeUpsertHooks, policyHook)
		policyBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		policyAfterUpsertMu.Lock()
		policyAfterUpsertHooks = append(policyAfterUpsertHooks, policyHook)
		policyAfterUpsertMu.Unlock()
	}
}

// One returns a single policy record from the query.
func (q policyQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Policy, error) {
	o := &Policy{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for policies")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Policy records from the query.
func (q policyQuery) All(ctx context.Context, exec boil.ContextExecutor) (PolicySlice, error) {
	var o []*Policy

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Policy slice")
	}

	if len(policyAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Policy records in the query.
func (q policyQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count policies rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q policyQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if policies exists")
	}

	return count > 0, nil
}

// Policies retrieves all the records using an executor.
func Policies(mods ...qm.QueryMod) policyQuery {
	mods = append(mods, qm.From("\"policies\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"policies\".*"})
	}

	return policyQuery{q}
}

// FindPolicy retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPolicy(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Policy, error) {
	policyObj := &Policy{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"policies\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, policyObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from policies")
	}

	if err = policyObj.doAfterSelectHooks(ctx, exec); err != nil {
		return policyObj, err
	}

	return policyObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Policy) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no policies provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(policyColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	policyInsertCacheMut.RLock()
	cache, cached := policyInsertCache[key]
	policyInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			policyAllColumns,
			policyColumnsWithDefault,
			policyColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(policyType, policyMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(policyType, policyMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"policies\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"policies\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into policies")
	}

	if !cached {
		policyInsertCacheMut.Lock()
		policyInsertCache[key] = cache
		policyInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Policy.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Policy) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	policyUpdateCacheMut.RLock()
	cache, cached := policyUpdateCache[key]
	policyUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			policyAllColumns,
			policyPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update policies, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"policies\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, policyPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(policyType, policyMapping, append(wl, policyPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update policies row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for policies")
	}

	if !cached {
		policyUpdateCacheMut.Lock()
		policyUpdateCache[key] = cache
		policyUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q policyQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for policies")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for policies")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PolicySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), policyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"policies\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, policyPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in policy slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all policy")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Policy) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no policies provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(policyColumnsWithDefault, o)

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

	policyUpsertCacheMut.RLock()
	cache, cached := policyUpsertCache[key]
	policyUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			policyAllColumns,
			policyColumnsWithDefault,
			policyColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			policyAllColumns,
			policyPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert policies, could not build update column list")
		}

		ret := strmangle.SetComplement(policyAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(policyPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert policies, could not build conflict column list")
			}

			conflict = make([]string, len(policyPrimaryKeyColumns))
			copy(conflict, policyPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"policies\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(policyType, policyMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(policyType, policyMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert policies")
	}

	if !cached {
		policyUpsertCacheMut.Lock()
		policyUpsertCache[key] = cache
		policyUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Policy record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Policy) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Policy provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), policyPrimaryKeyMapping)
	sql := "DELETE FROM \"policies\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from policies")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for policies")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q policyQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no policyQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from policies")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for policies")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PolicySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(policyBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), policyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"policies\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, policyPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from policy slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for policies")
	}

	if len(policyAfterDeleteHooks) != 0 {
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
func (o *Policy) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPolicy(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PolicySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PolicySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), policyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"policies\".* FROM \"policies\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, policyPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PolicySlice")
	}

	*o = slice

	return nil
}

// PolicyExists checks if the Policy row exists.
func PolicyExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"policies\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if policies exists")
	}

	return exists, nil
}

// Exists checks if the Policy row exists.
func (o *Policy) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return PolicyExists(ctx, exec, o.ID)
}
