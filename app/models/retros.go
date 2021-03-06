// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Retro is an object representing the database table.
type Retro struct {
	ID        uint        `boil:"id" json:"id" toml:"id" yaml:"id"`
	Username  string      `boil:"username" json:"username" toml:"username" yaml:"username"`
	Type      string      `boil:"type" json:"type" toml:"type" yaml:"type"`
	Message   null.String `boil:"message" json:"message,omitempty" toml:"message" yaml:"message,omitempty"`
	GroupID   int         `boil:"group_id" json:"group_id" toml:"group_id" yaml:"group_id"`
	CreatedAt time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *retroR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L retroL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RetroColumns = struct {
	ID        string
	Username  string
	Type      string
	Message   string
	GroupID   string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	Username:  "username",
	Type:      "type",
	Message:   "message",
	GroupID:   "group_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// Generated where

var RetroWhere = struct {
	ID        whereHelperuint
	Username  whereHelperstring
	Type      whereHelperstring
	Message   whereHelpernull_String
	GroupID   whereHelperint
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	ID:        whereHelperuint{field: `id`},
	Username:  whereHelperstring{field: `username`},
	Type:      whereHelperstring{field: `type`},
	Message:   whereHelpernull_String{field: `message`},
	GroupID:   whereHelperint{field: `group_id`},
	CreatedAt: whereHelpertime_Time{field: `created_at`},
	UpdatedAt: whereHelpertime_Time{field: `updated_at`},
}

// RetroRels is where relationship names are stored.
var RetroRels = struct {
}{}

// retroR is where relationships are stored.
type retroR struct {
}

// NewStruct creates a new relationship struct
func (*retroR) NewStruct() *retroR {
	return &retroR{}
}

// retroL is where Load methods for each relationship are stored.
type retroL struct{}

var (
	retroColumns               = []string{"id", "username", "type", "message", "group_id", "created_at", "updated_at"}
	retroColumnsWithoutDefault = []string{"username", "type", "message", "group_id", "created_at", "updated_at"}
	retroColumnsWithDefault    = []string{"id"}
	retroPrimaryKeyColumns     = []string{"id"}
)

type (
	// RetroSlice is an alias for a slice of pointers to Retro.
	// This should generally be used opposed to []Retro.
	RetroSlice []*Retro
	// RetroHook is the signature for custom Retro hook methods
	RetroHook func(boil.Executor, *Retro) error

	retroQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	retroType                 = reflect.TypeOf(&Retro{})
	retroMapping              = queries.MakeStructMapping(retroType)
	retroPrimaryKeyMapping, _ = queries.BindMapping(retroType, retroMapping, retroPrimaryKeyColumns)
	retroInsertCacheMut       sync.RWMutex
	retroInsertCache          = make(map[string]insertCache)
	retroUpdateCacheMut       sync.RWMutex
	retroUpdateCache          = make(map[string]updateCache)
	retroUpsertCacheMut       sync.RWMutex
	retroUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var retroBeforeInsertHooks []RetroHook
var retroBeforeUpdateHooks []RetroHook
var retroBeforeDeleteHooks []RetroHook
var retroBeforeUpsertHooks []RetroHook

var retroAfterInsertHooks []RetroHook
var retroAfterSelectHooks []RetroHook
var retroAfterUpdateHooks []RetroHook
var retroAfterDeleteHooks []RetroHook
var retroAfterUpsertHooks []RetroHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Retro) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range retroBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Retro) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range retroBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Retro) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range retroBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Retro) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range retroBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Retro) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range retroAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Retro) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range retroAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Retro) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range retroAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Retro) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range retroAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Retro) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range retroAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddRetroHook registers your hook function for all future operations.
func AddRetroHook(hookPoint boil.HookPoint, retroHook RetroHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		retroBeforeInsertHooks = append(retroBeforeInsertHooks, retroHook)
	case boil.BeforeUpdateHook:
		retroBeforeUpdateHooks = append(retroBeforeUpdateHooks, retroHook)
	case boil.BeforeDeleteHook:
		retroBeforeDeleteHooks = append(retroBeforeDeleteHooks, retroHook)
	case boil.BeforeUpsertHook:
		retroBeforeUpsertHooks = append(retroBeforeUpsertHooks, retroHook)
	case boil.AfterInsertHook:
		retroAfterInsertHooks = append(retroAfterInsertHooks, retroHook)
	case boil.AfterSelectHook:
		retroAfterSelectHooks = append(retroAfterSelectHooks, retroHook)
	case boil.AfterUpdateHook:
		retroAfterUpdateHooks = append(retroAfterUpdateHooks, retroHook)
	case boil.AfterDeleteHook:
		retroAfterDeleteHooks = append(retroAfterDeleteHooks, retroHook)
	case boil.AfterUpsertHook:
		retroAfterUpsertHooks = append(retroAfterUpsertHooks, retroHook)
	}
}

// OneG returns a single retro record from the query using the global executor.
func (q retroQuery) OneG() (*Retro, error) {
	return q.One(boil.GetDB())
}

// One returns a single retro record from the query.
func (q retroQuery) One(exec boil.Executor) (*Retro, error) {
	o := &Retro{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for retros")
	}

	if err := o.doAfterSelectHooks(exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all Retro records from the query using the global executor.
func (q retroQuery) AllG() (RetroSlice, error) {
	return q.All(boil.GetDB())
}

// All returns all Retro records from the query.
func (q retroQuery) All(exec boil.Executor) (RetroSlice, error) {
	var o []*Retro

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Retro slice")
	}

	if len(retroAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all Retro records in the query, and panics on error.
func (q retroQuery) CountG() (int64, error) {
	return q.Count(boil.GetDB())
}

// Count returns the count of all Retro records in the query.
func (q retroQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count retros rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q retroQuery) ExistsG() (bool, error) {
	return q.Exists(boil.GetDB())
}

// Exists checks if the row exists in the table.
func (q retroQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if retros exists")
	}

	return count > 0, nil
}

// Retros retrieves all the records using an executor.
func Retros(mods ...qm.QueryMod) retroQuery {
	mods = append(mods, qm.From("`retros`"))
	return retroQuery{NewQuery(mods...)}
}

// FindRetroG retrieves a single record by ID.
func FindRetroG(iD uint, selectCols ...string) (*Retro, error) {
	return FindRetro(boil.GetDB(), iD, selectCols...)
}

// FindRetro retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRetro(exec boil.Executor, iD uint, selectCols ...string) (*Retro, error) {
	retroObj := &Retro{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `retros` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, retroObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from retros")
	}

	return retroObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Retro) InsertG(columns boil.Columns) error {
	return o.Insert(boil.GetDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Retro) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no retros provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	if o.UpdatedAt.IsZero() {
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(retroColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	retroInsertCacheMut.RLock()
	cache, cached := retroInsertCache[key]
	retroInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			retroColumns,
			retroColumnsWithDefault,
			retroColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(retroType, retroMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(retroType, retroMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `retros` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `retros` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `retros` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, retroPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into retros")
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

	o.ID = uint(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == retroMapping["ID"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for retros")
	}

CacheNoHooks:
	if !cached {
		retroInsertCacheMut.Lock()
		retroInsertCache[key] = cache
		retroInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Retro record using the global executor.
// See Update for more documentation.
func (o *Retro) UpdateG(columns boil.Columns) error {
	return o.Update(boil.GetDB(), columns)
}

// Update uses an executor to update the Retro.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Retro) Update(exec boil.Executor, columns boil.Columns) error {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt = currTime

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(columns, nil)
	retroUpdateCacheMut.RLock()
	cache, cached := retroUpdateCache[key]
	retroUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			retroColumns,
			retroPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update retros, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `retros` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, retroPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(retroType, retroMapping, append(wl, retroPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update retros row")
	}

	if !cached {
		retroUpdateCacheMut.Lock()
		retroUpdateCache[key] = cache
		retroUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q retroQuery) UpdateAllG(cols M) error {
	return q.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q retroQuery) UpdateAll(exec boil.Executor, cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec(exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for retros")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o RetroSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o RetroSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), retroPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `retros` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, retroPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in retro slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Retro) UpsertG(updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(boil.GetDB(), updateColumns, insertColumns)
}

var mySQLRetroUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Retro) Upsert(exec boil.Executor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no retros provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	o.UpdatedAt = currTime

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(retroColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLRetroUniqueColumns, o)

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

	retroUpsertCacheMut.RLock()
	cache, cached := retroUpsertCache[key]
	retroUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			retroColumns,
			retroColumnsWithDefault,
			retroColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			retroColumns,
			retroPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert retros, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "retros", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `retros` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(retroType, retroMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(retroType, retroMapping, ret)
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

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for retros")
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

	o.ID = uint(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == retroMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(retroType, retroMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for retros")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRow(cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for retros")
	}

CacheNoHooks:
	if !cached {
		retroUpsertCacheMut.Lock()
		retroUpsertCache[key] = cache
		retroUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteG deletes a single Retro record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Retro) DeleteG() error {
	return o.Delete(boil.GetDB())
}

// Delete deletes a single Retro record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Retro) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Retro provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), retroPrimaryKeyMapping)
	sql := "DELETE FROM `retros` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from retros")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAll deletes all matching rows.
func (q retroQuery) DeleteAll(exec boil.Executor) error {
	if q.Query == nil {
		return errors.New("models: no retroQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec(exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from retros")
	}

	return nil
}

// DeleteAllG deletes all rows in the slice.
func (o RetroSlice) DeleteAllG() error {
	return o.DeleteAll(boil.GetDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o RetroSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Retro slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(retroBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), retroPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `retros` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, retroPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from retro slice")
	}

	if len(retroAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Retro) ReloadG() error {
	if o == nil {
		return errors.New("models: no Retro provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Retro) Reload(exec boil.Executor) error {
	ret, err := FindRetro(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RetroSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty RetroSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RetroSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RetroSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), retroPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `retros`.* FROM `retros` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, retroPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in RetroSlice")
	}

	*o = slice

	return nil
}

// RetroExistsG checks if the Retro row exists.
func RetroExistsG(iD uint) (bool, error) {
	return RetroExists(boil.GetDB(), iD)
}

// RetroExists checks if the Retro row exists.
func RetroExists(exec boil.Executor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `retros` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if retros exists")
	}

	return exists, nil
}
