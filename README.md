# go-dummy-service
Basic CRUD server written in go using gorm

## Installation
* Install and setup GO [go 1.18](https://go.dev/dl/)
* Install and setup PostgreSQL
* Install flyway (On Mac: `brew install flyway`)
* Create database `go_crud` with user `maanav` (you can tweak if you want, check `db/db.go`)
* `git clone git@github.com/maanavshah/go-gorm`
* `cd go-gorm`
* `go get -u`
* `flyway -configFiles="./conf/flyway.conf" migrate`
* `go run main.go`

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
