package main

import (
	"database/sql"
	"log"
	"log/slog"
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

	// Initialize structured logger
	logLevel := slog.LevelInfo
	if os.Getenv("LOG_LEVEL") == "debug" {
		logLevel = slog.LevelDebug
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
	}))
	slog.SetDefault(logger)

	baseURL := os.Getenv("YOKOY_URL")
	orgID := os.Getenv("YOKOY_ORG_ID")
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	syncStart := os.Getenv("SYNC_START")
	dbURI := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		slog.Error("Failed to open database connection", "error", err)
		os.Exit(1)
	}

	if err = db.Ping(); err != nil {
		slog.Error("Failed to ping database", "error", err)
		os.Exit(1)
	}

	y := yokoy.Sync{}
	y.SetLogger(logger)
	if err := y.SetSyncStart(syncStart); err != nil {
		slog.Error("Invalid sync start date", "error", err, "sync_start", syncStart)
		os.Exit(1)
	}
	y.SetURL(baseURL)
	y.SetDB(db)
	y.SetCredentials(clientID, clientSecret)
	y.SetOrganizationID(orgID)
	
	if err := y.Go(); err != nil {
		slog.Error("Sync failed", "error", err)
		os.Exit(1)
	}
}
