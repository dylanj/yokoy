package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/dylanj/yokoy/yokoy"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	baseURL := os.Getenv("YOKOY_URL")
	orgID := os.Getenv("YOKOY_ORG_ID")
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	syncStart := os.Getenv("SYNC_START")
	dbURI := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	y := yokoy.Sync{}
	y.SetSyncStart(syncStart)
	y.SetURL(baseURL)
	y.SetDB(db)
	y.SetCredentials(clientID, clientSecret)
	y.SetOrganizationID(orgID)
	y.Go()
}
