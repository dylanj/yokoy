# yokoy -> db sync

This is a quick and dirty app that sync yokoy data from the yokoy api to a postgres database.

### Objects

- [x] LegalEntity
- [x] User
- [x] Trip
- [x] CostCenter - Scoped to LegalEntity
- [x] Category - Scoped to LegalEntity
- [x] CompanyCard - Scoped to LegalEntity
- [x] Policy - Scoped to LegalEntity
- [x] TaxRate - Scoped to LegalEntity
- [x] Tag - Scoped to LegalEntity
- [x] Expense
    - [x] CostCenter
    - [x] TaxItem
    - [x] ApproverId
    - [x] EventLog
- [ ] Invoice - Scoped to LegalEntity
- [ ] InvoiceCategory - Scoped to LegalEntity
- [ ] InvoicePaymentTerms - Scoped to LegalEntity
- [ ] InvoiceSupplier - Scoped to LegalEntity
- [ ] InvoicePurchaseOrder - Scoped to LegalEntity


### Development Setup

Create a file `.env` with the following content and run the program with `go run main.go`. The environment variables will be loaded from .env and the program will start syncing data to your database.

```
CLIENT_ID="abc"
CLIENT_SECRET="abc"
YOKOY_URL="test.yokoy.ai"
YOKOY_ORG_ID="abc"
```

### Notes
- `approvalLimit` on cost center can be a string `"999999"` or an int `999999`
- there are no company cards in the test enviroment i am using
- there are no policies in the test environment i am using
- `expenses.eventLog.timestamp` - use format `Wed, 24 Aug 2022 08:58:34 GMT`
- `expenses.created`/`lastmodified` use format `1970-01-01T00:00:00.000Z`
