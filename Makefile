schema-up:
	goose -dir db/migration postgres "user=shawntyw password=shawntyw dbname=gt-onecv sslmode=disable" up

schema-down:
	goose -dir db/migration postgres "user=shawntyw password=shawntyw dbname=gt-onecv sslmode=disable" down

seed:
	goose -dir db/seed -table _db_seeds postgres "user=shawntyw password=shawntyw dbname=gt-onecv sslmode=disable" up

query:
	sqlc generate

run:
	go run main.go