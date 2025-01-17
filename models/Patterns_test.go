// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testPatterns(t *testing.T) {
	t.Parallel()

	query := Patterns()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testPatternsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Pattern{}
	if err = randomize.Struct(seed, o, patternDBTypes, true, patternColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pattern struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Patterns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPatternsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Pattern{}
	if err = randomize.Struct(seed, o, patternDBTypes, true, patternColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pattern struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Patterns().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Patterns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPatternsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Pattern{}
	if err = randomize.Struct(seed, o, patternDBTypes, true, patternColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pattern struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PatternSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Patterns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPatternsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Pattern{}
	if err = randomize.Struct(seed, o, patternDBTypes, true, patternColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pattern struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := PatternExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Pattern exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PatternExists to return true, but got false.")
	}
}

func testPatternsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Pattern{}
	if err = randomize.Struct(seed, o, patternDBTypes, true, patternColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pattern struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	patternFound, err := FindPattern(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if patternFound == nil {
		t.Error("want a record, got nil")
	}
}

func testPatternsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Pattern{}
	if err = randomize.Struct(seed, o, patternDBTypes, true, patternColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pattern struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Patterns().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testPatternsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Pattern{}
	if err = randomize.Struct(seed, o, patternDBTypes, true, patternColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pattern struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Patterns().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPatternsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	patternOne := &Pattern{}
	patternTwo := &Pattern{}
	if err = randomize.Struct(seed, patternOne, patternDBTypes, false, patternColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pattern struct: %s", err)
	}
	if err = randomize.Struct(seed, patternTwo, patternDBTypes, false, patternColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pattern struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = patternOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = patternTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Patterns().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPatternsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	patternOne := &Pattern{}
	patternTwo := &Pattern{}
	if err = randomize.Struct(seed, patternOne, patternDBTypes, false, patternColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pattern struct: %s", err)
	}
	if err = randomize.Struct(seed, patternTwo, patternDBTypes, false, patternColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pattern struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = patternOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = patternTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Patterns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func patternBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Pattern) error {
	*o = Pattern{}
	return nil
}

func patternAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Pattern) error {
	*o = Pattern{}
	return nil
}

func patternAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Pattern) error {
	*o = Pattern{}
	return nil
}

func patternBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Pattern) error {
	*o = Pattern{}
	return nil
}

func patternAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Pattern) error {
	*o = Pattern{}
	return nil
}

func patternBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Pattern) error {
	*o = Pattern{}
	return nil
}

func patternAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Pattern) error {
	*o = Pattern{}
	return nil
}

func patternBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Pattern) error {
	*o = Pattern{}
	return nil
}

func patternAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Pattern) error {
	*o = Pattern{}
	return nil
}

func testPatternsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Pattern{}
	o := &Pattern{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, patternDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Pattern object: %s", err)
	}

	AddPatternHook(boil.BeforeInsertHook, patternBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	patternBeforeInsertHooks = []PatternHook{}

	AddPatternHook(boil.AfterInsertHook, patternAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	patternAfterInsertHooks = []PatternHook{}

	AddPatternHook(boil.AfterSelectHook, patternAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	patternAfterSelectHooks = []PatternHook{}

	AddPatternHook(boil.BeforeUpdateHook, patternBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	patternBeforeUpdateHooks = []PatternHook{}

	AddPatternHook(boil.AfterUpdateHook, patternAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	patternAfterUpdateHooks = []PatternHook{}

	AddPatternHook(boil.BeforeDeleteHook, patternBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	patternBeforeDeleteHooks = []PatternHook{}

	AddPatternHook(boil.AfterDeleteHook, patternAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	patternAfterDeleteHooks = []PatternHook{}

	AddPatternHook(boil.BeforeUpsertHook, patternBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	patternBeforeUpsertHooks = []PatternHook{}

	AddPatternHook(boil.AfterUpsertHook, patternAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	patternAfterUpsertHooks = []PatternHook{}
}

func testPatternsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Pattern{}
	if err = randomize.Struct(seed, o, patternDBTypes, true, patternColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pattern struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Patterns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPatternsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Pattern{}
	if err = randomize.Struct(seed, o, patternDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Pattern struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(patternColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Patterns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPatternToManyPatternExamples(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Pattern
	var b, c Example

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, patternDBTypes, true, patternColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pattern struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, exampleDBTypes, false, exampleColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, exampleDBTypes, false, exampleColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.PatternID = a.ID
	c.PatternID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.PatternExamples().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.PatternID == b.PatternID {
			bFound = true
		}
		if v.PatternID == c.PatternID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := PatternSlice{&a}
	if err = a.L.LoadPatternExamples(ctx, tx, false, (*[]*Pattern)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.PatternExamples); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.PatternExamples = nil
	if err = a.L.LoadPatternExamples(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.PatternExamples); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testPatternToManyAddOpPatternExamples(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Pattern
	var b, c, d, e Example

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, patternDBTypes, false, strmangle.SetComplement(patternPrimaryKeyColumns, patternColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Example{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, exampleDBTypes, false, strmangle.SetComplement(examplePrimaryKeyColumns, exampleColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Example{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddPatternExamples(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.PatternID {
			t.Error("foreign key was wrong value", a.ID, first.PatternID)
		}
		if a.ID != second.PatternID {
			t.Error("foreign key was wrong value", a.ID, second.PatternID)
		}

		if first.R.Pattern != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Pattern != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.PatternExamples[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.PatternExamples[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.PatternExamples().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testPatternToOneServiceUsingService(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Pattern
	var foreign Service

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, patternDBTypes, false, patternColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pattern struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, serviceDBTypes, false, serviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Service struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.ServiceID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Service().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := PatternSlice{&local}
	if err = local.L.LoadService(ctx, tx, false, (*[]*Pattern)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Service == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Service = nil
	if err = local.L.LoadService(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Service == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPatternToOneSetOpServiceUsingService(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Pattern
	var b, c Service

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, patternDBTypes, false, strmangle.SetComplement(patternPrimaryKeyColumns, patternColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, serviceDBTypes, false, strmangle.SetComplement(servicePrimaryKeyColumns, serviceColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, serviceDBTypes, false, strmangle.SetComplement(servicePrimaryKeyColumns, serviceColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Service{&b, &c} {
		err = a.SetService(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Service != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ServicePatterns[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ServiceID != x.ID {
			t.Error("foreign key was wrong value", a.ServiceID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ServiceID))
		reflect.Indirect(reflect.ValueOf(&a.ServiceID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ServiceID != x.ID {
			t.Error("foreign key was wrong value", a.ServiceID, x.ID)
		}
	}
}

func testPatternsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Pattern{}
	if err = randomize.Struct(seed, o, patternDBTypes, true, patternColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pattern struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPatternsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Pattern{}
	if err = randomize.Struct(seed, o, patternDBTypes, true, patternColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pattern struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PatternSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPatternsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Pattern{}
	if err = randomize.Struct(seed, o, patternDBTypes, true, patternColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pattern struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Patterns().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	patternDBTypes = map[string]string{`ID`: `STRING (20, 50)`, `ServiceID`: `STRING`, `SequencePattern`: `STRING (1000)`, `TagPositions`: `STRING`, `DateCreated`: `DATETIME`, `DateLastMatched`: `DATETIME`, `OriginalMatchCount`: `INTEGER`, `CumulativeMatchCount`: `INTEGER`, `IgnorePattern`: `BOOLEAN`, `ComplexityScore`: `DOUBLE`}
	_              = bytes.MinRead
)

func testPatternsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(patternPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(patternColumns) == len(patternPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Pattern{}
	if err = randomize.Struct(seed, o, patternDBTypes, true, patternColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pattern struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Patterns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, patternDBTypes, true, patternPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Pattern struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testPatternsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(patternColumns) == len(patternPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Pattern{}
	if err = randomize.Struct(seed, o, patternDBTypes, true, patternColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pattern struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Patterns().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, patternDBTypes, true, patternPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Pattern struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(patternColumns, patternPrimaryKeyColumns) {
		fields = patternColumns
	} else {
		fields = strmangle.SetComplement(
			patternColumns,
			patternPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := PatternSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}
