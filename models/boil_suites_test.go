// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Examples", testExamples)
	t.Run("Patterns", testPatterns)
	t.Run("Services", testServices)
}

func TestDelete(t *testing.T) {
	t.Run("Examples", testExamplesDelete)
	t.Run("Patterns", testPatternsDelete)
	t.Run("Services", testServicesDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Examples", testExamplesQueryDeleteAll)
	t.Run("Patterns", testPatternsQueryDeleteAll)
	t.Run("Services", testServicesQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Examples", testExamplesSliceDeleteAll)
	t.Run("Patterns", testPatternsSliceDeleteAll)
	t.Run("Services", testServicesSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Examples", testExamplesExists)
	t.Run("Patterns", testPatternsExists)
	t.Run("Services", testServicesExists)
}

func TestFind(t *testing.T) {
	t.Run("Examples", testExamplesFind)
	t.Run("Patterns", testPatternsFind)
	t.Run("Services", testServicesFind)
}

func TestBind(t *testing.T) {
	t.Run("Examples", testExamplesBind)
	t.Run("Patterns", testPatternsBind)
	t.Run("Services", testServicesBind)
}

func TestOne(t *testing.T) {
	t.Run("Examples", testExamplesOne)
	t.Run("Patterns", testPatternsOne)
	t.Run("Services", testServicesOne)
}

func TestAll(t *testing.T) {
	t.Run("Examples", testExamplesAll)
	t.Run("Patterns", testPatternsAll)
	t.Run("Services", testServicesAll)
}

func TestCount(t *testing.T) {
	t.Run("Examples", testExamplesCount)
	t.Run("Patterns", testPatternsCount)
	t.Run("Services", testServicesCount)
}

func TestHooks(t *testing.T) {
	t.Run("Examples", testExamplesHooks)
	t.Run("Patterns", testPatternsHooks)
	t.Run("Services", testServicesHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Examples", testExamplesInsert)
	t.Run("Examples", testExamplesInsertWhitelist)
	t.Run("Patterns", testPatternsInsert)
	t.Run("Patterns", testPatternsInsertWhitelist)
	t.Run("Services", testServicesInsert)
	t.Run("Services", testServicesInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("ExampleToPatternUsingPattern", testExampleToOnePatternUsingPattern)
	t.Run("ExampleToServiceUsingService", testExampleToOneServiceUsingService)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("PatternToPatternExamples", testPatternToManyPatternExamples)
	t.Run("PatternToServiceIdServices", testPatternToManyServiceIdServices)
	t.Run("ServiceToServiceExamples", testServiceToManyServiceExamples)
	t.Run("ServiceToPatternIdPatterns", testServiceToManyPatternIdPatterns)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("ExampleToPatternUsingPatternExamples", testExampleToOneSetOpPatternUsingPattern)
	t.Run("ExampleToServiceUsingServiceExamples", testExampleToOneSetOpServiceUsingService)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("PatternToPatternExamples", testPatternToManyAddOpPatternExamples)
	t.Run("PatternToServiceIdServices", testPatternToManyAddOpServiceIdServices)
	t.Run("ServiceToServiceExamples", testServiceToManyAddOpServiceExamples)
	t.Run("ServiceToPatternIdPatterns", testServiceToManyAddOpPatternIdPatterns)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("PatternToServiceIdServices", testPatternToManySetOpServiceIdServices)
	t.Run("ServiceToPatternIdPatterns", testServiceToManySetOpPatternIdPatterns)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("PatternToServiceIdServices", testPatternToManyRemoveOpServiceIdServices)
	t.Run("ServiceToPatternIdPatterns", testServiceToManyRemoveOpPatternIdPatterns)
}

func TestReload(t *testing.T) {
	t.Run("Examples", testExamplesReload)
	t.Run("Patterns", testPatternsReload)
	t.Run("Services", testServicesReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Examples", testExamplesReloadAll)
	t.Run("Patterns", testPatternsReloadAll)
	t.Run("Services", testServicesReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Examples", testExamplesSelect)
	t.Run("Patterns", testPatternsSelect)
	t.Run("Services", testServicesSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Examples", testExamplesUpdate)
	t.Run("Patterns", testPatternsUpdate)
	t.Run("Services", testServicesUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Examples", testExamplesSliceUpdateAll)
	t.Run("Patterns", testPatternsSliceUpdateAll)
	t.Run("Services", testServicesSliceUpdateAll)
}
