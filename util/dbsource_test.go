package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateDbSource(t *testing.T) {
	dbSource := CreateDbSource(
		"user",
		"password",
		"host",
		"port",
		"dbName",
		"dbSSLMode",
	)

	assert.Equal(t, "user='user' password='password' host='host' port='port' dbname='dbName' sslmode='dbSSLMode'", dbSource)
}
