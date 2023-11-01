# yokoy -> db sync

This is a quick and dirty app that sync yokoy data from the yokoy api to a postgres database.

### Objects

- [x] LegalEntity
- [x] User
- [x] Trip
- [x] CostCenter - Scoped to LegalEntity
- [x] Category - Scoped to LegalEntity
- [x] CompanyCard - Scoped to LegalEntity
- [ ] ExpenseCategory - Scoped to LegalEntity
- [ ] Policy - Scoped to LegalEntity
- [ ] TaxRate - Scoped to LegalEntity
- [ ] Tag - Scoped to LegalEntity
- [ ] Expense
- [ ] Invoice - Scoped to LegalEntity
- [ ] InvoiceCategory - Scoped to LegalEntity
- [ ] InvoicePaymentTerms - Scoped to LegalEntity
- [ ] InvoiceSupplier - Scoped to LegalEntity
- [ ] InvoicePurchaseOrder - Scoped to LegalEntity

### Notes
- `approvalLimit` on cost center can be a string `"999999"` or an int `999999`
- there are no company cards in the test enviroment i am using
