package main

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	_ "github.com/lib/pq"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	"github.com/salesfive/yokoy/models"
	"github.com/salesfive/yokoy/yokoy"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
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

	// Open handle to database like normal
	db, err := sql.Open("postgres", "dbname=yokoy user=dylan sslmode=disable")
	if err != nil {
		panic(err)
	}

	boil.SetDB(db)

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//token := getAccessToken()
	//token := "YLf8hqYDHmrSAhlP5gKISfxzs0JWmQWUWlpnKVQhJo"
	//fmt.Println("token", token)

	baseUrl := "https://api." + os.Getenv("YOKOY_URL")
	//clientId := os.Getenv("CLIENT_ID")
	//clientSecret := os.Getenv("CLIENT_SECRET")
	//token := "fenvWKNcaAemeeAXPTwLwFcWS8xM6tsGnDfa6AkhGu"

	client := http.Client{Transport: NewAddHeaderTransport(nil)}
	//baseUrl = "https://776a-2003-e4-e74d-f800-6ce5-d473-a754-b527.ngrok-free.app"
	orgId := os.Getenv("YOKOY_ORG_ID")
	server := baseUrl + "/v1/organizations/" + orgId + "/"
	c, err := yokoy.NewClientWithResponses(
		server,
		yokoy.WithHTTPClient(&client),
	)
	if err != nil {
		fmt.Println("erro", err)
		return
	}

	fmt.Println(orgId, baseUrl)

	p := &yokoy.GetLegalEntitiesParams{}
	h, err := c.GetLegalEntitiesWithResponse(context.TODO(), p)
	if err != nil {
		fmt.Println("err")
		spew.Dump(h)
		spew.Dump(err)
		return
	}

	if h.StatusCode() != 200 {
		fmt.Println("got non 200")
		return
	}

	fmt.Println("got ", len(*h.JSON200.LegalEntities), "entities")

	db.Exec("truncate legal_entities;")

	for _, le := range *h.JSON200.LegalEntities {
		//spew.Dump(le)
		dbLE := models.LegalEntity{}
		dbLE.Name = null.StringFrom(*le.Name)
		dbLE.ID = *le.Id
		dbLE.Code = null.StringFrom(le.Code)
		dbLE.Language = null.StringFrom(string(le.Language))

		err := dbLE.Insert(context.TODO(), db, boil.Infer())
		if err != nil {
			panic(err)
		}
	}

	/*
		p := &yokoy.GetUsersParams{}
		h, err := c.GetUsersWithResponse(context.TODO(), p)
		if err != nil {
			fmt.Println("err")
			spew.Dump(h)
			spew.Dump(err)
			return
		}
	*/
	/*
		p := &yokoy.GetExpensesParams{}
		h, err := c.GetExpensesWithResponse(context.TODO(), p)
		if err != nil {
			fmt.Println("err")
			spew.Dump(h)
			spew.Dump(err)
			return
		}
	*/
	//spew.Dump(h.JSON200.Users)
	//fmt.Println("got", len(*h.JSON200.Expenses))
	//spew.Dump(h.JSON200.Expenses)
}
