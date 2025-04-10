# Run main app
run:
	go run ./cmd/main.go

# Run tests
test:
	go test ./... -v

# Run tests with coverage
cover:
	go test ./... -cover -v

# Generate HTML coverage report
cover-html:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

# Clean up generated files
clean:
	rm -f coverage.out
