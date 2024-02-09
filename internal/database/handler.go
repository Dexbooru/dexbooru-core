package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DatabaseHandler struct {
	*pgxpool.Pool
}

func (handler *DatabaseHandler) FetchFromQuery(query string) pgx.Row {
	row := handler.QueryRow(context.Background(), query)
	return row
}

func (handler *DatabaseHandler) FetchAllFromQuery(query string) pgx.Rows {
	rows, err := handler.Query(context.Background(), query)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return rows
}

func NewDatabaseHandler() *DatabaseHandler {
	dbConnString := os.Getenv("DATABASE_URL")		
	dbPoolConfig, err := pgxpool.ParseConfig(dbConnString)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	dbPoolConfig.MaxConns = 10
	dbPoolConfig.HealthCheckPeriod = 10 * time.Second
	dbPoolConfig.MaxConnLifetime = 30 * time.Second

	dbpool, err := pgxpool.NewWithConfig(context.Background(), dbPoolConfig)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &DatabaseHandler{
		dbpool,
	}
}
