package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
	"github.com/dylanj/yokoy/yokoy"
)

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
