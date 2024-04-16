package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *Queries
var connPool *pgxpool.Pool

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:Mozart%2330@localhost:5433/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	var err error
	connPool, err = pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(connPool)

	os.Exit(m.Run())
}
