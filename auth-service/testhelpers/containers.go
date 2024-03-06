package testhelpers

import (
	"context"
	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
	"time"
)

type PostgresContainer struct {
	*postgres.PostgresContainer
	ConnectionString string
}

func SetupTestDatabase() *PostgresContainer {
	ctx := context.Background()
	dbName := "users"
	dbUser := "user"
	dbPassword := "password"

	pgContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("docker.io/postgres:15.3-alpine"),
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		postgres.WithInitScripts("../../db.sql"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		log.Fatalf("failed to start container: %s", err)
	}
	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		log.Fatalf("could not get connection string: %s", err)
	}
	return &PostgresContainer{
		PostgresContainer: pgContainer,
		ConnectionString:  connStr,
	}
}

func (tdb *PostgresContainer) TearDown() {
	err := tdb.Terminate(context.Background())
	if err != nil {
		log.Fatalf("failed to terminate container: %s", err)
	}
}
