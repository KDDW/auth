package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

func NewDB() *bun.DB {
	connectionString := os.Getenv("POSTGRES_URL")
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(connectionString)))

	db := bun.NewDB(sqldb, pgdialect.New())

	if os.Getenv("DEBUG") == "true" {
		db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	}

	err := db.Ping()

	if err != nil {
		log.Fatal("cannot ping database", err)
	}

	return db
}
