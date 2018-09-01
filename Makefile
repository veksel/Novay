BRANCH="master"
BUILD_TIME = "$(shell date -u +\"%Y-%m-%dT%H:%M:%SZ\")"
REPONAME = "Novay"

build:
	@echo "Building binary..."
	@go build -i -o ./bin/Novay ./src/main.go
deps:
	@dep ensure
test:
	@go test ./... -cover
