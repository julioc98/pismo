test:
	go test ./... -v -coverprofile fmtcoverage.html fmt

run/api:
	go run api/main.go
