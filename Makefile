BINARY_NAME=server
MAIN_SERVER_FILE=cmd/server/main.go

run-server:
	@cd server && go run $(MAIN_SERVER_FILE)

build-server:
	@cd server && go build -o $(BINARY_NAME) $(MAIN_SERVER_FILE)

install-front:
	@cd Front/github_app && npm install

run-front:
	@cd Front/github_app && npm run dev

run-project:
	@cd Front/github_app && npm run dev &
	@cd server && go run $(MAIN_SERVER_FILE)
