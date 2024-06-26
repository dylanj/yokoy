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
	"github.com/shopspring/decimal"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// InvoiceLineItem is an object representing the database table.
type InvoiceLineItem struct {
	ID                  int             `boil:"id" json:"id" toml:"id" yaml:"id"`
	InvoiceID           null.String     `boil:"invoice_id" json:"invoice_id,omitempty" toml:"invoice_id" yaml:"invoice_id,omitempty"`
	CategoryID          null.String     `boil:"category_id" json:"category_id,omitempty" toml:"category_id" yaml:"category_id,omitempty"`
	CostObjectID        null.String     `boil:"cost_object_id" json:"cost_object_id,omitempty" toml:"cost_object_id" yaml:"cost_object_id,omitempty"`
	Description         null.String     `boil:"description" json:"description,omitempty" toml:"description" yaml:"description,omitempty"`
	Gross               decimal.Decimal `boil:"gross" json:"gross" toml:"gross" yaml:"gross"`
	ItemPrice           decimal.Decimal `boil:"item_price" json:"item_price" toml:"item_price" yaml:"item_price"`
	Net                 decimal.Decimal `boil:"net" json:"net" toml:"net" yaml:"net"`
	PurchaseOrderID     null.String     `boil:"purchase_order_id" json:"purchase_order_id,omitempty" toml:"purchase_order_id" yaml:"purchase_order_id,omitempty"`
	PurchaseOrderItemID null.String     `boil:"purchase_order_item_id" json:"purchase_order_item_id,omitempty" toml:"purchase_order_item_id" yaml:"purchase_order_item_id,omitempty"`
	Quantity            decimal.Decimal `boil:"quantity" json:"quantity" toml:"quantity" yaml:"quantity"`
	Tags                null.String     `boil:"tags" json:"tags,omitempty" toml:"tags" yaml:"tags,omitempty"`
	TaxRateID           null.String     `boil:"tax_rate_id" json:"tax_rate_id,omitempty" toml:"tax_rate_id" yaml:"tax_rate_id,omitempty"`
	Unit                null.String     `boil:"unit" json:"unit,omitempty" toml:"unit" yaml:"unit,omitempty"`

	R *invoiceLineItemR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L invoiceLineItemL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var InvoiceLineItemColumns = struct {
	ID                  string
	InvoiceID           string
	CategoryID          string
	CostObjectID        string
	Description         string
	Gross               string
	ItemPrice           string
	Net                 string
	PurchaseOrderID     string
	PurchaseOrderItemID string
	Quantity            string
	Tags                string
	TaxRateID           string
	Unit                string
}{
	ID:                  "id",
	InvoiceID:           "invoice_id",
	CategoryID:          "category_id",
	CostObjectID:        "cost_object_id",
	Description:         "description",
	Gross:               "gross",
	ItemPrice:           "item_price",
	Net:                 "net",
	PurchaseOrderID:     "purchase_order_id",
	PurchaseOrderItemID: "purchase_order_item_id",
	Quantity:            "quantity",
	Tags:                "tags",
	TaxRateID:           "tax_rate_id",
	Unit:                "unit",
}

var InvoiceLineItemTableColumns = struct {
	ID                  string
	InvoiceID           string
	CategoryID          string
	CostObjectID        string
	Description         string
	Gross               string
	ItemPrice           string
	Net                 string
	PurchaseOrderID     string
	PurchaseOrderItemID string
	Quantity            string
	Tags                string
	TaxRateID           string
	Unit                string
}{
	ID:                  "invoice_line_items.id",
	InvoiceID:           "invoice_line_items.invoice_id",
	CategoryID:          "invoice_line_items.category_id",
	CostObjectID:        "invoice_line_items.cost_object_id",
	Description:         "invoice_line_items.description",
	Gross:               "invoice_line_items.gross",
	ItemPrice:           "invoice_line_items.item_price",
	Net:                 "invoice_line_items.net",
	PurchaseOrderID:     "invoice_line_items.purchase_order_id",
	PurchaseOrderItemID: "invoice_line_items.purchase_order_item_id",
	Quantity:            "invoice_line_items.quantity",
	Tags:                "invoice_line_items.tags",
	TaxRateID:           "invoice_line_items.tax_rate_id",
	Unit:                "invoice_line_items.unit",
}

// Generated where

var InvoiceLineItemWhere = struct {
	ID                  whereHelperint
	InvoiceID           whereHelpernull_String
	CategoryID          whereHelpernull_String
	CostObjectID        whereHelpernull_String
	Description         whereHelpernull_String
	Gross               whereHelperdecimal_Decimal
	ItemPrice           whereHelperdecimal_Decimal
	Net                 whereHelperdecimal_Decimal
	PurchaseOrderID     whereHelpernull_String
	PurchaseOrderItemID whereHelpernull_String
	Quantity            whereHelperdecimal_Decimal
	Tags                whereHelpernull_String
	TaxRateID           whereHelpernull_String
	Unit                whereHelpernull_String
}{
	ID:                  whereHelperint{field: "\"invoice_line_items\".\"id\""},
	InvoiceID:           whereHelpernull_String{field: "\"invoice_line_items\".\"invoice_id\""},
	CategoryID:          whereHelpernull_String{field: "\"invoice_line_items\".\"category_id\""},
	CostObjectID:        whereHelpernull_String{field: "\"invoice_line_items\".\"cost_object_id\""},
	Description:         whereHelpernull_String{field: "\"invoice_line_items\".\"description\""},
	Gross:               whereHelperdecimal_Decimal{field: "\"invoice_line_items\".\"gross\""},
	ItemPrice:           whereHelperdecimal_Decimal{field: "\"invoice_line_items\".\"item_price\""},
	Net:                 whereHelperdecimal_Decimal{field: "\"invoice_line_items\".\"net\""},
	PurchaseOrderID:     whereHelpernull_String{field: "\"invoice_line_items\".\"purchase_order_id\""},
	PurchaseOrderItemID: whereHelpernull_String{field: "\"invoice_line_items\".\"purchase_order_item_id\""},
	Quantity:            whereHelperdecimal_Decimal{field: "\"invoice_line_items\".\"quantity\""},
	Tags:                whereHelpernull_String{field: "\"invoice_line_items\".\"tags\""},
	TaxRateID:           whereHelpernull_String{field: "\"invoice_line_items\".\"tax_rate_id\""},
	Unit:                whereHelpernull_String{field: "\"invoice_line_items\".\"unit\""},
}

// InvoiceLineItemRels is where relationship names are stored.
var InvoiceLineItemRels = struct {
}{}

// invoiceLineItemR is where relationships are stored.
type invoiceLineItemR struct {
}

// NewStruct creates a new relationship struct
func (*invoiceLineItemR) NewStruct() *invoiceLineItemR {
	return &invoiceLineItemR{}
}

// invoiceLineItemL is where Load methods for each relationship are stored.
type invoiceLineItemL struct{}

var (
	invoiceLineItemAllColumns            = []string{"id", "invoice_id", "category_id", "cost_object_id", "description", "gross", "item_price", "net", "purchase_order_id", "purchase_order_item_id", "quantity", "tags", "tax_rate_id", "unit"}
	invoiceLineItemColumnsWithoutDefault = []string{"gross", "item_price", "net", "quantity"}
	invoiceLineItemColumnsWithDefault    = []string{"id", "invoice_id", "category_id", "cost_object_id", "description", "purchase_order_id", "purchase_order_item_id", "tags", "tax_rate_id", "unit"}
	invoiceLineItemPrimaryKeyColumns     = []string{"id"}
	invoiceLineItemGeneratedColumns      = []string{}
)

type (
	// InvoiceLineItemSlice is an alias for a slice of pointers to InvoiceLineItem.
	// This should almost always be used instead of []InvoiceLineItem.
	InvoiceLineItemSlice []*InvoiceLineItem
	// InvoiceLineItemHook is the signature for custom InvoiceLineItem hook methods
	InvoiceLineItemHook func(context.Context, boil.ContextExecutor, *InvoiceLineItem) error

	invoiceLineItemQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	invoiceLineItemType                 = reflect.TypeOf(&InvoiceLineItem{})
	invoiceLineItemMapping              = queries.MakeStructMapping(invoiceLineItemType)
	invoiceLineItemPrimaryKeyMapping, _ = queries.BindMapping(invoiceLineItemType, invoiceLineItemMapping, invoiceLineItemPrimaryKeyColumns)
	invoiceLineItemInsertCacheMut       sync.RWMutex
	invoiceLineItemInsertCache          = make(map[string]insertCache)
	invoiceLineItemUpdateCacheMut       sync.RWMutex
	invoiceLineItemUpdateCache          = make(map[string]updateCache)
	invoiceLineItemUpsertCacheMut       sync.RWMutex
	invoiceLineItemUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var invoiceLineItemAfterSelectMu sync.Mutex
var invoiceLineItemAfterSelectHooks []InvoiceLineItemHook

var invoiceLineItemBeforeInsertMu sync.Mutex
var invoiceLineItemBeforeInsertHooks []InvoiceLineItemHook
var invoiceLineItemAfterInsertMu sync.Mutex
var invoiceLineItemAfterInsertHooks []InvoiceLineItemHook

var invoiceLineItemBeforeUpdateMu sync.Mutex
var invoiceLineItemBeforeUpdateHooks []InvoiceLineItemHook
var invoiceLineItemAfterUpdateMu sync.Mutex
var invoiceLineItemAfterUpdateHooks []InvoiceLineItemHook

var invoiceLineItemBeforeDeleteMu sync.Mutex
var invoiceLineItemBeforeDeleteHooks []InvoiceLineItemHook
var invoiceLineItemAfterDeleteMu sync.Mutex
var invoiceLineItemAfterDeleteHooks []InvoiceLineItemHook

var invoiceLineItemBeforeUpsertMu sync.Mutex
var invoiceLineItemBeforeUpsertHooks []InvoiceLineItemHook
var invoiceLineItemAfterUpsertMu sync.Mutex
var invoiceLineItemAfterUpsertHooks []InvoiceLineItemHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *InvoiceLineItem) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range invoiceLineItemAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *InvoiceLineItem) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range invoiceLineItemBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *InvoiceLineItem) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range invoiceLineItemAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *InvoiceLineItem) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range invoiceLineItemBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *InvoiceLineItem) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range invoiceLineItemAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *InvoiceLineItem) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range invoiceLineItemBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *InvoiceLineItem) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range invoiceLineItemAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *InvoiceLineItem) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range invoiceLineItemBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *InvoiceLineItem) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range invoiceLineItemAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddInvoiceLineItemHook registers your hook function for all future operations.
func AddInvoiceLineItemHook(hookPoint boil.HookPoint, invoiceLineItemHook InvoiceLineItemHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		invoiceLineItemAfterSelectMu.Lock()
		invoiceLineItemAfterSelectHooks = append(invoiceLineItemAfterSelectHooks, invoiceLineItemHook)
		invoiceLineItemAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		invoiceLineItemBeforeInsertMu.Lock()
		invoiceLineItemBeforeInsertHooks = append(invoiceLineItemBeforeInsertHooks, invoiceLineItemHook)
		invoiceLineItemBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		invoiceLineItemAfterInsertMu.Lock()
		invoiceLineItemAfterInsertHooks = append(invoiceLineItemAfterInsertHooks, invoiceLineItemHook)
		invoiceLineItemAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		invoiceLineItemBeforeUpdateMu.Lock()
		invoiceLineItemBeforeUpdateHooks = append(invoiceLineItemBeforeUpdateHooks, invoiceLineItemHook)
		invoiceLineItemBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		invoiceLineItemAfterUpdateMu.Lock()
		invoiceLineItemAfterUpdateHooks = append(invoiceLineItemAfterUpdateHooks, invoiceLineItemHook)
		invoiceLineItemAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		invoiceLineItemBeforeDeleteMu.Lock()
		invoiceLineItemBeforeDeleteHooks = append(invoiceLineItemBeforeDeleteHooks, invoiceLineItemHook)
		invoiceLineItemBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		invoiceLineItemAfterDeleteMu.Lock()
		invoiceLineItemAfterDeleteHooks = append(invoiceLineItemAfterDeleteHooks, invoiceLineItemHook)
		invoiceLineItemAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		invoiceLineItemBeforeUpsertMu.Lock()
		invoiceLineItemBeforeUpsertHooks = append(invoiceLineItemBeforeUpsertHooks, invoiceLineItemHook)
		invoiceLineItemBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		invoiceLineItemAfterUpsertMu.Lock()
		invoiceLineItemAfterUpsertHooks = append(invoiceLineItemAfterUpsertHooks, invoiceLineItemHook)
		invoiceLineItemAfterUpsertMu.Unlock()
	}
}

// One returns a single invoiceLineItem record from the query.
func (q invoiceLineItemQuery) One(ctx context.Context, exec boil.ContextExecutor) (*InvoiceLineItem, error) {
	o := &InvoiceLineItem{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for invoice_line_items")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all InvoiceLineItem records from the query.
func (q invoiceLineItemQuery) All(ctx context.Context, exec boil.ContextExecutor) (InvoiceLineItemSlice, error) {
	var o []*InvoiceLineItem

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to InvoiceLineItem slice")
	}

	if len(invoiceLineItemAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all InvoiceLineItem records in the query.
func (q invoiceLineItemQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count invoice_line_items rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q invoiceLineItemQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if invoice_line_items exists")
	}

	return count > 0, nil
}

// InvoiceLineItems retrieves all the records using an executor.
func InvoiceLineItems(mods ...qm.QueryMod) invoiceLineItemQuery {
	mods = append(mods, qm.From("\"invoice_line_items\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"invoice_line_items\".*"})
	}

	return invoiceLineItemQuery{q}
}

// FindInvoiceLineItem retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindInvoiceLineItem(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*InvoiceLineItem, error) {
	invoiceLineItemObj := &InvoiceLineItem{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"invoice_line_items\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, invoiceLineItemObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from invoice_line_items")
	}

	if err = invoiceLineItemObj.doAfterSelectHooks(ctx, exec); err != nil {
		return invoiceLineItemObj, err
	}

	return invoiceLineItemObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *InvoiceLineItem) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no invoice_line_items provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(invoiceLineItemColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	invoiceLineItemInsertCacheMut.RLock()
	cache, cached := invoiceLineItemInsertCache[key]
	invoiceLineItemInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			invoiceLineItemAllColumns,
			invoiceLineItemColumnsWithDefault,
			invoiceLineItemColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(invoiceLineItemType, invoiceLineItemMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(invoiceLineItemType, invoiceLineItemMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"invoice_line_items\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"invoice_line_items\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into invoice_line_items")
	}

	if !cached {
		invoiceLineItemInsertCacheMut.Lock()
		invoiceLineItemInsertCache[key] = cache
		invoiceLineItemInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the InvoiceLineItem.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *InvoiceLineItem) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	invoiceLineItemUpdateCacheMut.RLock()
	cache, cached := invoiceLineItemUpdateCache[key]
	invoiceLineItemUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			invoiceLineItemAllColumns,
			invoiceLineItemPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update invoice_line_items, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"invoice_line_items\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, invoiceLineItemPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(invoiceLineItemType, invoiceLineItemMapping, append(wl, invoiceLineItemPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update invoice_line_items row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for invoice_line_items")
	}

	if !cached {
		invoiceLineItemUpdateCacheMut.Lock()
		invoiceLineItemUpdateCache[key] = cache
		invoiceLineItemUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q invoiceLineItemQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for invoice_line_items")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for invoice_line_items")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o InvoiceLineItemSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), invoiceLineItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"invoice_line_items\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, invoiceLineItemPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in invoiceLineItem slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all invoiceLineItem")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *InvoiceLineItem) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no invoice_line_items provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(invoiceLineItemColumnsWithDefault, o)

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

	invoiceLineItemUpsertCacheMut.RLock()
	cache, cached := invoiceLineItemUpsertCache[key]
	invoiceLineItemUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			invoiceLineItemAllColumns,
			invoiceLineItemColumnsWithDefault,
			invoiceLineItemColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			invoiceLineItemAllColumns,
			invoiceLineItemPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert invoice_line_items, could not build update column list")
		}

		ret := strmangle.SetComplement(invoiceLineItemAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(invoiceLineItemPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert invoice_line_items, could not build conflict column list")
			}

			conflict = make([]string, len(invoiceLineItemPrimaryKeyColumns))
			copy(conflict, invoiceLineItemPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"invoice_line_items\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(invoiceLineItemType, invoiceLineItemMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(invoiceLineItemType, invoiceLineItemMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert invoice_line_items")
	}

	if !cached {
		invoiceLineItemUpsertCacheMut.Lock()
		invoiceLineItemUpsertCache[key] = cache
		invoiceLineItemUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single InvoiceLineItem record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *InvoiceLineItem) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no InvoiceLineItem provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), invoiceLineItemPrimaryKeyMapping)
	sql := "DELETE FROM \"invoice_line_items\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from invoice_line_items")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for invoice_line_items")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q invoiceLineItemQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no invoiceLineItemQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from invoice_line_items")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for invoice_line_items")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o InvoiceLineItemSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(invoiceLineItemBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), invoiceLineItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"invoice_line_items\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, invoiceLineItemPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from invoiceLineItem slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for invoice_line_items")
	}

	if len(invoiceLineItemAfterDeleteHooks) != 0 {
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
func (o *InvoiceLineItem) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindInvoiceLineItem(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *InvoiceLineItemSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := InvoiceLineItemSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), invoiceLineItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"invoice_line_items\".* FROM \"invoice_line_items\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, invoiceLineItemPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in InvoiceLineItemSlice")
	}

	*o = slice

	return nil
}

// InvoiceLineItemExists checks if the InvoiceLineItem row exists.
func InvoiceLineItemExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"invoice_line_items\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if invoice_line_items exists")
	}

	return exists, nil
}

// Exists checks if the InvoiceLineItem row exists.
func (o *InvoiceLineItem) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return InvoiceLineItemExists(ctx, exec, o.ID)
}
