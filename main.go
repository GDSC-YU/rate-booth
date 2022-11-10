package main

import (
	api "backend/api"
	db "backend/db/sqlc"
	"backend/util"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	dbSource := util.CreateDbSource(
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
		config.DBSSLMode,
	)
	conn, err := sql.Open(config.DBDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	driver, err := postgres.WithInstance(conn, &postgres.Config{})
	if err != nil {
		log.Fatal("cannot create db driver:", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migration",
		config.DBDriver,
		driver,
	)
	if err != nil {
		log.Fatal("cannot create migration instance:", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal("cannot run migration:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.Port)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

	fmt.Println("Running on port:", config.Port)
}
