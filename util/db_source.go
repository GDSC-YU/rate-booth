package util

import "fmt"

func CreateDbSource(
	dbUser string,
	dbPassword string,
	dbHost string,
	dbPort string,
	dbName string,
	dbSSLMode string,
) string {
	dbSource := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
		dbSSLMode,
	)

	return dbSource
}
