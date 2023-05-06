## Golang project

# Prerequisites
* golang latest version
* make latest version

# Lifecycle

| Command | Description |
| ------- | ----------- |
| build   | Compile the source code of the application to a binary named awesome-api |
| run     |Run the application in background by executing the binary awesome-api, and write logs into a file named awesome-api.log |
| stop    | Stop the application with the command kill XXXXX where XXXXX is the Process ID of the application |
| clean   | Delete the binary awesome-api and the log file awesome-api.log |
| test    | Check expected behaviour of the app |
| lint    | Static linting on go files using golangci-lint |
| unit-tests | Execute (successfully) the Golang unit tests |
| integration-tests | Execute the Golang integration tests |
| help    | Show this help message |
