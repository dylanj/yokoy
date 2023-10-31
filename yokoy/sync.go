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

	"github.com/salesfive/yokoy/api"
)

type Sync struct {
	baseURL        string
	clientID       string
	clientSecret   string
	organizationID string
	accessToken    string

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

}

func (s *Sync) SyncTrips() {
	fmt.Println("syncing trips")

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
	// fetch all legal entites
	// fetch all users
	// fetch all trips
}
