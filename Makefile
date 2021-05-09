CONTAINER_IMAGE?=moon-coin-api
RELEASE?=$(shell git tag --points-at HEAD)

create-mssql:
	docker-compose -f ./build/mssql/docker-compose.yaml up -d

delete-mssql:
	docker-compose -f ./build/mssql/docker-compose.yaml down

run:
	go run ./cmd/main.go

clean:
	rm -f goapp

test: clean
	go test -v -cover ./...

build-api: test
	docker build . --no-cache -t $(CONTAINER_IMAGE):$(RELEASE) -f build/Dockerfile

create-api: build-api
	docker-compose -f ./build/backend/docker-compose.yaml up -d