APPMAIN=./cmd/app/main.go
BINARY=./app
TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)

export TAG

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

test: ### Run integration test
	go test -v -cover -covermode=atomic ./...

build: ### Build the app to binary
	go build -tags migrate -ldflags "-X main.version=$(TAG)" -o ${BINARY} ${APPMAIN}

unittest: ### Run unit tests
	go test -v -short ./...

genmock: ### Generate mocks
	mockery --all --keeptree

clean: ### Clear build binaru
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

run: ### Run app in docker
	docker compose -p "cake-store-app" up --build -d

stop: ### Stop app from docker
	docker compose down

run-dev: ### Run app in dev mode (host)
	nodemon --exec go run -tags migrate ${APPMAIN}

migrate-create:  ### Create new migration
	migrate create -ext sql -dir migrations $(name)

migrate-up: ### Apply the migrations
	migrate -path migrations -database '$(PG_URL)?sslmode=disable' up

.PHONY: help clean install unittest build docker run stop run-dev migrate-create migrate-up
