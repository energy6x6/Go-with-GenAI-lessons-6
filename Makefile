.PHONY: test test-verbose cover vet fmt

# Run all tests (same as CI, minus the coverage gate and job summary).
test:
	go test -count=1 ./...

# Same, but verbose: shows every subtest name and PASS/FAIL line.
test-verbose:
	go test -v -count=1 ./...

# Coverage report per package, printed to the terminal.
cover:
	go test -count=1 -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

# Open an HTML coverage report in your browser (highlights untested lines).
cover-html:
	go test -count=1 -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

vet:
	go vet ./...

fmt:
	gofmt -l .
