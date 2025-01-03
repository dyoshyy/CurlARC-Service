# CurlARC-Service
This is a service that provides a RESTful API for the CurlARC application.

This repository uses the following technologies:
- Go
- Echo (Web Framework)
- GORM (ORM)
- Atlas (Migration Tool)
- PostgreSQL

## Set up
The following command launches api server & db server.
```sh
$ docker compose up
```

## Development Usage
### Check the database
```sh
$ docker exec -it $(container_id) bash
$ psql -U app -d app
$ \dt
$ SELECT * FROM ${table_name};
```
### How to proxy the flyio database
```sh
$ flyctl proxy 5432 -a ${app_name}
```

### Apply migration file
move to the `./migrations` directory and run the following command.
```sh
$ make migrate-up
```

### Generate mocks
Generate repository and usecase mocks.
```sh
$ make mockgen
```

### Run tests
```sh
$ make test
```