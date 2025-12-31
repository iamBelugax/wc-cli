run:
	@go run main.go $(ARGS)

test:
	@go test -v ./...

bench:
	@go test -bench=. -benchmem