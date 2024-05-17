.PHONY=run
run:
	@go run ./cmd/web


.PHONY=test
test: 
	@go test -v ./...


.PHONY=run_migrate
run_migrate:
	@migrate -path=./migrations -database="postgres://green_toys_user:password@localhost/green_toys_db?sslmode=disable" up
