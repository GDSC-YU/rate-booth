#!/bin/sh

set -e

echo "run db migration"
/app/migrate -path /app/migration -database "postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=$DB_SSL_MODE" -verbose up

echo "start the app"
exec "$@"