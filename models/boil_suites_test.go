// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	t.Run("Categories", testCategories)
	t.Run("CompanyCards", testCompanyCards)
	t.Run("CostCenters", testCostCenters)
	t.Run("ExpenseCostCenters", testExpenseCostCenters)
	t.Run("ExpenseTaxItems", testExpenseTaxItems)
	t.Run("Expenses", testExpenses)
	t.Run("LegalEntities", testLegalEntities)
	t.Run("Trips", testTrips)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("Categories", testCategoriesDelete)
	t.Run("CompanyCards", testCompanyCardsDelete)
	t.Run("CostCenters", testCostCentersDelete)
	t.Run("ExpenseCostCenters", testExpenseCostCentersDelete)
	t.Run("ExpenseTaxItems", testExpenseTaxItemsDelete)
	t.Run("Expenses", testExpensesDelete)
	t.Run("LegalEntities", testLegalEntitiesDelete)
	t.Run("Trips", testTripsDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Categories", testCategoriesQueryDeleteAll)
	t.Run("CompanyCards", testCompanyCardsQueryDeleteAll)
	t.Run("CostCenters", testCostCentersQueryDeleteAll)
	t.Run("ExpenseCostCenters", testExpenseCostCentersQueryDeleteAll)
	t.Run("ExpenseTaxItems", testExpenseTaxItemsQueryDeleteAll)
	t.Run("Expenses", testExpensesQueryDeleteAll)
	t.Run("LegalEntities", testLegalEntitiesQueryDeleteAll)
	t.Run("Trips", testTripsQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Categories", testCategoriesSliceDeleteAll)
	t.Run("CompanyCards", testCompanyCardsSliceDeleteAll)
	t.Run("CostCenters", testCostCentersSliceDeleteAll)
	t.Run("ExpenseCostCenters", testExpenseCostCentersSliceDeleteAll)
	t.Run("ExpenseTaxItems", testExpenseTaxItemsSliceDeleteAll)
	t.Run("Expenses", testExpensesSliceDeleteAll)
	t.Run("LegalEntities", testLegalEntitiesSliceDeleteAll)
	t.Run("Trips", testTripsSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Categories", testCategoriesExists)
	t.Run("CompanyCards", testCompanyCardsExists)
	t.Run("CostCenters", testCostCentersExists)
	t.Run("ExpenseCostCenters", testExpenseCostCentersExists)
	t.Run("ExpenseTaxItems", testExpenseTaxItemsExists)
	t.Run("Expenses", testExpensesExists)
	t.Run("LegalEntities", testLegalEntitiesExists)
	t.Run("Trips", testTripsExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("Categories", testCategoriesFind)
	t.Run("CompanyCards", testCompanyCardsFind)
	t.Run("CostCenters", testCostCentersFind)
	t.Run("ExpenseCostCenters", testExpenseCostCentersFind)
	t.Run("ExpenseTaxItems", testExpenseTaxItemsFind)
	t.Run("Expenses", testExpensesFind)
	t.Run("LegalEntities", testLegalEntitiesFind)
	t.Run("Trips", testTripsFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("Categories", testCategoriesBind)
	t.Run("CompanyCards", testCompanyCardsBind)
	t.Run("CostCenters", testCostCentersBind)
	t.Run("ExpenseCostCenters", testExpenseCostCentersBind)
	t.Run("ExpenseTaxItems", testExpenseTaxItemsBind)
	t.Run("Expenses", testExpensesBind)
	t.Run("LegalEntities", testLegalEntitiesBind)
	t.Run("Trips", testTripsBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("Categories", testCategoriesOne)
	t.Run("CompanyCards", testCompanyCardsOne)
	t.Run("CostCenters", testCostCentersOne)
	t.Run("ExpenseCostCenters", testExpenseCostCentersOne)
	t.Run("ExpenseTaxItems", testExpenseTaxItemsOne)
	t.Run("Expenses", testExpensesOne)
	t.Run("LegalEntities", testLegalEntitiesOne)
	t.Run("Trips", testTripsOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("Categories", testCategoriesAll)
	t.Run("CompanyCards", testCompanyCardsAll)
	t.Run("CostCenters", testCostCentersAll)
	t.Run("ExpenseCostCenters", testExpenseCostCentersAll)
	t.Run("ExpenseTaxItems", testExpenseTaxItemsAll)
	t.Run("Expenses", testExpensesAll)
	t.Run("LegalEntities", testLegalEntitiesAll)
	t.Run("Trips", testTripsAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("Categories", testCategoriesCount)
	t.Run("CompanyCards", testCompanyCardsCount)
	t.Run("CostCenters", testCostCentersCount)
	t.Run("ExpenseCostCenters", testExpenseCostCentersCount)
	t.Run("ExpenseTaxItems", testExpenseTaxItemsCount)
	t.Run("Expenses", testExpensesCount)
	t.Run("LegalEntities", testLegalEntitiesCount)
	t.Run("Trips", testTripsCount)
	t.Run("Users", testUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("Categories", testCategoriesHooks)
	t.Run("CompanyCards", testCompanyCardsHooks)
	t.Run("CostCenters", testCostCentersHooks)
	t.Run("ExpenseCostCenters", testExpenseCostCentersHooks)
	t.Run("ExpenseTaxItems", testExpenseTaxItemsHooks)
	t.Run("Expenses", testExpensesHooks)
	t.Run("LegalEntities", testLegalEntitiesHooks)
	t.Run("Trips", testTripsHooks)
	t.Run("Users", testUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Categories", testCategoriesInsert)
	t.Run("Categories", testCategoriesInsertWhitelist)
	t.Run("CompanyCards", testCompanyCardsInsert)
	t.Run("CompanyCards", testCompanyCardsInsertWhitelist)
	t.Run("CostCenters", testCostCentersInsert)
	t.Run("CostCenters", testCostCentersInsertWhitelist)
	t.Run("ExpenseCostCenters", testExpenseCostCentersInsert)
	t.Run("ExpenseCostCenters", testExpenseCostCentersInsertWhitelist)
	t.Run("ExpenseTaxItems", testExpenseTaxItemsInsert)
	t.Run("ExpenseTaxItems", testExpenseTaxItemsInsertWhitelist)
	t.Run("Expenses", testExpensesInsert)
	t.Run("Expenses", testExpensesInsertWhitelist)
	t.Run("LegalEntities", testLegalEntitiesInsert)
	t.Run("LegalEntities", testLegalEntitiesInsertWhitelist)
	t.Run("Trips", testTripsInsert)
	t.Run("Trips", testTripsInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {}

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
func TestToManyAdd(t *testing.T) {}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Categories", testCategoriesReload)
	t.Run("CompanyCards", testCompanyCardsReload)
	t.Run("CostCenters", testCostCentersReload)
	t.Run("ExpenseCostCenters", testExpenseCostCentersReload)
	t.Run("ExpenseTaxItems", testExpenseTaxItemsReload)
	t.Run("Expenses", testExpensesReload)
	t.Run("LegalEntities", testLegalEntitiesReload)
	t.Run("Trips", testTripsReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Categories", testCategoriesReloadAll)
	t.Run("CompanyCards", testCompanyCardsReloadAll)
	t.Run("CostCenters", testCostCentersReloadAll)
	t.Run("ExpenseCostCenters", testExpenseCostCentersReloadAll)
	t.Run("ExpenseTaxItems", testExpenseTaxItemsReloadAll)
	t.Run("Expenses", testExpensesReloadAll)
	t.Run("LegalEntities", testLegalEntitiesReloadAll)
	t.Run("Trips", testTripsReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Categories", testCategoriesSelect)
	t.Run("CompanyCards", testCompanyCardsSelect)
	t.Run("CostCenters", testCostCentersSelect)
	t.Run("ExpenseCostCenters", testExpenseCostCentersSelect)
	t.Run("ExpenseTaxItems", testExpenseTaxItemsSelect)
	t.Run("Expenses", testExpensesSelect)
	t.Run("LegalEntities", testLegalEntitiesSelect)
	t.Run("Trips", testTripsSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Categories", testCategoriesUpdate)
	t.Run("CompanyCards", testCompanyCardsUpdate)
	t.Run("CostCenters", testCostCentersUpdate)
	t.Run("ExpenseCostCenters", testExpenseCostCentersUpdate)
	t.Run("ExpenseTaxItems", testExpenseTaxItemsUpdate)
	t.Run("Expenses", testExpensesUpdate)
	t.Run("LegalEntities", testLegalEntitiesUpdate)
	t.Run("Trips", testTripsUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Categories", testCategoriesSliceUpdateAll)
	t.Run("CompanyCards", testCompanyCardsSliceUpdateAll)
	t.Run("CostCenters", testCostCentersSliceUpdateAll)
	t.Run("ExpenseCostCenters", testExpenseCostCentersSliceUpdateAll)
	t.Run("ExpenseTaxItems", testExpenseTaxItemsSliceUpdateAll)
	t.Run("Expenses", testExpensesSliceUpdateAll)
	t.Run("LegalEntities", testLegalEntitiesSliceUpdateAll)
	t.Run("Trips", testTripsSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}
