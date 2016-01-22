test:
	@go test -v .

run:
	@go run main.go

dev:
	@watch -n0.5 make test
