// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// ThinkingTree is an object representing the database table.
type ThinkingTree struct {
	ID           int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserID       string    `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	Theme        string    `boil:"theme" json:"theme" toml:"theme" yaml:"theme"`
	ThinkingTree string    `boil:"thinking_tree" json:"thinking_tree" toml:"thinking_tree" yaml:"thinking_tree"`
	InsertDate   time.Time `boil:"insert_date" json:"insert_date" toml:"insert_date" yaml:"insert_date"`

	R *thinkingTreeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L thinkingTreeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ThinkingTreeColumns = struct {
	ID           string
	UserID       string
	Theme        string
	ThinkingTree string
	InsertDate   string
}{
	ID:           "id",
	UserID:       "user_id",
	Theme:        "theme",
	ThinkingTree: "thinking_tree",
	InsertDate:   "insert_date",
}

var ThinkingTreeTableColumns = struct {
	ID           string
	UserID       string
	Theme        string
	ThinkingTree string
	InsertDate   string
}{
	ID:           "thinking_tree.id",
	UserID:       "thinking_tree.user_id",
	Theme:        "thinking_tree.theme",
	ThinkingTree: "thinking_tree.thinking_tree",
	InsertDate:   "thinking_tree.insert_date",
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

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var ThinkingTreeWhere = struct {
	ID           whereHelperint
	UserID       whereHelperstring
	Theme        whereHelperstring
	ThinkingTree whereHelperstring
	InsertDate   whereHelpertime_Time
}{
	ID:           whereHelperint{field: "`thinking_tree`.`id`"},
	UserID:       whereHelperstring{field: "`thinking_tree`.`user_id`"},
	Theme:        whereHelperstring{field: "`thinking_tree`.`theme`"},
	ThinkingTree: whereHelperstring{field: "`thinking_tree`.`thinking_tree`"},
	InsertDate:   whereHelpertime_Time{field: "`thinking_tree`.`insert_date`"},
}

// ThinkingTreeRels is where relationship names are stored.
var ThinkingTreeRels = struct {
}{}

// thinkingTreeR is where relationships are stored.
type thinkingTreeR struct {
}

// NewStruct creates a new relationship struct
func (*thinkingTreeR) NewStruct() *thinkingTreeR {
	return &thinkingTreeR{}
}

// thinkingTreeL is where Load methods for each relationship are stored.
type thinkingTreeL struct{}

var (
	thinkingTreeAllColumns            = []string{"id", "user_id", "theme", "thinking_tree", "insert_date"}
	thinkingTreeColumnsWithoutDefault = []string{"user_id", "theme", "thinking_tree", "insert_date"}
	thinkingTreeColumnsWithDefault    = []string{"id"}
	thinkingTreePrimaryKeyColumns     = []string{"id"}
	thinkingTreeGeneratedColumns      = []string{}
)

type (
	// ThinkingTreeSlice is an alias for a slice of pointers to ThinkingTree.
	// This should almost always be used instead of []ThinkingTree.
	ThinkingTreeSlice []*ThinkingTree
	// ThinkingTreeHook is the signature for custom ThinkingTree hook methods
	ThinkingTreeHook func(context.Context, boil.ContextExecutor, *ThinkingTree) error

	thinkingTreeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	thinkingTreeType                 = reflect.TypeOf(&ThinkingTree{})
	thinkingTreeMapping              = queries.MakeStructMapping(thinkingTreeType)
	thinkingTreePrimaryKeyMapping, _ = queries.BindMapping(thinkingTreeType, thinkingTreeMapping, thinkingTreePrimaryKeyColumns)
	thinkingTreeInsertCacheMut       sync.RWMutex
	thinkingTreeInsertCache          = make(map[string]insertCache)
	thinkingTreeUpdateCacheMut       sync.RWMutex
	thinkingTreeUpdateCache          = make(map[string]updateCache)
	thinkingTreeUpsertCacheMut       sync.RWMutex
	thinkingTreeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var thinkingTreeAfterSelectHooks []ThinkingTreeHook

var thinkingTreeBeforeInsertHooks []ThinkingTreeHook
var thinkingTreeAfterInsertHooks []ThinkingTreeHook

var thinkingTreeBeforeUpdateHooks []ThinkingTreeHook
var thinkingTreeAfterUpdateHooks []ThinkingTreeHook

var thinkingTreeBeforeDeleteHooks []ThinkingTreeHook
var thinkingTreeAfterDeleteHooks []ThinkingTreeHook

var thinkingTreeBeforeUpsertHooks []ThinkingTreeHook
var thinkingTreeAfterUpsertHooks []ThinkingTreeHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ThinkingTree) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thinkingTreeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ThinkingTree) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thinkingTreeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ThinkingTree) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thinkingTreeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ThinkingTree) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thinkingTreeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ThinkingTree) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thinkingTreeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ThinkingTree) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thinkingTreeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ThinkingTree) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thinkingTreeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ThinkingTree) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thinkingTreeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ThinkingTree) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thinkingTreeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddThinkingTreeHook registers your hook function for all future operations.
func AddThinkingTreeHook(hookPoint boil.HookPoint, thinkingTreeHook ThinkingTreeHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		thinkingTreeAfterSelectHooks = append(thinkingTreeAfterSelectHooks, thinkingTreeHook)
	case boil.BeforeInsertHook:
		thinkingTreeBeforeInsertHooks = append(thinkingTreeBeforeInsertHooks, thinkingTreeHook)
	case boil.AfterInsertHook:
		thinkingTreeAfterInsertHooks = append(thinkingTreeAfterInsertHooks, thinkingTreeHook)
	case boil.BeforeUpdateHook:
		thinkingTreeBeforeUpdateHooks = append(thinkingTreeBeforeUpdateHooks, thinkingTreeHook)
	case boil.AfterUpdateHook:
		thinkingTreeAfterUpdateHooks = append(thinkingTreeAfterUpdateHooks, thinkingTreeHook)
	case boil.BeforeDeleteHook:
		thinkingTreeBeforeDeleteHooks = append(thinkingTreeBeforeDeleteHooks, thinkingTreeHook)
	case boil.AfterDeleteHook:
		thinkingTreeAfterDeleteHooks = append(thinkingTreeAfterDeleteHooks, thinkingTreeHook)
	case boil.BeforeUpsertHook:
		thinkingTreeBeforeUpsertHooks = append(thinkingTreeBeforeUpsertHooks, thinkingTreeHook)
	case boil.AfterUpsertHook:
		thinkingTreeAfterUpsertHooks = append(thinkingTreeAfterUpsertHooks, thinkingTreeHook)
	}
}

// One returns a single thinkingTree record from the query.
func (q thinkingTreeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ThinkingTree, error) {
	o := &ThinkingTree{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for thinking_tree")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ThinkingTree records from the query.
func (q thinkingTreeQuery) All(ctx context.Context, exec boil.ContextExecutor) (ThinkingTreeSlice, error) {
	var o []*ThinkingTree

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ThinkingTree slice")
	}

	if len(thinkingTreeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ThinkingTree records in the query.
func (q thinkingTreeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count thinking_tree rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q thinkingTreeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if thinking_tree exists")
	}

	return count > 0, nil
}

// ThinkingTrees retrieves all the records using an executor.
func ThinkingTrees(mods ...qm.QueryMod) thinkingTreeQuery {
	mods = append(mods, qm.From("`thinking_tree`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`thinking_tree`.*"})
	}

	return thinkingTreeQuery{q}
}

// FindThinkingTree retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindThinkingTree(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*ThinkingTree, error) {
	thinkingTreeObj := &ThinkingTree{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `thinking_tree` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, thinkingTreeObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from thinking_tree")
	}

	if err = thinkingTreeObj.doAfterSelectHooks(ctx, exec); err != nil {
		return thinkingTreeObj, err
	}

	return thinkingTreeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ThinkingTree) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no thinking_tree provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(thinkingTreeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	thinkingTreeInsertCacheMut.RLock()
	cache, cached := thinkingTreeInsertCache[key]
	thinkingTreeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			thinkingTreeAllColumns,
			thinkingTreeColumnsWithDefault,
			thinkingTreeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(thinkingTreeType, thinkingTreeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(thinkingTreeType, thinkingTreeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `thinking_tree` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `thinking_tree` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `thinking_tree` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, thinkingTreePrimaryKeyColumns))
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into thinking_tree")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == thinkingTreeMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for thinking_tree")
	}

CacheNoHooks:
	if !cached {
		thinkingTreeInsertCacheMut.Lock()
		thinkingTreeInsertCache[key] = cache
		thinkingTreeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ThinkingTree.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ThinkingTree) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	thinkingTreeUpdateCacheMut.RLock()
	cache, cached := thinkingTreeUpdateCache[key]
	thinkingTreeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			thinkingTreeAllColumns,
			thinkingTreePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update thinking_tree, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `thinking_tree` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, thinkingTreePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(thinkingTreeType, thinkingTreeMapping, append(wl, thinkingTreePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update thinking_tree row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for thinking_tree")
	}

	if !cached {
		thinkingTreeUpdateCacheMut.Lock()
		thinkingTreeUpdateCache[key] = cache
		thinkingTreeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q thinkingTreeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for thinking_tree")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for thinking_tree")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ThinkingTreeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), thinkingTreePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `thinking_tree` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, thinkingTreePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in thinkingTree slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all thinkingTree")
	}
	return rowsAff, nil
}

var mySQLThinkingTreeUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ThinkingTree) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no thinking_tree provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(thinkingTreeColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLThinkingTreeUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
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
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	thinkingTreeUpsertCacheMut.RLock()
	cache, cached := thinkingTreeUpsertCache[key]
	thinkingTreeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			thinkingTreeAllColumns,
			thinkingTreeColumnsWithDefault,
			thinkingTreeColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			thinkingTreeAllColumns,
			thinkingTreePrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert thinking_tree, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`thinking_tree`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `thinking_tree` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(thinkingTreeType, thinkingTreeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(thinkingTreeType, thinkingTreeMapping, ret)
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for thinking_tree")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == thinkingTreeMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(thinkingTreeType, thinkingTreeMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for thinking_tree")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for thinking_tree")
	}

CacheNoHooks:
	if !cached {
		thinkingTreeUpsertCacheMut.Lock()
		thinkingTreeUpsertCache[key] = cache
		thinkingTreeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ThinkingTree record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ThinkingTree) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ThinkingTree provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), thinkingTreePrimaryKeyMapping)
	sql := "DELETE FROM `thinking_tree` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from thinking_tree")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for thinking_tree")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q thinkingTreeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no thinkingTreeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from thinking_tree")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for thinking_tree")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ThinkingTreeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(thinkingTreeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), thinkingTreePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `thinking_tree` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, thinkingTreePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from thinkingTree slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for thinking_tree")
	}

	if len(thinkingTreeAfterDeleteHooks) != 0 {
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
func (o *ThinkingTree) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindThinkingTree(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ThinkingTreeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ThinkingTreeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), thinkingTreePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `thinking_tree`.* FROM `thinking_tree` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, thinkingTreePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ThinkingTreeSlice")
	}

	*o = slice

	return nil
}

// ThinkingTreeExists checks if the ThinkingTree row exists.
func ThinkingTreeExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `thinking_tree` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if thinking_tree exists")
	}

	return exists, nil
}

// Exists checks if the ThinkingTree row exists.
func (o *ThinkingTree) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ThinkingTreeExists(ctx, exec, o.ID)
}
