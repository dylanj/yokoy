package main

import (
	"database/sql"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
	"github.com/salesfive/yokoy/yokoy"
)

func getAccessToken() string {
	baseUrl := os.Getenv("YOKOY_URL")
	clientId := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")

	client := &http.Client{}
	authUrl := "https://accounts." + baseUrl + "/oauth2/token"
	v := url.Values{}
	v.Set("grant_type", "client_credentials")
	req, err := http.NewRequest("POST", authUrl, strings.NewReader(v.Encode()))
	req.SetBasicAuth(clientId, clientSecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	return string(bodyText)
}

type AddHeaderTransport struct {
	T http.RoundTripper
}

func (adt *AddHeaderTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer rxDrcrl8cDLKbXi3OXTbvp5D8Pgy8rZfhvI0oPdCJl")
	req.Header.Set("X-Yk-Auth-Method", "yokoy")
	return adt.T.RoundTrip(req)
}

func NewAddHeaderTransport(T http.RoundTripper) *AddHeaderTransport {
	if T == nil {
		T = http.DefaultTransport
	}
	return &AddHeaderTransport{T}
}

func main() {
	db, err := sql.Open("postgres", "dbname=yokoy user=dylan sslmode=disable")
	if err != nil {
		panic(err)
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	baseURL := os.Getenv("YOKOY_URL")
	orgID := os.Getenv("YOKOY_ORG_ID")
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")

	y := yokoy.Sync{}
	y.SetURL(baseURL)
	y.SetDB(db)
	y.SetCredentials(clientID, clientSecret)
	y.SetOrganizationID(orgID)
	y.Go()
}
