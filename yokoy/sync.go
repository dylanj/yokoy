package yokoy

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
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
func (s *Sync) SetSyncStart(date string) {
	d, err := time.Parse("2006-01-02", date)
	if err != nil {
		panic(err)
	}
	s.syncStart = d
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

func (s *Sync) SyncLegalEntities() {
	fmt.Println("syncing legal entities")
	ctx := context.TODO()

	legalEntities, err := fetchLegalEntities(ctx, s.apiClient)
	if err != nil {
		log.Fatalln(err)
	}

	s.legalEntityIds = make([]string, 0)
	for _, le := range *legalEntities {
		s.legalEntityIds = append(s.legalEntityIds, *le.Id)
	}

	fmt.Println("got", len(*legalEntities), "legal entities")

	fmt.Println("truncating legal entities")
	truncateLegalEntities(ctx, s.db)

	fmt.Println("inserting legal entities into db")
	insertLegalEntities(ctx, s.db, legalEntities)

	// fetch legal expenses
	// truncate db
	// insert legal expenses
}

func (s *Sync) SyncUsers() {
	fmt.Println("syncing users")

	ctx := context.TODO()

	users, err := fetchUsers(ctx, s.apiClient)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("got", len(*users), "users")

	fmt.Println("truncating users")
	truncateUsers(ctx, s.db)

	fmt.Println("inserting users into db")
	insertUsers(ctx, s.db, users)
}

func (s *Sync) SyncTrips() {
	fmt.Println("syncing trips")

	ctx := context.TODO()

	Trips, err := fetchTrips(ctx, s.apiClient)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("got", len(*Trips), "trips")

	fmt.Println("truncating trips")
	truncateTrips(ctx, s.db)

	fmt.Println("inserting trips into db")
	insertTrips(ctx, s.db, Trips)
}

func (s *Sync) SyncExpenses() {
	fmt.Println("syncing expenses")

	ctx := context.TODO()

	fmt.Println("truncating expenses")
	truncateExpenses(ctx, s.db)
	for _, filter := range forEachMonthSince("created", s.syncStart) {
		fmt.Println("fetching", filter)

		Expenses, err := fetchExpenses(ctx, s.apiClient, filter)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("got", len(*Expenses), "expenses")

		fmt.Println("inserting expenses into db")
		err = insertExpenses(ctx, s.db, Expenses)
		if err != nil {
			panic(err)
		}
	}
}

func (s *Sync) SyncCompanyCards() {
	fmt.Println("syncing company cards")

	ctx := context.TODO()

	truncateCompanyCards(ctx, s.db)
	for _, legalEntityId := range s.legalEntityIds {
		ccs, err := fetchCompanyCards(ctx, legalEntityId, s.apiClient)

		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("got", len(*ccs), "cards in le", legalEntityId)

		insertCompanyCards(ctx, s.db, legalEntityId, ccs)
	}
}

func (s *Sync) SyncCategories() {
	fmt.Println("syncing categories")

	ctx := context.TODO()

	truncateCategories(ctx, s.db)
	for _, legalEntityId := range s.legalEntityIds {
		ccs, err := fetchCategories(ctx, legalEntityId, s.apiClient)

		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("got", len(*ccs), "categories in le", legalEntityId)

		insertCategories(ctx, s.db, legalEntityId, ccs)
	}
}

func (s *Sync) SyncPolicies() {
	fmt.Println("syncing policies")

	ctx := context.TODO()

	truncatePolicies(ctx, s.db)
	for _, legalEntityId := range s.legalEntityIds {
		ccs, err := fetchPolicies(ctx, legalEntityId, s.apiClient)

		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("got", len(*ccs), "policies in le", legalEntityId)

		insertPolicies(ctx, s.db, legalEntityId, ccs)
	}
}

func (s *Sync) SyncTaxRates() {
	fmt.Println("syncing tax rates")

	ctx := context.TODO()

	truncateTaxRates(ctx, s.db)
	for _, legalEntityId := range s.legalEntityIds {
		ccs, err := fetchTaxRates(ctx, legalEntityId, s.apiClient)

		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("got", len(*ccs), "tax rates in le", legalEntityId)

		insertTaxRates(ctx, s.db, legalEntityId, ccs)
	}
}

func (s *Sync) SyncCostCenters() {
	fmt.Println("syncing cost centers")

	ctx := context.TODO()

	truncateCostCenters(ctx, s.db)
	for _, legalEntityId := range s.legalEntityIds {
		ccs, err := fetchCostCenters(ctx, legalEntityId, s.apiClient)

		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("got", len(*ccs), "cost centers in le", legalEntityId)

		insertCostCenters(ctx, s.db, legalEntityId, ccs)
	}
}

func (s *Sync) SyncInvoices() {
	fmt.Println("syncing invoices")

	ctx := context.TODO()

	truncateInvoices(ctx, s.db)
	for _, legalEntityId := range s.legalEntityIds {
		ccs, err := fetchInvoices(ctx, legalEntityId, s.apiClient)

		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("got", len(*ccs), "invoices in le", legalEntityId)

		insertInvoices(ctx, s.db, legalEntityId, ccs)
	}
}

func (s *Sync) SyncTags() {
	fmt.Println("syncing tags")

	ctx := context.TODO()

	truncateTags(ctx, s.db)
	for _, legalEntityId := range s.legalEntityIds {
		ccs, err := fetchTags(ctx, legalEntityId, s.apiClient)

		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("got", len(*ccs), "tags in le", legalEntityId)

		insertTags(ctx, s.db, legalEntityId, ccs)
	}
}

func (s *Sync) SyncInvoiceCategories() {
	fmt.Println("syncing invoice categories")

	ctx := context.TODO()

	truncateInvoiceCategories(ctx, s.db)
	for _, legalEntityId := range s.legalEntityIds {
		ccs, err := fetchInvoiceCategories(ctx, legalEntityId, s.apiClient)

		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("got", len(*ccs), "invoice categories in le", legalEntityId)

		insertInvoiceCategories(ctx, s.db, legalEntityId, ccs)
	}
}

func (s *Sync) SyncSuppliers() {
	fmt.Println("syncing suppliers")

	ctx := context.TODO()

	truncateSuppliers(ctx, s.db)
	for _, legalEntityId := range s.legalEntityIds {
		ccs, err := fetchSuppliers(ctx, legalEntityId, s.apiClient)

		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("got", len(*ccs), "suppliers in le", legalEntityId)

		insertSuppliers(ctx, s.db, legalEntityId, ccs)
	}
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

func (s *Sync) setup() {
	var err error

	if len(s.baseURL) == 0 {
		log.Fatalln("no baseurl specified")
	}

	// grab access token
	if len(s.accessToken) == 0 {
		log.Println("fetching access token")
		err = s.fetchAccessToken()
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("obtained access token")
	}

	err = s.configApiClient()
	if err != nil {
		log.Fatalln(err)
	}
}

func (s *Sync) Go() {
	s.setup()
	fmt.Println("syncing")
	s.SyncLegalEntities()
	s.SyncUsers()
	s.SyncTrips()
	s.SyncCompanyCards()
	s.SyncCategories()
	s.SyncCostCenters()
	s.SyncExpenses()
	s.SyncPolicies()
	s.SyncTaxRates()
	s.SyncTags()
	s.SyncInvoices()
	s.SyncInvoiceCategories()
	s.SyncSuppliers()
	fmt.Println("done")
}
