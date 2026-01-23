package yokoy

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/dylanj/yokoy/api"
)

type Sync struct {
	baseURL        string
	clientID       string
	clientSecret   string
	organizationID string
	accessToken    string
	syncStart      time.Time

	legalEntityIds []string

	db        *sql.DB
	apiClient *api.ClientWithResponses
	logger    *slog.Logger
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	Expires     int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func (s *Sync) fetchAccessToken() error {
	client := &http.Client{}
	authUrl := "https://accounts." + s.baseURL + "/oauth2/token"
	v := url.Values{}
	v.Set("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", authUrl, strings.NewReader(v.Encode()))
	if err != nil {
		return err
	}

	req.SetBasicAuth(s.clientID, s.clientSecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var tok AccessTokenResponse
	err = json.Unmarshal(b, &tok)
	if err != nil {
		return err
	}

	s.accessToken = tok.AccessToken
	return nil
}
func (s *Sync) SetSyncStart(date string) error {
	d, err := time.Parse("2006-01-02", date)
	if err != nil {
		return fmt.Errorf("invalid sync start date format: %w", err)
	}
	s.syncStart = d
	return nil
}
func (s *Sync) SetCredentials(clientID, clientSecret string) {
	s.clientID = clientID
	s.clientSecret = clientSecret
}

func (s *Sync) SetToken(tok string) {
	s.accessToken = tok
}

func (s *Sync) SetURL(url string) {
	// todo validate
	s.baseURL = url
}

func (s *Sync) SetDB(db *sql.DB) {
	// todo validate
	s.db = db
}

func (s *Sync) SetOrganizationID(orgID string) {
	// todo validate
	s.organizationID = orgID
}

func (s *Sync) SetLogger(logger *slog.Logger) {
	s.logger = logger
}

func (s *Sync) log() *slog.Logger {
	if s.logger == nil {
		return slog.Default()
	}
	return s.logger
}

func (s *Sync) SyncLegalEntities() error {
	start := time.Now()
	s.log().Info("Syncing legal entities")
	ctx := context.TODO()

	legalEntities, err := fetchLegalEntities(ctx, s.apiClient)
	if err != nil {
		s.log().Error("Failed to fetch legal entities", "error", err)
		return fmt.Errorf("fetch legal entities: %w", err)
	}

	s.legalEntityIds = make([]string, 0)
	for _, le := range *legalEntities {
		s.legalEntityIds = append(s.legalEntityIds, *le.Id)
	}

	s.log().Info("Fetched legal entities", "count", len(*legalEntities))

	s.log().Debug("Truncating legal entities")
	truncateLegalEntities(ctx, s.db)

	s.log().Debug("Inserting legal entities into database")
	if err := insertLegalEntities(ctx, s.db, legalEntities); err != nil {
		s.log().Error("Failed to insert legal entities", "error", err)
		return fmt.Errorf("insert legal entities: %w", err)
	}

	s.log().Info("Completed syncing legal entities", "duration", time.Since(start))
	return nil
}

func (s *Sync) SyncUsers() error {
	start := time.Now()
	s.log().Info("Syncing users")
	ctx := context.TODO()

	users, err := fetchUsers(ctx, s.apiClient)
	if err != nil {
		s.log().Error("Failed to fetch users", "error", err)
		return fmt.Errorf("fetch users: %w", err)
	}

	s.log().Info("Fetched users", "count", len(*users))

	s.log().Debug("Truncating users")
	truncateUsers(ctx, s.db)

	s.log().Debug("Inserting users into database")
	if err := insertUsers(ctx, s.db, users); err != nil {
		s.log().Error("Failed to insert users", "error", err)
		return fmt.Errorf("insert users: %w", err)
	}

	s.log().Info("Completed syncing users", "duration", time.Since(start))
	return nil
}

func (s *Sync) SyncTrips() error {
	start := time.Now()
	s.log().Info("Syncing trips")
	ctx := context.TODO()

	trips, err := fetchTrips(ctx, s.apiClient)
	if err != nil {
		s.log().Error("Failed to fetch trips", "error", err)
		return fmt.Errorf("fetch trips: %w", err)
	}

	s.log().Info("Fetched trips", "count", len(*trips))

	s.log().Debug("Truncating trips")
	truncateTrips(ctx, s.db)

	s.log().Debug("Inserting trips into database")
	if err := insertTrips(ctx, s.db, trips); err != nil {
		s.log().Error("Failed to insert trips", "error", err)
		return fmt.Errorf("insert trips: %w", err)
	}

	s.log().Info("Completed syncing trips", "duration", time.Since(start))
	return nil
}

func (s *Sync) SyncExpenses() error {
	start := time.Now()
	s.log().Info("Syncing expenses")
	ctx := context.TODO()

	s.log().Debug("Truncating expenses")
	truncateExpenses(ctx, s.db)

	filters := forEachMonthSince("created", s.syncStart)
	s.log().Info("Processing expense filters", "count", len(filters))

	for i, filter := range filters {
		s.log().Debug("Fetching expenses", "filter", filter, "progress", fmt.Sprintf("%d/%d", i+1, len(filters)))

		expenses, err := fetchExpenses(ctx, s.apiClient, filter)
		if err != nil {
			s.log().Error("Failed to fetch expenses", "filter", filter, "error", err)
			return fmt.Errorf("fetch expenses for filter %s: %w", filter, err)
		}

		s.log().Info("Fetched expenses", "count", len(*expenses), "filter", filter)

		s.log().Debug("Inserting expenses into database", "filter", filter)
		if err := insertExpenses(ctx, s.db, expenses); err != nil {
			s.log().Error("Failed to insert expenses", "filter", filter, "error", err)
			return fmt.Errorf("insert expenses for filter %s: %w", filter, err)
		}
	}

	s.log().Info("Completed syncing expenses", "duration", time.Since(start))
	return nil
}

func (s *Sync) SyncTransactions() error {
	start := time.Now()
	s.log().Info("Syncing transactions")
	ctx := context.TODO()

	s.log().Debug("Truncating transactions")
	truncateTransactions(ctx, s.db)

	filters := forEachMonthSince("transactionDate", s.syncStart)
	s.log().Info("Processing transaction filters", "count", len(filters))

	for i, filter := range filters {
		s.log().Debug("Fetching transactions", "filter", filter, "progress", fmt.Sprintf("%d/%d", i+1, len(filters)))

		transactions, err := fetchTransactions(ctx, s.apiClient, filter)
		if err != nil {
			s.log().Error("Failed to fetch transactions", "filter", filter, "error", err)
			return fmt.Errorf("fetch transactions for filter %s: %w", filter, err)
		}

		s.log().Info("Fetched transactions", "count", len(*transactions), "filter", filter)

		s.log().Debug("Inserting transactions into database", "filter", filter)
		if err := insertTransactions(ctx, s.db, transactions); err != nil {
			s.log().Error("Failed to insert transactions", "filter", filter, "error", err)
			return fmt.Errorf("insert transactions for filter %s: %w", filter, err)
		}
	}

	s.log().Info("Completed syncing transactions", "duration", time.Since(start))
	return nil
}

func (s *Sync) SyncCompanyCards() error {
	start := time.Now()
	s.log().Info("Syncing company cards")
	ctx := context.TODO()

	s.log().Debug("Truncating company cards")
	truncateCompanyCards(ctx, s.db)
	for _, legalEntityId := range s.legalEntityIds {
		ccs, err := fetchCompanyCards(ctx, legalEntityId, s.apiClient)
		if err != nil {
			s.log().Error("Failed to fetch company cards", "legal_entity_id", legalEntityId, "error", err)
			return fmt.Errorf("fetch company cards for legal entity %s: %w", legalEntityId, err)
		}

		s.log().Info("Fetched company cards", "count", len(*ccs), "legal_entity_id", legalEntityId)

		if err := insertCompanyCards(ctx, s.db, legalEntityId, ccs); err != nil {
			s.log().Error("Failed to insert company cards", "legal_entity_id", legalEntityId, "error", err)
			return fmt.Errorf("insert company cards for legal entity %s: %w", legalEntityId, err)
		}
	}

	s.log().Info("Completed syncing company cards", "duration", time.Since(start))
	return nil
}

func (s *Sync) SyncCategories() error {
	start := time.Now()
	s.log().Info("Syncing categories")
	ctx := context.TODO()

	s.log().Debug("Truncating categories")
	truncateCategories(ctx, s.db)
	for _, legalEntityId := range s.legalEntityIds {
		categories, err := fetchCategories(ctx, legalEntityId, s.apiClient)
		if err != nil {
			s.log().Error("Failed to fetch categories", "legal_entity_id", legalEntityId, "error", err)
			return fmt.Errorf("fetch categories for legal entity %s: %w", legalEntityId, err)
		}

		s.log().Info("Fetched categories", "count", len(*categories), "legal_entity_id", legalEntityId)

		if err := insertCategories(ctx, s.db, legalEntityId, categories); err != nil {
			s.log().Error("Failed to insert categories", "legal_entity_id", legalEntityId, "error", err)
			return fmt.Errorf("insert categories for legal entity %s: %w", legalEntityId, err)
		}
	}

	s.log().Info("Completed syncing categories", "duration", time.Since(start))
	return nil
}

func (s *Sync) SyncPolicies() error {
	start := time.Now()
	s.log().Info("Syncing policies")
	ctx := context.TODO()

	s.log().Debug("Truncating policies")
	truncatePolicies(ctx, s.db)
	for _, legalEntityId := range s.legalEntityIds {
		policies, err := fetchPolicies(ctx, legalEntityId, s.apiClient)
		if err != nil {
			s.log().Error("Failed to fetch policies", "legal_entity_id", legalEntityId, "error", err)
			return fmt.Errorf("fetch policies for legal entity %s: %w", legalEntityId, err)
		}

		s.log().Info("Fetched policies", "count", len(*policies), "legal_entity_id", legalEntityId)

		if err := insertPolicies(ctx, s.db, legalEntityId, policies); err != nil {
			s.log().Error("Failed to insert policies", "legal_entity_id", legalEntityId, "error", err)
			return fmt.Errorf("insert policies for legal entity %s: %w", legalEntityId, err)
		}
	}

	s.log().Info("Completed syncing policies", "duration", time.Since(start))
	return nil
}

func (s *Sync) SyncTaxRates() error {
	start := time.Now()
	s.log().Info("Syncing tax rates")
	ctx := context.TODO()

	s.log().Debug("Truncating tax rates")
	truncateTaxRates(ctx, s.db)
	for _, legalEntityId := range s.legalEntityIds {
		taxRates, err := fetchTaxRates(ctx, legalEntityId, s.apiClient)
		if err != nil {
			s.log().Error("Failed to fetch tax rates", "legal_entity_id", legalEntityId, "error", err)
			return fmt.Errorf("fetch tax rates for legal entity %s: %w", legalEntityId, err)
		}

		s.log().Info("Fetched tax rates", "count", len(*taxRates), "legal_entity_id", legalEntityId)

		if err := insertTaxRates(ctx, s.db, legalEntityId, taxRates); err != nil {
			s.log().Error("Failed to insert tax rates", "legal_entity_id", legalEntityId, "error", err)
			return fmt.Errorf("insert tax rates for legal entity %s: %w", legalEntityId, err)
		}
	}

	s.log().Info("Completed syncing tax rates", "duration", time.Since(start))
	return nil
}

func (s *Sync) SyncCostCenters() error {
	start := time.Now()
	s.log().Info("Syncing cost centers")
	ctx := context.TODO()

	s.log().Debug("Truncating cost centers")
	truncateCostCenters(ctx, s.db)
	for _, legalEntityId := range s.legalEntityIds {
		costCenters, err := fetchCostCenters(ctx, legalEntityId, s.apiClient)
		if err != nil {
			s.log().Error("Failed to fetch cost centers", "legal_entity_id", legalEntityId, "error", err)
			return fmt.Errorf("fetch cost centers for legal entity %s: %w", legalEntityId, err)
		}

		s.log().Info("Fetched cost centers", "count", len(*costCenters), "legal_entity_id", legalEntityId)

		if err := insertCostCenters(ctx, s.db, legalEntityId, costCenters); err != nil {
			s.log().Error("Failed to insert cost centers", "legal_entity_id", legalEntityId, "error", err)
			return fmt.Errorf("insert cost centers for legal entity %s: %w", legalEntityId, err)
		}
	}

	s.log().Info("Completed syncing cost centers", "duration", time.Since(start))
	return nil
}

func (s *Sync) SyncInvoices() error {
	start := time.Now()
	s.log().Info("Syncing invoices")
	ctx := context.TODO()

	s.log().Debug("Truncating invoices")
	truncateInvoices(ctx, s.db)
	for _, legalEntityId := range s.legalEntityIds {
		invoices, err := fetchInvoices(ctx, legalEntityId, s.apiClient)
		if err != nil {
			s.log().Error("Failed to fetch invoices", "legal_entity_id", legalEntityId, "error", err)
			return fmt.Errorf("fetch invoices for legal entity %s: %w", legalEntityId, err)
		}

		s.log().Info("Fetched invoices", "count", len(*invoices), "legal_entity_id", legalEntityId)

		if err := insertInvoices(ctx, s.db, legalEntityId, invoices); err != nil {
			s.log().Error("Failed to insert invoices", "legal_entity_id", legalEntityId, "error", err)
			return fmt.Errorf("insert invoices for legal entity %s: %w", legalEntityId, err)
		}
	}

	s.log().Info("Completed syncing invoices", "duration", time.Since(start))
	return nil
}

func (s *Sync) SyncTags() error {
	start := time.Now()
	s.log().Info("Syncing tags")
	ctx := context.TODO()

	s.log().Debug("Truncating tags")
	truncateTags(ctx, s.db)
	for _, legalEntityId := range s.legalEntityIds {
		tags, err := fetchTags(ctx, legalEntityId, s.apiClient)
		if err != nil {
			s.log().Error("Failed to fetch tags", "legal_entity_id", legalEntityId, "error", err)
			return fmt.Errorf("fetch tags for legal entity %s: %w", legalEntityId, err)
		}

		s.log().Info("Fetched tags", "count", len(*tags), "legal_entity_id", legalEntityId)

		if err := insertTags(ctx, s.db, legalEntityId, tags); err != nil {
			s.log().Error("Failed to insert tags", "legal_entity_id", legalEntityId, "error", err)
			return fmt.Errorf("insert tags for legal entity %s: %w", legalEntityId, err)
		}
	}

	s.log().Info("Completed syncing tags", "duration", time.Since(start))
	return nil
}

func (s *Sync) SyncInvoiceCategories() error {
	start := time.Now()
	s.log().Info("Syncing invoice categories")
	ctx := context.TODO()

	s.log().Debug("Truncating invoice categories")
	truncateInvoiceCategories(ctx, s.db)
	for _, legalEntityId := range s.legalEntityIds {
		invoiceCategories, err := fetchInvoiceCategories(ctx, legalEntityId, s.apiClient)
		if err != nil {
			s.log().Error("Failed to fetch invoice categories", "legal_entity_id", legalEntityId, "error", err)
			return fmt.Errorf("fetch invoice categories for legal entity %s: %w", legalEntityId, err)
		}

		s.log().Info("Fetched invoice categories", "count", len(*invoiceCategories), "legal_entity_id", legalEntityId)

		if err := insertInvoiceCategories(ctx, s.db, legalEntityId, invoiceCategories); err != nil {
			s.log().Error("Failed to insert invoice categories", "legal_entity_id", legalEntityId, "error", err)
			return fmt.Errorf("insert invoice categories for legal entity %s: %w", legalEntityId, err)
		}
	}

	s.log().Info("Completed syncing invoice categories", "duration", time.Since(start))
	return nil
}

func (s *Sync) SyncSuppliers() error {
	start := time.Now()
	s.log().Info("Syncing suppliers")
	ctx := context.TODO()

	s.log().Debug("Truncating suppliers")
	truncateSuppliers(ctx, s.db)
	for _, legalEntityId := range s.legalEntityIds {
		suppliers, err := fetchSuppliers(ctx, legalEntityId, s.apiClient)
		if err != nil {
			s.log().Error("Failed to fetch suppliers", "legal_entity_id", legalEntityId, "error", err)
			return fmt.Errorf("fetch suppliers for legal entity %s: %w", legalEntityId, err)
		}

		s.log().Info("Fetched suppliers", "count", len(*suppliers), "legal_entity_id", legalEntityId)

		if err := insertSuppliers(ctx, s.db, legalEntityId, suppliers); err != nil {
			s.log().Error("Failed to insert suppliers", "legal_entity_id", legalEntityId, "error", err)
			return fmt.Errorf("insert suppliers for legal entity %s: %w", legalEntityId, err)
		}
	}

	s.log().Info("Completed syncing suppliers", "duration", time.Since(start))
	return nil
}

func (s *Sync) configApiClient() error {
	client := http.Client{Transport: NewTransport(s.accessToken)}
	server := "https://api." + s.baseURL + "/v1/organizations/" + s.organizationID + "/"

	c, err := api.NewClientWithResponses(
		server,
		api.WithHTTPClient(&client),
	)
	if err != nil {
		return err
	}

	s.apiClient = c
	return nil
}

func (s *Sync) setup() error {
	if len(s.baseURL) == 0 {
		return fmt.Errorf("no baseurl specified")
	}

	// grab access token
	if len(s.accessToken) == 0 {
		s.log().Info("Fetching access token")
		if err := s.fetchAccessToken(); err != nil {
			s.log().Error("Failed to fetch access token", "error", err)
			return fmt.Errorf("fetch access token: %w", err)
		}
		s.log().Info("Obtained access token")
	}

	if err := s.configApiClient(); err != nil {
		s.log().Error("Failed to configure API client", "error", err)
		return fmt.Errorf("configure API client: %w", err)
	}

	return nil
}

func (s *Sync) Go() error {
	start := time.Now()
	s.log().Info("Starting sync process")

	if err := s.setup(); err != nil {
		return fmt.Errorf("setup failed: %w", err)
	}

	syncSteps := []struct {
		name string
		fn   func() error
	}{
		{"legal entities", s.SyncLegalEntities},
		{"users", s.SyncUsers},
		{"trips", s.SyncTrips},
		{"company cards", s.SyncCompanyCards},
		{"categories", s.SyncCategories},
		{"cost centers", s.SyncCostCenters},
		{"expenses", s.SyncExpenses},
		{"transactions", s.SyncTransactions},
		{"policies", s.SyncPolicies},
		{"tax rates", s.SyncTaxRates},
		{"tags", s.SyncTags},
		{"invoices", s.SyncInvoices},
		{"invoice categories", s.SyncInvoiceCategories},
		{"suppliers", s.SyncSuppliers},
	}

	for _, step := range syncSteps {
		if err := step.fn(); err != nil {
			s.log().Error("Sync step failed", "step", step.name, "error", err)
			return fmt.Errorf("sync %s failed: %w", step.name, err)
		}
	}

	s.log().Info("Sync process completed successfully", "duration", time.Since(start))
	return nil
}
