linter:
	golangci-lint run

test:
	go test ./...

test-coverage: 
	go test -cover ./...

test-cover-mkfile:
	go test -cover -coverprofile="coverage.out" ./...

test-cover-report:
	go tool cover -html="coverage.out"