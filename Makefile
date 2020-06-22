test:
	go test ./... -v -coverprofile fmtcoverage.html fmt

test/docker:
	docker-compose run pismo make test

run/api:
	go run api/main.go

run/db:
	docker-compose up pismodb

run/docker:
	docker-compose up --build