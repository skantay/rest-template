include .env
export

compose-up: ### Run docker-compose
	docker-compose up --build -d && docker-compose logs -f
.PHONY: compose-up

compose-down: ### Down docker-compose
	docker-compose down --remove-orphans
.PHONY: compose-down

linter-golangci: ### check by golangci linter
	golangci-lint run
.PHONY: linter-golangci

migrate-create:  ### create new migration
	migrate create -ext sql -dir migrations '$(PROJECT_NAME)'
.PHONY: migrate-create

migrate-up: ### migration up
	migrate -path migrations -database '$(PG_URL_LOCALHOST)?sslmode=disable' up
.PHONY: migrate-up

migrate-down: ### migration down
	echo "y" | migrate -path migrations -database '$(PG_URL_LOCALHOST)?sslmode=disable' down
.PHONY: migrate-down

test: ### run test
	go test -v ./...

coverage-html: ### run test with coverage and open html report
	go test -coverprofile=cvr.out ./...
	go tool cover -html=cvr.out
	rm cvr.out
.PHONY: coverage-html

coverage: ### run test with coverage
	go test -coverprofile=cvr.out ./...
	go tool cover -func=cvr.out
	rm cvr.out
.PHONY: coverage
