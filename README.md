# yokoy postgres sync

This application syncs data from your [yokoy](https://yokoy.io) organization to a postgres databse via the [yokoy api](https://docs.yokoy.ai/).

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
- [-] Invoice - Scoped to LegalEntity
    - [-] InvoiceLineItem
- [x] InvoiceCategory - Scoped to LegalEntity
- [ ] InvoicePaymentTerms - Scoped to LegalEntity
- [x] InvoiceSupplier - Scoped to LegalEntity
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
- `invoices.serverDate` returns string|[]string|null where strings denote dates
- `invoices.status` returns an integer enum, docs say it is a string
- `invoices.paymentInformation` is documented as `invoices.bankAccount`
- `invoice suppliers.country_code` is inconsistent with every other reference to country
- invoice supplier documentation references org level `/suppliers` which is not documented
- `invoice supplier.taxNumber` example is a `number`, documented as `string`
- `invoice supplier.zipCode` example is a `number`, documented as `string`
- yokoy api client generated with swagger code generation - consider a different generator

### Fetch OpenAPI Definition

**note**: there are some manual changes to the `swagger.json` due to issues with
the source swagger.json

```
# download swagger as json
curl https://api.yokoy.ai/v1/swagger.json > swagger.json

# convert to yaml (so we can inspect)
yq eval -P swagger.json > swagger.yaml

# hack: we need to remove the 200response code from export-tasks/{exportTaskId}/artefacts so that we can generate a valid api

# generate code with oapi-codegen
oapi-codegen -package api swagger.yaml &> api/yokoy.gen.go

# hack: we need to comment out the usage of InvoiceStatus as the API returns ints and the api docs say string
```

### Generate models from DB

This project uses [SQLBoiler](https://github.com/volatiletech/sqlboiler) to
generate code for communicating with the database. You can modify the
migrate.sql, recreate the database and generate the models with sqlboiler.

```
# install sqlboiler
go install github.com/volatiletech/sqlboiler/v4@latest
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest

# modify migrate.sql
# ...

# recreate db and run migrations
dropdb yokoy; createdb yokoy; cat migrate.sql | psql yokoy

# generate models with sqlboiler
sqlboiler psql
```

### Thanks

Special thanks to [Yokoy Switzerland](https://yokoy.io) Ltd for having an Open API.
