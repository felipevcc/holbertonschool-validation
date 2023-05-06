build: ## Compile the source code of the application to a binary named awesome-api
	make lint
	go get github.com/gorilla/mux
	go build

run: ## Run the application in background by executing the binary awesome-api, and write logs into a file named awesome-api.log
	./awesome-api >./awesome-api.log 2>&1 &

stop: ## Stop the application with the command kill XXXXX where XXXXX is the Process ID of the application. 
	kill "$(shell pgrep awesome-api)"

clean: ## Delete the binary awesome-api and the log file awesome-api.log
	-make stop
	-rm ./awesome-api ./awesome-api.log ./coverage-units.out

test: ## Check expected behaviour of the app
	make unit-tests
	make integration-tests

lint: ## Static linting on go files using golangci-lint
	golangci-lint run

unit-tests: ## Execute (successfully) the Golang unit tests:
	@go test -v -short -coverprofile=coverage-units.out

integration-tests: ## Execute the Golang integration tests
	@go test -v -coverprofile=coverage-integrations.out

help: ## Show this help message
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'
