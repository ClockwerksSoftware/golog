TEST_OPTIONS="-v"
TEST_COVER_PROFILE=".coverage"

all: test

test:
	go test -cover -covermode=atomic -coverprofile=$(TEST_COVER_PROFILE) ./... $(TEST_OPTIONS)

coverage: test
	go tool cover -html=$(TEST_COVER_PROFILE)
	go tool cover -func=$(TEST_COVER_PROFILE)

lint:
	golangci-lint run
