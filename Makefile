
test: dep
	go test ./... -v

dev-db:
	sqlite3 ./db/dev.db < ./db/schema.sql

dev-data:
	sqlite3 ./db/dev.db < ./db/devdata.sql

dev-api: build-api
	./cmd/server/server -db ./db/dev.db
 
dep:
	GOPRIVATE=github.com/sano-home go mod vendor

build-api: dep
	cd cmd/server && go build && cd -

build-collector:
	cd cmd/collector && go build && cd -


.PHONY: test dev-db dev-data dev-api dep build-api build-collector
