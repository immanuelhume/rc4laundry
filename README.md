## Environment Variables

`DATABASE_URL` This connection string is used by PostGraphile. It could be: `postgres://postgres:postgres@db:5432/rc4laundry_test`.

`POSTGRES_URL` This connection string is used by golang-migrate. It's the same as `DATABASE_URL`, just with SSL disabled: `postgres://postgres:postgres@db:5432/rc4laundry_test?sslmode=false`.

On top of these environment variables, we need a `.env.local` file with these variables, used as the `env_file` for the postgres service in docker-compose.

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
