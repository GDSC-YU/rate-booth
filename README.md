# ratebooth

[![codecov](https://codecov.io/gh/GDSC-YU/rate-booth/branch/main/graph/badge.svg?token=2s7qJUPgaa)](https://codecov.io/gh/GDSC-YU/rate-booth)

## Getting Stared

Install the following dependencies:

1. [go v1.19](https://go.dev)
2. [go-task v3.14.1](https://taskfile.dev)
3. [migrate v4.15.2](https://github.com/golang-migrate/migrate)
4. [docker](https://www.docker.com)
5. [docker compose plugin](https://docs.docker.com/compose/install)
6. [sqlc v1.16.0](https://sqlc.dev)
7. [mockgen v1.6.0](https://github.com/golang/mock#installation)

Run the following tasks in order:

1. `docker compose up`
2. `task migrate-up`

Then run `go run main.go` to the run the API

## Create a migration

We use migrations to introduce any change into the database, from new tables to altering existing ones.

Use the following command to create a migration:

```sh
migrate create -seq -ext sql -dir db/migration migration_name
```

> NOTE: replace `migration_name` with the preferred migration name

## Build the api docker image

Run: `docker build -t rate-booth-api .`
