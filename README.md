# ratebooth

## Getting Stared

Install the following dependencies:

1. [go v1.19](https://go.dev)
2. [go-task v3.14.1](https://taskfile.dev)
3. [migrate v4.15.2](https://github.com/golang-migrate/migrate)
4. [docker](https://www.docker.com)
5. [docker compose plugin](https://docs.docker.com/compose/install)
6. [sqlc v1.15.0](https://sqlc.dev)
7. [mockgen v1.6.0](https://github.com/golang/mock#installation)

Run the following tasks in order:

1. `docker compose up`
2. `task migrate-up`

Then run `go run main.go` to the run the API
