include .env

.EXPORT_ALL_VARIABLES:

dev:
	go run main.go

build:
	CGO_ENABLED=0
	
	go build -ldflags="-s -w" -o ./bin/go-pg ./cmd/pg/pg.go
	go build -ldflags="-s -w" -o ./bin/go-pq ./cmd/pq/pq.go

bench:
	go test -benchmem -run=^$ go-postgres -bench .
	
upx: build
	upx ./bin/go-pg
	upx ./bin/go-pq

run-pg: build
	./bin/go-pg

run-pq: build
	./bin/go-pq

.PHONY: dev build upx run-pg run-pq