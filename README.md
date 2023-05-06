# Post Website

Project for DevOps on Holberton School

## Prerequisites

* [GoHugo](https://gohugo.io/) on extended version - [v0.84](https://github.com/gohugoio/hugo/releases/tag/v0.84.0) or [above](https://github.com/gohugoio/hugo/releases)
* [Make](https://www.gnu.org/software/make/) - v3.81 or above
* [Golang](https://go.dev/) - v1.15 or above
* [NodeJS](https://nodejs.org/en) - v14 or above
* [NPM](https://www.npmjs.com/) - v7 or above
* [golangci-lint](https://golangci-lint.run/)
* [markdownlint-cli](https://github.com/igorshubovych/markdownlint-cli)
* [markdown-link-check](https://www.npmjs.com/package/markdown-link-check)
* [HBTN's W3C-Validator](https://github.com/hs-hq/W3C-Validator)

## Lifecycle

| Command | Description |
| ------- | ----------- |
| build   | Generate the website from the markdown and configuration files in the directory dist/ and compile the source code of the application to a binary named awesome-api |
| clean   | Cleanup the content of the directory dist/, the binary awesome-api and the log file awesome-api.log |
| post    | Create a new blog post whose filename and title come from the environment variables POST_TITLE and POST_NAME |
| run     | Run the application in background by executing the binary awesome-api, and write logs into a file named awesome-api.log |
| stop    | Stop the application with the command kill XXXXX where XXXXX is the Process ID of the application |
| test    | Check expected behaviour of the app |
| lint    | Static linting on go files using golangci-lint |
| unit-tests | Execute (successfully) the Golang unit tests |
| integration-tests | Execute the Golang integration tests |
| check   | Markdown linting on syntax and dead links |
| validate | Should validate the file ./dist/index.html by using the command line Holberton’s W3C Validator |
| help    | Shows this help message |
