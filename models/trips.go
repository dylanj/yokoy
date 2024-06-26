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

// Trip is an object representing the database table.
type Trip struct {
	ID            string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	Currency      null.String `boil:"currency" json:"currency,omitempty" toml:"currency" yaml:"currency,omitempty"`
	StartDateTime null.Time   `boil:"start_date_time" json:"start_date_time,omitempty" toml:"start_date_time" yaml:"start_date_time,omitempty"`
	EndDateTime   null.Time   `boil:"end_date_time" json:"end_date_time,omitempty" toml:"end_date_time" yaml:"end_date_time,omitempty"`
	LegalEntityID null.String `boil:"legal_entity_id" json:"legal_entity_id,omitempty" toml:"legal_entity_id" yaml:"legal_entity_id,omitempty"`
	Name          null.String `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	Status        null.String `boil:"status" json:"status,omitempty" toml:"status" yaml:"status,omitempty"`
	TotalClaim    null.Int    `boil:"total_claim" json:"total_claim,omitempty" toml:"total_claim" yaml:"total_claim,omitempty"`
	UserID        null.String `boil:"user_id" json:"user_id,omitempty" toml:"user_id" yaml:"user_id,omitempty"`

	R *tripR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L tripL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TripColumns = struct {
	ID            string
	Currency      string
	StartDateTime string
	EndDateTime   string
	LegalEntityID string
	Name          string
	Status        string
	TotalClaim    string
	UserID        string
}{
	ID:            "id",
	Currency:      "currency",
	StartDateTime: "start_date_time",
	EndDateTime:   "end_date_time",
	LegalEntityID: "legal_entity_id",
	Name:          "name",
	Status:        "status",
	TotalClaim:    "total_claim",
	UserID:        "user_id",
}

var TripTableColumns = struct {
	ID            string
	Currency      string
	StartDateTime string
	EndDateTime   string
	LegalEntityID string
	Name          string
	Status        string
	TotalClaim    string
	UserID        string
}{
	ID:            "trips.id",
	Currency:      "trips.currency",
	StartDateTime: "trips.start_date_time",
	EndDateTime:   "trips.end_date_time",
	LegalEntityID: "trips.legal_entity_id",
	Name:          "trips.name",
	Status:        "trips.status",
	TotalClaim:    "trips.total_claim",
	UserID:        "trips.user_id",
}

// Generated where

var TripWhere = struct {
	ID            whereHelperstring
	Currency      whereHelpernull_String
	StartDateTime whereHelpernull_Time
	EndDateTime   whereHelpernull_Time
	LegalEntityID whereHelpernull_String
	Name          whereHelpernull_String
	Status        whereHelpernull_String
	TotalClaim    whereHelpernull_Int
	UserID        whereHelpernull_String
}{
	ID:            whereHelperstring{field: "\"trips\".\"id\""},
	Currency:      whereHelpernull_String{field: "\"trips\".\"currency\""},
	StartDateTime: whereHelpernull_Time{field: "\"trips\".\"start_date_time\""},
	EndDateTime:   whereHelpernull_Time{field: "\"trips\".\"end_date_time\""},
	LegalEntityID: whereHelpernull_String{field: "\"trips\".\"legal_entity_id\""},
	Name:          whereHelpernull_String{field: "\"trips\".\"name\""},
	Status:        whereHelpernull_String{field: "\"trips\".\"status\""},
	TotalClaim:    whereHelpernull_Int{field: "\"trips\".\"total_claim\""},
	UserID:        whereHelpernull_String{field: "\"trips\".\"user_id\""},
}

// TripRels is where relationship names are stored.
var TripRels = struct {
}{}

// tripR is where relationships are stored.
type tripR struct {
}

// NewStruct creates a new relationship struct
func (*tripR) NewStruct() *tripR {
	return &tripR{}
}

// tripL is where Load methods for each relationship are stored.
type tripL struct{}

var (
	tripAllColumns            = []string{"id", "currency", "start_date_time", "end_date_time", "legal_entity_id", "name", "status", "total_claim", "user_id"}
	tripColumnsWithoutDefault = []string{"id"}
	tripColumnsWithDefault    = []string{"currency", "start_date_time", "end_date_time", "legal_entity_id", "name", "status", "total_claim", "user_id"}
	tripPrimaryKeyColumns     = []string{"id"}
	tripGeneratedColumns      = []string{}
)

type (
	// TripSlice is an alias for a slice of pointers to Trip.
	// This should almost always be used instead of []Trip.
	TripSlice []*Trip
	// TripHook is the signature for custom Trip hook methods
	TripHook func(context.Context, boil.ContextExecutor, *Trip) error

	tripQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	tripType                 = reflect.TypeOf(&Trip{})
	tripMapping              = queries.MakeStructMapping(tripType)
	tripPrimaryKeyMapping, _ = queries.BindMapping(tripType, tripMapping, tripPrimaryKeyColumns)
	tripInsertCacheMut       sync.RWMutex
	tripInsertCache          = make(map[string]insertCache)
	tripUpdateCacheMut       sync.RWMutex
	tripUpdateCache          = make(map[string]updateCache)
	tripUpsertCacheMut       sync.RWMutex
	tripUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var tripAfterSelectMu sync.Mutex
var tripAfterSelectHooks []TripHook

var tripBeforeInsertMu sync.Mutex
var tripBeforeInsertHooks []TripHook
var tripAfterInsertMu sync.Mutex
var tripAfterInsertHooks []TripHook

var tripBeforeUpdateMu sync.Mutex
var tripBeforeUpdateHooks []TripHook
var tripAfterUpdateMu sync.Mutex
var tripAfterUpdateHooks []TripHook

var tripBeforeDeleteMu sync.Mutex
var tripBeforeDeleteHooks []TripHook
var tripAfterDeleteMu sync.Mutex
var tripAfterDeleteHooks []TripHook

var tripBeforeUpsertMu sync.Mutex
var tripBeforeUpsertHooks []TripHook
var tripAfterUpsertMu sync.Mutex
var tripAfterUpsertHooks []TripHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Trip) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tripAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Trip) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tripBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Trip) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tripAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Trip) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tripBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Trip) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tripAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Trip) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tripBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Trip) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tripAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Trip) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tripBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Trip) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tripAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTripHook registers your hook function for all future operations.
func AddTripHook(hookPoint boil.HookPoint, tripHook TripHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		tripAfterSelectMu.Lock()
		tripAfterSelectHooks = append(tripAfterSelectHooks, tripHook)
		tripAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		tripBeforeInsertMu.Lock()
		tripBeforeInsertHooks = append(tripBeforeInsertHooks, tripHook)
		tripBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		tripAfterInsertMu.Lock()
		tripAfterInsertHooks = append(tripAfterInsertHooks, tripHook)
		tripAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		tripBeforeUpdateMu.Lock()
		tripBeforeUpdateHooks = append(tripBeforeUpdateHooks, tripHook)
		tripBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		tripAfterUpdateMu.Lock()
		tripAfterUpdateHooks = append(tripAfterUpdateHooks, tripHook)
		tripAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		tripBeforeDeleteMu.Lock()
		tripBeforeDeleteHooks = append(tripBeforeDeleteHooks, tripHook)
		tripBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		tripAfterDeleteMu.Lock()
		tripAfterDeleteHooks = append(tripAfterDeleteHooks, tripHook)
		tripAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		tripBeforeUpsertMu.Lock()
		tripBeforeUpsertHooks = append(tripBeforeUpsertHooks, tripHook)
		tripBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		tripAfterUpsertMu.Lock()
		tripAfterUpsertHooks = append(tripAfterUpsertHooks, tripHook)
		tripAfterUpsertMu.Unlock()
	}
}

// One returns a single trip record from the query.
func (q tripQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Trip, error) {
	o := &Trip{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for trips")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Trip records from the query.
func (q tripQuery) All(ctx context.Context, exec boil.ContextExecutor) (TripSlice, error) {
	var o []*Trip

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Trip slice")
	}

	if len(tripAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Trip records in the query.
func (q tripQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count trips rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q tripQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if trips exists")
	}

	return count > 0, nil
}

// Trips retrieves all the records using an executor.
func Trips(mods ...qm.QueryMod) tripQuery {
	mods = append(mods, qm.From("\"trips\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"trips\".*"})
	}

	return tripQuery{q}
}

// FindTrip retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTrip(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Trip, error) {
	tripObj := &Trip{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"trips\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, tripObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from trips")
	}

	if err = tripObj.doAfterSelectHooks(ctx, exec); err != nil {
		return tripObj, err
	}

	return tripObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Trip) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no trips provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(tripColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	tripInsertCacheMut.RLock()
	cache, cached := tripInsertCache[key]
	tripInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			tripAllColumns,
			tripColumnsWithDefault,
			tripColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(tripType, tripMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(tripType, tripMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"trips\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"trips\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into trips")
	}

	if !cached {
		tripInsertCacheMut.Lock()
		tripInsertCache[key] = cache
		tripInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Trip.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Trip) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	tripUpdateCacheMut.RLock()
	cache, cached := tripUpdateCache[key]
	tripUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			tripAllColumns,
			tripPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update trips, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"trips\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, tripPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(tripType, tripMapping, append(wl, tripPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update trips row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for trips")
	}

	if !cached {
		tripUpdateCacheMut.Lock()
		tripUpdateCache[key] = cache
		tripUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q tripQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for trips")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for trips")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TripSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tripPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"trips\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, tripPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in trip slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all trip")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Trip) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no trips provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(tripColumnsWithDefault, o)

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

	tripUpsertCacheMut.RLock()
	cache, cached := tripUpsertCache[key]
	tripUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			tripAllColumns,
			tripColumnsWithDefault,
			tripColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			tripAllColumns,
			tripPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert trips, could not build update column list")
		}

		ret := strmangle.SetComplement(tripAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(tripPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert trips, could not build conflict column list")
			}

			conflict = make([]string, len(tripPrimaryKeyColumns))
			copy(conflict, tripPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"trips\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(tripType, tripMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(tripType, tripMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert trips")
	}

	if !cached {
		tripUpsertCacheMut.Lock()
		tripUpsertCache[key] = cache
		tripUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Trip record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Trip) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Trip provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), tripPrimaryKeyMapping)
	sql := "DELETE FROM \"trips\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from trips")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for trips")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q tripQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no tripQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from trips")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for trips")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TripSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(tripBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tripPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"trips\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, tripPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from trip slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for trips")
	}

	if len(tripAfterDeleteHooks) != 0 {
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
func (o *Trip) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTrip(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TripSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TripSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tripPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"trips\".* FROM \"trips\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, tripPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TripSlice")
	}

	*o = slice

	return nil
}

// TripExists checks if the Trip row exists.
func TripExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"trips\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if trips exists")
	}

	return exists, nil
}

// Exists checks if the Trip row exists.
func (o *Trip) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return TripExists(ctx, exec, o.ID)
}
