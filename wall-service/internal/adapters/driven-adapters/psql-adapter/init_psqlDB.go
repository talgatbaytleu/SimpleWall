package psqladapter

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	MainDB *pgxpool.Pool
	TestDB *pgxpool.Pool
	ctx    = context.Background()
)

func InitDB(dbURL string) {
	var err error
	MainDB, err = pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	err = MainDB.Ping(ctx)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	log.SetOutput(os.Stdout)
	log.Println("Connected to database successfully!")
}

func CloseDB() {
	MainDB.Close()
}

func InitTestDB(dbURL string) {
	var err error
	TestDB, err = pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	log.SetOutput(os.Stdout)
	log.Println("Connected to database successfully!")
}

func CloseTestDB() {
	TestDB.Close()
}
