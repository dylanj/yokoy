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

// CompanyCard is an object representing the database table.
type CompanyCard struct {
	ID               string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	AccountReference null.String `boil:"account_reference" json:"account_reference,omitempty" toml:"account_reference" yaml:"account_reference,omitempty"`
	Currency         null.String `boil:"currency" json:"currency,omitempty" toml:"currency" yaml:"currency,omitempty"`
	Description      null.String `boil:"description" json:"description,omitempty" toml:"description" yaml:"description,omitempty"`
	Name             null.String `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	Number           null.String `boil:"number" json:"number,omitempty" toml:"number" yaml:"number,omitempty"`
	OwnerID          null.String `boil:"owner_id" json:"owner_id,omitempty" toml:"owner_id" yaml:"owner_id,omitempty"`
	StatusActive     null.Bool   `boil:"status_active" json:"status_active,omitempty" toml:"status_active" yaml:"status_active,omitempty"`
	LegalEntityID    null.String `boil:"legal_entity_id" json:"legal_entity_id,omitempty" toml:"legal_entity_id" yaml:"legal_entity_id,omitempty"`

	R *companyCardR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L companyCardL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CompanyCardColumns = struct {
	ID               string
	AccountReference string
	Currency         string
	Description      string
	Name             string
	Number           string
	OwnerID          string
	StatusActive     string
	LegalEntityID    string
}{
	ID:               "id",
	AccountReference: "account_reference",
	Currency:         "currency",
	Description:      "description",
	Name:             "name",
	Number:           "number",
	OwnerID:          "owner_id",
	StatusActive:     "status_active",
	LegalEntityID:    "legal_entity_id",
}

var CompanyCardTableColumns = struct {
	ID               string
	AccountReference string
	Currency         string
	Description      string
	Name             string
	Number           string
	OwnerID          string
	StatusActive     string
	LegalEntityID    string
}{
	ID:               "company_cards.id",
	AccountReference: "company_cards.account_reference",
	Currency:         "company_cards.currency",
	Description:      "company_cards.description",
	Name:             "company_cards.name",
	Number:           "company_cards.number",
	OwnerID:          "company_cards.owner_id",
	StatusActive:     "company_cards.status_active",
	LegalEntityID:    "company_cards.legal_entity_id",
}

// Generated where

var CompanyCardWhere = struct {
	ID               whereHelperstring
	AccountReference whereHelpernull_String
	Currency         whereHelpernull_String
	Description      whereHelpernull_String
	Name             whereHelpernull_String
	Number           whereHelpernull_String
	OwnerID          whereHelpernull_String
	StatusActive     whereHelpernull_Bool
	LegalEntityID    whereHelpernull_String
}{
	ID:               whereHelperstring{field: "\"company_cards\".\"id\""},
	AccountReference: whereHelpernull_String{field: "\"company_cards\".\"account_reference\""},
	Currency:         whereHelpernull_String{field: "\"company_cards\".\"currency\""},
	Description:      whereHelpernull_String{field: "\"company_cards\".\"description\""},
	Name:             whereHelpernull_String{field: "\"company_cards\".\"name\""},
	Number:           whereHelpernull_String{field: "\"company_cards\".\"number\""},
	OwnerID:          whereHelpernull_String{field: "\"company_cards\".\"owner_id\""},
	StatusActive:     whereHelpernull_Bool{field: "\"company_cards\".\"status_active\""},
	LegalEntityID:    whereHelpernull_String{field: "\"company_cards\".\"legal_entity_id\""},
}

// CompanyCardRels is where relationship names are stored.
var CompanyCardRels = struct {
}{}

// companyCardR is where relationships are stored.
type companyCardR struct {
}

// NewStruct creates a new relationship struct
func (*companyCardR) NewStruct() *companyCardR {
	return &companyCardR{}
}

// companyCardL is where Load methods for each relationship are stored.
type companyCardL struct{}

var (
	companyCardAllColumns            = []string{"id", "account_reference", "currency", "description", "name", "number", "owner_id", "status_active", "legal_entity_id"}
	companyCardColumnsWithoutDefault = []string{"id"}
	companyCardColumnsWithDefault    = []string{"account_reference", "currency", "description", "name", "number", "owner_id", "status_active", "legal_entity_id"}
	companyCardPrimaryKeyColumns     = []string{"id"}
	companyCardGeneratedColumns      = []string{}
)

type (
	// CompanyCardSlice is an alias for a slice of pointers to CompanyCard.
	// This should almost always be used instead of []CompanyCard.
	CompanyCardSlice []*CompanyCard
	// CompanyCardHook is the signature for custom CompanyCard hook methods
	CompanyCardHook func(context.Context, boil.ContextExecutor, *CompanyCard) error

	companyCardQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	companyCardType                 = reflect.TypeOf(&CompanyCard{})
	companyCardMapping              = queries.MakeStructMapping(companyCardType)
	companyCardPrimaryKeyMapping, _ = queries.BindMapping(companyCardType, companyCardMapping, companyCardPrimaryKeyColumns)
	companyCardInsertCacheMut       sync.RWMutex
	companyCardInsertCache          = make(map[string]insertCache)
	companyCardUpdateCacheMut       sync.RWMutex
	companyCardUpdateCache          = make(map[string]updateCache)
	companyCardUpsertCacheMut       sync.RWMutex
	companyCardUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var companyCardAfterSelectHooks []CompanyCardHook

var companyCardBeforeInsertHooks []CompanyCardHook
var companyCardAfterInsertHooks []CompanyCardHook

var companyCardBeforeUpdateHooks []CompanyCardHook
var companyCardAfterUpdateHooks []CompanyCardHook

var companyCardBeforeDeleteHooks []CompanyCardHook
var companyCardAfterDeleteHooks []CompanyCardHook

var companyCardBeforeUpsertHooks []CompanyCardHook
var companyCardAfterUpsertHooks []CompanyCardHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CompanyCard) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range companyCardAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CompanyCard) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range companyCardBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CompanyCard) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range companyCardAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CompanyCard) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range companyCardBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CompanyCard) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range companyCardAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CompanyCard) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range companyCardBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CompanyCard) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range companyCardAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CompanyCard) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range companyCardBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CompanyCard) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range companyCardAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCompanyCardHook registers your hook function for all future operations.
func AddCompanyCardHook(hookPoint boil.HookPoint, companyCardHook CompanyCardHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		companyCardAfterSelectHooks = append(companyCardAfterSelectHooks, companyCardHook)
	case boil.BeforeInsertHook:
		companyCardBeforeInsertHooks = append(companyCardBeforeInsertHooks, companyCardHook)
	case boil.AfterInsertHook:
		companyCardAfterInsertHooks = append(companyCardAfterInsertHooks, companyCardHook)
	case boil.BeforeUpdateHook:
		companyCardBeforeUpdateHooks = append(companyCardBeforeUpdateHooks, companyCardHook)
	case boil.AfterUpdateHook:
		companyCardAfterUpdateHooks = append(companyCardAfterUpdateHooks, companyCardHook)
	case boil.BeforeDeleteHook:
		companyCardBeforeDeleteHooks = append(companyCardBeforeDeleteHooks, companyCardHook)
	case boil.AfterDeleteHook:
		companyCardAfterDeleteHooks = append(companyCardAfterDeleteHooks, companyCardHook)
	case boil.BeforeUpsertHook:
		companyCardBeforeUpsertHooks = append(companyCardBeforeUpsertHooks, companyCardHook)
	case boil.AfterUpsertHook:
		companyCardAfterUpsertHooks = append(companyCardAfterUpsertHooks, companyCardHook)
	}
}

// One returns a single companyCard record from the query.
func (q companyCardQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CompanyCard, error) {
	o := &CompanyCard{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for company_cards")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CompanyCard records from the query.
func (q companyCardQuery) All(ctx context.Context, exec boil.ContextExecutor) (CompanyCardSlice, error) {
	var o []*CompanyCard

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CompanyCard slice")
	}

	if len(companyCardAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CompanyCard records in the query.
func (q companyCardQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count company_cards rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q companyCardQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if company_cards exists")
	}

	return count > 0, nil
}

// CompanyCards retrieves all the records using an executor.
func CompanyCards(mods ...qm.QueryMod) companyCardQuery {
	mods = append(mods, qm.From("\"company_cards\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"company_cards\".*"})
	}

	return companyCardQuery{q}
}

// FindCompanyCard retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCompanyCard(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*CompanyCard, error) {
	companyCardObj := &CompanyCard{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"company_cards\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, companyCardObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from company_cards")
	}

	if err = companyCardObj.doAfterSelectHooks(ctx, exec); err != nil {
		return companyCardObj, err
	}

	return companyCardObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CompanyCard) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no company_cards provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(companyCardColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	companyCardInsertCacheMut.RLock()
	cache, cached := companyCardInsertCache[key]
	companyCardInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			companyCardAllColumns,
			companyCardColumnsWithDefault,
			companyCardColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(companyCardType, companyCardMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(companyCardType, companyCardMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"company_cards\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"company_cards\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into company_cards")
	}

	if !cached {
		companyCardInsertCacheMut.Lock()
		companyCardInsertCache[key] = cache
		companyCardInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CompanyCard.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CompanyCard) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	companyCardUpdateCacheMut.RLock()
	cache, cached := companyCardUpdateCache[key]
	companyCardUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			companyCardAllColumns,
			companyCardPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update company_cards, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"company_cards\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, companyCardPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(companyCardType, companyCardMapping, append(wl, companyCardPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update company_cards row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for company_cards")
	}

	if !cached {
		companyCardUpdateCacheMut.Lock()
		companyCardUpdateCache[key] = cache
		companyCardUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q companyCardQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for company_cards")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for company_cards")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CompanyCardSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), companyCardPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"company_cards\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, companyCardPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in companyCard slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all companyCard")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CompanyCard) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no company_cards provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(companyCardColumnsWithDefault, o)

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

	companyCardUpsertCacheMut.RLock()
	cache, cached := companyCardUpsertCache[key]
	companyCardUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			companyCardAllColumns,
			companyCardColumnsWithDefault,
			companyCardColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			companyCardAllColumns,
			companyCardPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert company_cards, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(companyCardPrimaryKeyColumns))
			copy(conflict, companyCardPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"company_cards\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(companyCardType, companyCardMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(companyCardType, companyCardMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert company_cards")
	}

	if !cached {
		companyCardUpsertCacheMut.Lock()
		companyCardUpsertCache[key] = cache
		companyCardUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CompanyCard record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CompanyCard) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CompanyCard provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), companyCardPrimaryKeyMapping)
	sql := "DELETE FROM \"company_cards\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from company_cards")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for company_cards")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q companyCardQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no companyCardQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from company_cards")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for company_cards")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CompanyCardSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(companyCardBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), companyCardPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"company_cards\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, companyCardPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from companyCard slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for company_cards")
	}

	if len(companyCardAfterDeleteHooks) != 0 {
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
func (o *CompanyCard) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCompanyCard(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CompanyCardSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CompanyCardSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), companyCardPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"company_cards\".* FROM \"company_cards\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, companyCardPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CompanyCardSlice")
	}

	*o = slice

	return nil
}

// CompanyCardExists checks if the CompanyCard row exists.
func CompanyCardExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"company_cards\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if company_cards exists")
	}

	return exists, nil
}

// Exists checks if the CompanyCard row exists.
func (o *CompanyCard) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return CompanyCardExists(ctx, exec, o.ID)
}
