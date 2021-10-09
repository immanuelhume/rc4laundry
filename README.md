## Requirements

- Docker v20.10+
- Docker Compose v1.29+
- Go v1.17+ (optional, for tests only)

## Environment Variables

`DATABASE_URL` This connection string is used by [PostGraphile](https://github.com/graphile/postgraphile). It could be: `postgres://postgres:postgres@db:5432/rc4laundry_test`.

`MIGRATIONS_URL` This connection string is used by [golang-migrate](https://github.com/golang-migrate/migrate). It's the same as `DATABASE_URL`, just with SSL disabled: `postgres://postgres:postgres@db:5432/rc4laundry_test?sslmode=false` for now.

Export these variables, or place them in a `.env` file. It must be named `.env` for docker-compose to detect it automatically.

We also need a `.env.db.local` file with these three variables, used in the `env_file` field for the postgres service in docker-compose. It looks something like these.

```
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=rc4laundry_test
```

Once these environment variables are in place, run this.

```
docker-compose up
```

If there are no errors, the GraphQL API playground and docs can be accessed at http://localhost:5433/graphiql. The API endpoint is http://localhost:5433/graphql.

## Tests

Tests on the database are written in Go. This will run the tests locally.

```
cd tests
go test
```
