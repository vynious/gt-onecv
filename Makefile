schema-up:
	goose -dir db/migration postgres "user=shawntyw password=shawntyw dbname=gt-onecv sslmode=disable" up

schema-down:
	goose -dir db/migration postgres "user=shawntyw password=shawntyw dbname=gt-onecv sslmode=disable" down

query:
	sqlc generate

run:
	go run main.go

test:
	go test ./tests -v
