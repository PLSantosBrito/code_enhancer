BINARY_NAME=server
MAIN_FILE=cmd/server/main.go

run:
	@cd server && go run $(MAIN_FILE)

build:
	@cd server && go build -o $(BINARY_NAME) $(MAIN_FILE)