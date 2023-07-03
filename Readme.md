# FireBond Assignment API

To run the application you can go the traditional terminal way, use the custom made make commands or docker-compose command

## Usage

### 1. Run Using docker

Install Postgres

```bash
	docker run --name postgres12  --network <network> -p 5432:5432 -e POSTGRES_USER=<username> -e POSTGRES_PASSWORD=<password> -d postgres:12-alpine
```

create db

```bash
docker exec -it postgres12 createdb --username=<username> --owner=<owner> <dbname>
```

Build docker image

```bash
docker build -t firebond_assessment:latest .
```

Run container

```bash
docker run --name firebond_assessment --network firebond-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://<username>:<password>@postgres12:5432/<dbname>?sslmode=disable" firebond_assessment:latest
```

# OR using docker-compose

This command will install the required dependencies, set up a postgres db and dockerise the api

```bash
docker-compose up
```

### 2. Prepare make commands

To install postgres container:

```bash
make postgres
```

To create db for the application

```bash
make createdb
```

To create db for test

```bash
make createtestdb
```

To drop db

```bash
make dropdb
```

To create migration

```bash
make createmigration
```

To apply migration

```bash
make migrateup
```

To apply migration on test db

```bash
make migratetestdbup
```

To run test

```bash
make test
```

To start application

```bash
make server
```
