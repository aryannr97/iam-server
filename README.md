# iam-server

Microservice serving REST APIs for user CRUD operations

## Pre-requisites
- [Golang-1.14 & above](https://golang.org/doc/install)
- [Docker](https://www.docker.com/products/docker-desktop)
- [Data-Server](https://github.com/aryannr97/data-server)

## Build & Execution

```
docker-compose up --build
```

## Test coverage

```
go test ./... -cover
```

## API Documentation

Once server is up, swagger docs can be accessed through following endpoint

```
http://<hostname>:<port>/swagger/index.html
```
