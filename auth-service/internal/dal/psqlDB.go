package dal

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	DB  *pgxpool.Pool
	ctx = context.Background()
)

func InitializeDB(dbURL string) {
	var err error
	DB, err = pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	fmt.Println("Connected to database successfully!")
}

func CloseDB() {
	DB.Close()
}
