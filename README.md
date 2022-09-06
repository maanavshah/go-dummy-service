# go-dummy-service
Basic CRUD server written in go using gorm

## Installation
* Install and setup GO [go 1.18](https://go.dev/dl/)
* Install and setup PostgreSQL
* Install flyway (On Mac: `brew install flyway`)
* Install `gofmt` (linter), `gopls` (language server by VSCode), `golang/tools` (Static analysis of `.go` files) and other Go packages suggested by VSCode. Use `Opt + Cmd + L` to autoformat `.go` files using `gofmt` once installed
* Install air for live reloading go apps: `go install github.com/cosmtrek/air@latest`
* Create database `go_crud` with user `root` (you can tweak if you want, check `db/db.go`)
* `git clone git@github.com/maanavshah/go-dummy-service`
* `cd go-gorm`
* `go get -u`
* Run `cp conf/conf.yaml.back conf/conf.yaml` to copy application config
* Run `cp conf/flyway.conf.back conf/flyway.conf` to copy flyway config
* Verify database creds and path in `conf/conf.yaml`, run migrations using: `flyway -configFiles="./conf/flyway.conf" migrate`
* Start server `go run main.go`
* (or) start server with hot reload enabled, run `air` in project root

## CURLs

* Health APIs

```
curl --location --request GET 'http://localhost:8000/ping'

curl --location --request GET 'http://localhost:8000/health'
```

* Book APIs 

```
curl --location --request GET 'http://localhost:8000/books'

curl --location --request POST 'http://localhost:8000/books' \
--header 'Content-Type: application/json' \
--data-raw '{
"title": "Harry Potter",
"author": "J K Rowling",
"desc": "HP Books"
}'
```
