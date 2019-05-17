// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Pattern is an object representing the database table.
type Pattern struct {
	ID                   string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	SequencePattern      string    `boil:"sequence_pattern" json:"sequence_pattern" toml:"sequence_pattern" yaml:"sequence_pattern"`
	DateCreated          time.Time `boil:"date_created" json:"date_created" toml:"date_created" yaml:"date_created"`
	ServiceID            string    `boil:"service_id" json:"service_id" toml:"service_id" yaml:"service_id"`
	ThresholdReached     bool      `boil:"threshold_reached" json:"threshold_reached" toml:"threshold_reached" yaml:"threshold_reached"`
	DateLastMatched      time.Time `boil:"date_last_matched" json:"date_last_matched" toml:"date_last_matched" yaml:"date_last_matched"`
	OriginalMatchCount   int64     `boil:"original_match_count" json:"original_match_count" toml:"original_match_count" yaml:"original_match_count"`
	CumulativeMatchCount int64     `boil:"cumulative_match_count" json:"cumulative_match_count" toml:"cumulative_match_count" yaml:"cumulative_match_count"`

	R *patternR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L patternL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PatternColumns = struct {
	ID                   string
	SequencePattern      string
	DateCreated          string
	ServiceID            string
	ThresholdReached     string
	DateLastMatched      string
	OriginalMatchCount   string
	CumulativeMatchCount string
}{
	ID:                   "id",
	SequencePattern:      "sequence_pattern",
	DateCreated:          "date_created",
	ServiceID:            "service_id",
	ThresholdReached:     "threshold_reached",
	DateLastMatched:      "date_last_matched",
	OriginalMatchCount:   "original_match_count",
	CumulativeMatchCount: "cumulative_match_count",
}

// Generated where

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

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var PatternWhere = struct {
	ID                   whereHelperstring
	SequencePattern      whereHelperstring
	DateCreated          whereHelpertime_Time
	ServiceID            whereHelperstring
	ThresholdReached     whereHelperbool
	DateLastMatched      whereHelpertime_Time
	OriginalMatchCount   whereHelperint64
	CumulativeMatchCount whereHelperint64
}{
	ID:                   whereHelperstring{field: `id`},
	SequencePattern:      whereHelperstring{field: `sequence_pattern`},
	DateCreated:          whereHelpertime_Time{field: `date_created`},
	ServiceID:            whereHelperstring{field: `service_id`},
	ThresholdReached:     whereHelperbool{field: `threshold_reached`},
	DateLastMatched:      whereHelpertime_Time{field: `date_last_matched`},
	OriginalMatchCount:   whereHelperint64{field: `original_match_count`},
	CumulativeMatchCount: whereHelperint64{field: `cumulative_match_count`},
}

// PatternRels is where relationship names are stored.
var PatternRels = struct {
	Service         string
	PatternExamples string
}{
	Service:         "Service",
	PatternExamples: "PatternExamples",
}

// patternR is where relationships are stored.
type patternR struct {
	Service         *Service
	PatternExamples ExampleSlice
}

// NewStruct creates a new relationship struct
func (*patternR) NewStruct() *patternR {
	return &patternR{}
}

// patternL is where Load methods for each relationship are stored.
type patternL struct{}

var (
	patternColumns               = []string{"id", "sequence_pattern", "date_created", "service_id", "threshold_reached", "date_last_matched", "original_match_count", "cumulative_match_count"}
	patternColumnsWithoutDefault = []string{"id", "sequence_pattern", "date_created", "service_id", "date_last_matched", "original_match_count", "cumulative_match_count"}
	patternColumnsWithDefault    = []string{"threshold_reached"}
	patternPrimaryKeyColumns     = []string{"id"}
)

type (
	// PatternSlice is an alias for a slice of pointers to Pattern.
	// This should generally be used opposed to []Pattern.
	PatternSlice []*Pattern
	// PatternHook is the signature for custom Pattern hook methods
	PatternHook func(context.Context, boil.ContextExecutor, *Pattern) error

	patternQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	patternType                 = reflect.TypeOf(&Pattern{})
	patternMapping              = queries.MakeStructMapping(patternType)
	patternPrimaryKeyMapping, _ = queries.BindMapping(patternType, patternMapping, patternPrimaryKeyColumns)
	patternInsertCacheMut       sync.RWMutex
	patternInsertCache          = make(map[string]insertCache)
	patternUpdateCacheMut       sync.RWMutex
	patternUpdateCache          = make(map[string]updateCache)
	patternUpsertCacheMut       sync.RWMutex
	patternUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var patternBeforeInsertHooks []PatternHook
var patternBeforeUpdateHooks []PatternHook
var patternBeforeDeleteHooks []PatternHook
var patternBeforeUpsertHooks []PatternHook

var patternAfterInsertHooks []PatternHook
var patternAfterSelectHooks []PatternHook
var patternAfterUpdateHooks []PatternHook
var patternAfterDeleteHooks []PatternHook
var patternAfterUpsertHooks []PatternHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Pattern) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range patternBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Pattern) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range patternBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Pattern) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range patternBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Pattern) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range patternBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Pattern) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range patternAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Pattern) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range patternAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Pattern) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range patternAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Pattern) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range patternAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Pattern) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range patternAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPatternHook registers your hook function for all future operations.
func AddPatternHook(hookPoint boil.HookPoint, patternHook PatternHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		patternBeforeInsertHooks = append(patternBeforeInsertHooks, patternHook)
	case boil.BeforeUpdateHook:
		patternBeforeUpdateHooks = append(patternBeforeUpdateHooks, patternHook)
	case boil.BeforeDeleteHook:
		patternBeforeDeleteHooks = append(patternBeforeDeleteHooks, patternHook)
	case boil.BeforeUpsertHook:
		patternBeforeUpsertHooks = append(patternBeforeUpsertHooks, patternHook)
	case boil.AfterInsertHook:
		patternAfterInsertHooks = append(patternAfterInsertHooks, patternHook)
	case boil.AfterSelectHook:
		patternAfterSelectHooks = append(patternAfterSelectHooks, patternHook)
	case boil.AfterUpdateHook:
		patternAfterUpdateHooks = append(patternAfterUpdateHooks, patternHook)
	case boil.AfterDeleteHook:
		patternAfterDeleteHooks = append(patternAfterDeleteHooks, patternHook)
	case boil.AfterUpsertHook:
		patternAfterUpsertHooks = append(patternAfterUpsertHooks, patternHook)
	}
}

// One returns a single pattern record from the query.
func (q patternQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Pattern, error) {
	o := &Pattern{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for Patterns")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Pattern records from the query.
func (q patternQuery) All(ctx context.Context, exec boil.ContextExecutor) (PatternSlice, error) {
	var o []*Pattern

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Pattern slice")
	}

	if len(patternAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Pattern records in the query.
func (q patternQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count Patterns rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q patternQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if Patterns exists")
	}

	return count > 0, nil
}

// Service pointed to by the foreign key.
func (o *Pattern) Service(mods ...qm.QueryMod) serviceQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.ServiceID),
	}

	queryMods = append(queryMods, mods...)

	query := Services(queryMods...)
	queries.SetFrom(query.Query, "\"Services\"")

	return query
}

// PatternExamples retrieves all the Example's Examples with an executor via pattern_id column.
func (o *Pattern) PatternExamples(mods ...qm.QueryMod) exampleQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"Examples\".\"pattern_id\"=?", o.ID),
	)

	query := Examples(queryMods...)
	queries.SetFrom(query.Query, "\"Examples\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"Examples\".*"})
	}

	return query
}

// LoadService allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (patternL) LoadService(ctx context.Context, e boil.ContextExecutor, singular bool, maybePattern interface{}, mods queries.Applicator) error {
	var slice []*Pattern
	var object *Pattern

	if singular {
		object = maybePattern.(*Pattern)
	} else {
		slice = *maybePattern.(*[]*Pattern)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &patternR{}
		}
		args = append(args, object.ServiceID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &patternR{}
			}

			for _, a := range args {
				if a == obj.ServiceID {
					continue Outer
				}
			}

			args = append(args, obj.ServiceID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`Services`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Service")
	}

	var resultSlice []*Service
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Service")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for Services")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for Services")
	}

	if len(patternAfterSelectHooks) != 0 {
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
		object.R.Service = foreign
		if foreign.R == nil {
			foreign.R = &serviceR{}
		}
		foreign.R.ServicePatterns = append(foreign.R.ServicePatterns, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ServiceID == foreign.ID {
				local.R.Service = foreign
				if foreign.R == nil {
					foreign.R = &serviceR{}
				}
				foreign.R.ServicePatterns = append(foreign.R.ServicePatterns, local)
				break
			}
		}
	}

	return nil
}

// LoadPatternExamples allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (patternL) LoadPatternExamples(ctx context.Context, e boil.ContextExecutor, singular bool, maybePattern interface{}, mods queries.Applicator) error {
	var slice []*Pattern
	var object *Pattern

	if singular {
		object = maybePattern.(*Pattern)
	} else {
		slice = *maybePattern.(*[]*Pattern)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &patternR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &patternR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`Examples`), qm.WhereIn(`pattern_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Examples")
	}

	var resultSlice []*Example
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Examples")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on Examples")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for Examples")
	}

	if len(exampleAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.PatternExamples = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &exampleR{}
			}
			foreign.R.Pattern = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.PatternID {
				local.R.PatternExamples = append(local.R.PatternExamples, foreign)
				if foreign.R == nil {
					foreign.R = &exampleR{}
				}
				foreign.R.Pattern = local
				break
			}
		}
	}

	return nil
}

// SetService of the pattern to the related item.
// Sets o.R.Service to related.
// Adds o to related.R.ServicePatterns.
func (o *Pattern) SetService(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Service) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"Patterns\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, []string{"service_id"}),
		strmangle.WhereClause("\"", "\"", 0, patternPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ServiceID = related.ID
	if o.R == nil {
		o.R = &patternR{
			Service: related,
		}
	} else {
		o.R.Service = related
	}

	if related.R == nil {
		related.R = &serviceR{
			ServicePatterns: PatternSlice{o},
		}
	} else {
		related.R.ServicePatterns = append(related.R.ServicePatterns, o)
	}

	return nil
}

// AddPatternExamples adds the given related objects to the existing relationships
// of the Pattern, optionally inserting them as new records.
// Appends related to o.R.PatternExamples.
// Sets related.R.Pattern appropriately.
func (o *Pattern) AddPatternExamples(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Example) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.PatternID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"Examples\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 0, []string{"pattern_id"}),
				strmangle.WhereClause("\"", "\"", 0, examplePrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.PatternID = o.ID
		}
	}

	if o.R == nil {
		o.R = &patternR{
			PatternExamples: related,
		}
	} else {
		o.R.PatternExamples = append(o.R.PatternExamples, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &exampleR{
				Pattern: o,
			}
		} else {
			rel.R.Pattern = o
		}
	}
	return nil
}

// Patterns retrieves all the records using an executor.
func Patterns(mods ...qm.QueryMod) patternQuery {
	mods = append(mods, qm.From("\"Patterns\""))
	return patternQuery{NewQuery(mods...)}
}

// FindPattern retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPattern(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Pattern, error) {
	patternObj := &Pattern{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"Patterns\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, patternObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from Patterns")
	}

	return patternObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Pattern) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no Patterns provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(patternColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	patternInsertCacheMut.RLock()
	cache, cached := patternInsertCache[key]
	patternInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			patternColumns,
			patternColumnsWithDefault,
			patternColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(patternType, patternMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(patternType, patternMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"Patterns\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"Patterns\" () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT \"%s\" FROM \"Patterns\" WHERE %s", strings.Join(returnColumns, "\",\""), strmangle.WhereClause("\"", "\"", 0, patternPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into Patterns")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for Patterns")
	}

CacheNoHooks:
	if !cached {
		patternInsertCacheMut.Lock()
		patternInsertCache[key] = cache
		patternInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Pattern.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Pattern) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	patternUpdateCacheMut.RLock()
	cache, cached := patternUpdateCache[key]
	patternUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			patternColumns,
			patternPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update Patterns, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"Patterns\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, patternPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(patternType, patternMapping, append(wl, patternPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update Patterns row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for Patterns")
	}

	if !cached {
		patternUpdateCacheMut.Lock()
		patternUpdateCache[key] = cache
		patternUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q patternQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for Patterns")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for Patterns")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PatternSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), patternPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"Patterns\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, patternPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in pattern slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all pattern")
	}
	return rowsAff, nil
}

// Delete deletes a single Pattern record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Pattern) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Pattern provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), patternPrimaryKeyMapping)
	sql := "DELETE FROM \"Patterns\" WHERE \"id\"=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from Patterns")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for Patterns")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q patternQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no patternQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from Patterns")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for Patterns")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PatternSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Pattern slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(patternBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), patternPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"Patterns\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, patternPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from pattern slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for Patterns")
	}

	if len(patternAfterDeleteHooks) != 0 {
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
func (o *Pattern) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPattern(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PatternSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PatternSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), patternPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"Patterns\".* FROM \"Patterns\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, patternPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PatternSlice")
	}

	*o = slice

	return nil
}

// PatternExists checks if the Pattern row exists.
func PatternExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"Patterns\" where \"id\"=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if Patterns exists")
	}

	return exists, nil
}
