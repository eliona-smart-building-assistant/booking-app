// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package appdb

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

// EventResource is an object representing the database table.
type EventResource struct {
	EventID int32 `boil:"event_id" json:"event_id" toml:"event_id" yaml:"event_id"`
	AssetID int32 `boil:"asset_id" json:"asset_id" toml:"asset_id" yaml:"asset_id"`

	R *eventResourceR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L eventResourceL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var EventResourceColumns = struct {
	EventID string
	AssetID string
}{
	EventID: "event_id",
	AssetID: "asset_id",
}

var EventResourceTableColumns = struct {
	EventID string
	AssetID string
}{
	EventID: "event_resource.event_id",
	AssetID: "event_resource.asset_id",
}

// Generated where

var EventResourceWhere = struct {
	EventID whereHelperint32
	AssetID whereHelperint32
}{
	EventID: whereHelperint32{field: "\"booking\".\"event_resource\".\"event_id\""},
	AssetID: whereHelperint32{field: "\"booking\".\"event_resource\".\"asset_id\""},
}

// EventResourceRels is where relationship names are stored.
var EventResourceRels = struct {
	Event string
}{
	Event: "Event",
}

// eventResourceR is where relationships are stored.
type eventResourceR struct {
	Event *Event `boil:"Event" json:"Event" toml:"Event" yaml:"Event"`
}

// NewStruct creates a new relationship struct
func (*eventResourceR) NewStruct() *eventResourceR {
	return &eventResourceR{}
}

func (r *eventResourceR) GetEvent() *Event {
	if r == nil {
		return nil
	}
	return r.Event
}

// eventResourceL is where Load methods for each relationship are stored.
type eventResourceL struct{}

var (
	eventResourceAllColumns            = []string{"event_id", "asset_id"}
	eventResourceColumnsWithoutDefault = []string{"event_id", "asset_id"}
	eventResourceColumnsWithDefault    = []string{}
	eventResourcePrimaryKeyColumns     = []string{"event_id", "asset_id"}
	eventResourceGeneratedColumns      = []string{}
)

type (
	// EventResourceSlice is an alias for a slice of pointers to EventResource.
	// This should almost always be used instead of []EventResource.
	EventResourceSlice []*EventResource
	// EventResourceHook is the signature for custom EventResource hook methods
	EventResourceHook func(context.Context, boil.ContextExecutor, *EventResource) error

	eventResourceQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	eventResourceType                 = reflect.TypeOf(&EventResource{})
	eventResourceMapping              = queries.MakeStructMapping(eventResourceType)
	eventResourcePrimaryKeyMapping, _ = queries.BindMapping(eventResourceType, eventResourceMapping, eventResourcePrimaryKeyColumns)
	eventResourceInsertCacheMut       sync.RWMutex
	eventResourceInsertCache          = make(map[string]insertCache)
	eventResourceUpdateCacheMut       sync.RWMutex
	eventResourceUpdateCache          = make(map[string]updateCache)
	eventResourceUpsertCacheMut       sync.RWMutex
	eventResourceUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var eventResourceAfterSelectMu sync.Mutex
var eventResourceAfterSelectHooks []EventResourceHook

var eventResourceBeforeInsertMu sync.Mutex
var eventResourceBeforeInsertHooks []EventResourceHook
var eventResourceAfterInsertMu sync.Mutex
var eventResourceAfterInsertHooks []EventResourceHook

var eventResourceBeforeUpdateMu sync.Mutex
var eventResourceBeforeUpdateHooks []EventResourceHook
var eventResourceAfterUpdateMu sync.Mutex
var eventResourceAfterUpdateHooks []EventResourceHook

var eventResourceBeforeDeleteMu sync.Mutex
var eventResourceBeforeDeleteHooks []EventResourceHook
var eventResourceAfterDeleteMu sync.Mutex
var eventResourceAfterDeleteHooks []EventResourceHook

var eventResourceBeforeUpsertMu sync.Mutex
var eventResourceBeforeUpsertHooks []EventResourceHook
var eventResourceAfterUpsertMu sync.Mutex
var eventResourceAfterUpsertHooks []EventResourceHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *EventResource) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventResourceAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *EventResource) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventResourceBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *EventResource) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventResourceAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *EventResource) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventResourceBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *EventResource) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventResourceAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *EventResource) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventResourceBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *EventResource) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventResourceAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *EventResource) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventResourceBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *EventResource) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventResourceAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddEventResourceHook registers your hook function for all future operations.
func AddEventResourceHook(hookPoint boil.HookPoint, eventResourceHook EventResourceHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		eventResourceAfterSelectMu.Lock()
		eventResourceAfterSelectHooks = append(eventResourceAfterSelectHooks, eventResourceHook)
		eventResourceAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		eventResourceBeforeInsertMu.Lock()
		eventResourceBeforeInsertHooks = append(eventResourceBeforeInsertHooks, eventResourceHook)
		eventResourceBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		eventResourceAfterInsertMu.Lock()
		eventResourceAfterInsertHooks = append(eventResourceAfterInsertHooks, eventResourceHook)
		eventResourceAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		eventResourceBeforeUpdateMu.Lock()
		eventResourceBeforeUpdateHooks = append(eventResourceBeforeUpdateHooks, eventResourceHook)
		eventResourceBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		eventResourceAfterUpdateMu.Lock()
		eventResourceAfterUpdateHooks = append(eventResourceAfterUpdateHooks, eventResourceHook)
		eventResourceAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		eventResourceBeforeDeleteMu.Lock()
		eventResourceBeforeDeleteHooks = append(eventResourceBeforeDeleteHooks, eventResourceHook)
		eventResourceBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		eventResourceAfterDeleteMu.Lock()
		eventResourceAfterDeleteHooks = append(eventResourceAfterDeleteHooks, eventResourceHook)
		eventResourceAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		eventResourceBeforeUpsertMu.Lock()
		eventResourceBeforeUpsertHooks = append(eventResourceBeforeUpsertHooks, eventResourceHook)
		eventResourceBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		eventResourceAfterUpsertMu.Lock()
		eventResourceAfterUpsertHooks = append(eventResourceAfterUpsertHooks, eventResourceHook)
		eventResourceAfterUpsertMu.Unlock()
	}
}

// OneG returns a single eventResource record from the query using the global executor.
func (q eventResourceQuery) OneG(ctx context.Context) (*EventResource, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single eventResource record from the query.
func (q eventResourceQuery) One(ctx context.Context, exec boil.ContextExecutor) (*EventResource, error) {
	o := &EventResource{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "appdb: failed to execute a one query for event_resource")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all EventResource records from the query using the global executor.
func (q eventResourceQuery) AllG(ctx context.Context) (EventResourceSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all EventResource records from the query.
func (q eventResourceQuery) All(ctx context.Context, exec boil.ContextExecutor) (EventResourceSlice, error) {
	var o []*EventResource

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "appdb: failed to assign all query results to EventResource slice")
	}

	if len(eventResourceAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all EventResource records in the query using the global executor
func (q eventResourceQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all EventResource records in the query.
func (q eventResourceQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "appdb: failed to count event_resource rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q eventResourceQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q eventResourceQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "appdb: failed to check if event_resource exists")
	}

	return count > 0, nil
}

// Event pointed to by the foreign key.
func (o *EventResource) Event(mods ...qm.QueryMod) eventQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.EventID),
	}

	queryMods = append(queryMods, mods...)

	return Events(queryMods...)
}

// LoadEvent allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (eventResourceL) LoadEvent(ctx context.Context, e boil.ContextExecutor, singular bool, maybeEventResource interface{}, mods queries.Applicator) error {
	var slice []*EventResource
	var object *EventResource

	if singular {
		var ok bool
		object, ok = maybeEventResource.(*EventResource)
		if !ok {
			object = new(EventResource)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeEventResource)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeEventResource))
			}
		}
	} else {
		s, ok := maybeEventResource.(*[]*EventResource)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeEventResource)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeEventResource))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &eventResourceR{}
		}
		args[object.EventID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &eventResourceR{}
			}

			args[obj.EventID] = struct{}{}

		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`booking.event`),
		qm.WhereIn(`booking.event.id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Event")
	}

	var resultSlice []*Event
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Event")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for event")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for event")
	}

	if len(eventAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Event = foreign
		if foreign.R == nil {
			foreign.R = &eventR{}
		}
		foreign.R.EventResources = append(foreign.R.EventResources, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.EventID == foreign.ID {
				local.R.Event = foreign
				if foreign.R == nil {
					foreign.R = &eventR{}
				}
				foreign.R.EventResources = append(foreign.R.EventResources, local)
				break
			}
		}
	}

	return nil
}

// SetEventG of the eventResource to the related item.
// Sets o.R.Event to related.
// Adds o to related.R.EventResources.
// Uses the global database handle.
func (o *EventResource) SetEventG(ctx context.Context, insert bool, related *Event) error {
	return o.SetEvent(ctx, boil.GetContextDB(), insert, related)
}

// SetEvent of the eventResource to the related item.
// Sets o.R.Event to related.
// Adds o to related.R.EventResources.
func (o *EventResource) SetEvent(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Event) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"booking\".\"event_resource\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"event_id"}),
		strmangle.WhereClause("\"", "\"", 2, eventResourcePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.EventID, o.AssetID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.EventID = related.ID
	if o.R == nil {
		o.R = &eventResourceR{
			Event: related,
		}
	} else {
		o.R.Event = related
	}

	if related.R == nil {
		related.R = &eventR{
			EventResources: EventResourceSlice{o},
		}
	} else {
		related.R.EventResources = append(related.R.EventResources, o)
	}

	return nil
}

// EventResources retrieves all the records using an executor.
func EventResources(mods ...qm.QueryMod) eventResourceQuery {
	mods = append(mods, qm.From("\"booking\".\"event_resource\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"booking\".\"event_resource\".*"})
	}

	return eventResourceQuery{q}
}

// FindEventResourceG retrieves a single record by ID.
func FindEventResourceG(ctx context.Context, eventID int32, assetID int32, selectCols ...string) (*EventResource, error) {
	return FindEventResource(ctx, boil.GetContextDB(), eventID, assetID, selectCols...)
}

// FindEventResource retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindEventResource(ctx context.Context, exec boil.ContextExecutor, eventID int32, assetID int32, selectCols ...string) (*EventResource, error) {
	eventResourceObj := &EventResource{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"booking\".\"event_resource\" where \"event_id\"=$1 AND \"asset_id\"=$2", sel,
	)

	q := queries.Raw(query, eventID, assetID)

	err := q.Bind(ctx, exec, eventResourceObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "appdb: unable to select from event_resource")
	}

	if err = eventResourceObj.doAfterSelectHooks(ctx, exec); err != nil {
		return eventResourceObj, err
	}

	return eventResourceObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *EventResource) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *EventResource) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("appdb: no event_resource provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(eventResourceColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	eventResourceInsertCacheMut.RLock()
	cache, cached := eventResourceInsertCache[key]
	eventResourceInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			eventResourceAllColumns,
			eventResourceColumnsWithDefault,
			eventResourceColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(eventResourceType, eventResourceMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(eventResourceType, eventResourceMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"booking\".\"event_resource\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"booking\".\"event_resource\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "appdb: unable to insert into event_resource")
	}

	if !cached {
		eventResourceInsertCacheMut.Lock()
		eventResourceInsertCache[key] = cache
		eventResourceInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single EventResource record using the global executor.
// See Update for more documentation.
func (o *EventResource) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the EventResource.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *EventResource) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	eventResourceUpdateCacheMut.RLock()
	cache, cached := eventResourceUpdateCache[key]
	eventResourceUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			eventResourceAllColumns,
			eventResourcePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("appdb: unable to update event_resource, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"booking\".\"event_resource\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, eventResourcePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(eventResourceType, eventResourceMapping, append(wl, eventResourcePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "appdb: unable to update event_resource row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "appdb: failed to get rows affected by update for event_resource")
	}

	if !cached {
		eventResourceUpdateCacheMut.Lock()
		eventResourceUpdateCache[key] = cache
		eventResourceUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q eventResourceQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q eventResourceQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "appdb: unable to update all for event_resource")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "appdb: unable to retrieve rows affected for event_resource")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o EventResourceSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o EventResourceSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("appdb: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), eventResourcePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"booking\".\"event_resource\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, eventResourcePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "appdb: unable to update all in eventResource slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "appdb: unable to retrieve rows affected all in update all eventResource")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *EventResource) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns, opts...)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *EventResource) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("appdb: no event_resource provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(eventResourceColumnsWithDefault, o)

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

	eventResourceUpsertCacheMut.RLock()
	cache, cached := eventResourceUpsertCache[key]
	eventResourceUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			eventResourceAllColumns,
			eventResourceColumnsWithDefault,
			eventResourceColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			eventResourceAllColumns,
			eventResourcePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("appdb: unable to upsert event_resource, could not build update column list")
		}

		ret := strmangle.SetComplement(eventResourceAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(eventResourcePrimaryKeyColumns) == 0 {
				return errors.New("appdb: unable to upsert event_resource, could not build conflict column list")
			}

			conflict = make([]string, len(eventResourcePrimaryKeyColumns))
			copy(conflict, eventResourcePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"booking\".\"event_resource\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(eventResourceType, eventResourceMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(eventResourceType, eventResourceMapping, ret)
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
		return errors.Wrap(err, "appdb: unable to upsert event_resource")
	}

	if !cached {
		eventResourceUpsertCacheMut.Lock()
		eventResourceUpsertCache[key] = cache
		eventResourceUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single EventResource record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *EventResource) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single EventResource record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *EventResource) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("appdb: no EventResource provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), eventResourcePrimaryKeyMapping)
	sql := "DELETE FROM \"booking\".\"event_resource\" WHERE \"event_id\"=$1 AND \"asset_id\"=$2"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "appdb: unable to delete from event_resource")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "appdb: failed to get rows affected by delete for event_resource")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q eventResourceQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q eventResourceQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("appdb: no eventResourceQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "appdb: unable to delete all from event_resource")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "appdb: failed to get rows affected by deleteall for event_resource")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o EventResourceSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o EventResourceSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(eventResourceBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), eventResourcePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"booking\".\"event_resource\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, eventResourcePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "appdb: unable to delete all from eventResource slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "appdb: failed to get rows affected by deleteall for event_resource")
	}

	if len(eventResourceAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *EventResource) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("appdb: no EventResource provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *EventResource) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindEventResource(ctx, exec, o.EventID, o.AssetID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EventResourceSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("appdb: empty EventResourceSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EventResourceSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := EventResourceSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), eventResourcePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"booking\".\"event_resource\".* FROM \"booking\".\"event_resource\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, eventResourcePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "appdb: unable to reload all in EventResourceSlice")
	}

	*o = slice

	return nil
}

// EventResourceExistsG checks if the EventResource row exists.
func EventResourceExistsG(ctx context.Context, eventID int32, assetID int32) (bool, error) {
	return EventResourceExists(ctx, boil.GetContextDB(), eventID, assetID)
}

// EventResourceExists checks if the EventResource row exists.
func EventResourceExists(ctx context.Context, exec boil.ContextExecutor, eventID int32, assetID int32) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"booking\".\"event_resource\" where \"event_id\"=$1 AND \"asset_id\"=$2 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, eventID, assetID)
	}
	row := exec.QueryRowContext(ctx, sql, eventID, assetID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "appdb: unable to check if event_resource exists")
	}

	return exists, nil
}

// Exists checks if the EventResource row exists.
func (o *EventResource) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return EventResourceExists(ctx, exec, o.EventID, o.AssetID)
}
