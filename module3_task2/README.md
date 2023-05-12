# Post Website

Project for DevOps on Holberton School

## Prerequisites

* [GoHugo](https://gohugo.io/) on extended version -
[v0.84](https://github.com/gohugoio/hugo/releases/tag/v0.84.0)
or [above](https://github.com/gohugoio/hugo/releases)
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
| build | Compile the app to a binary named awesome-api and generate dist/ dir |
| clean   | Clean the dist/ directory, the binary awesome-api and extra files |
| post    | Create a new blog post |
| run | Run the app in bg with awesome-api binary, logging to awesome-api.log |
| stop    | Stop application by process ID |
| test    | Check expected behaviour of the app |
| lint    | Static linting on go files using golangci-lint |
| unit-tests | Execute (successfully) the Golang unit tests |
| integration-tests | Execute the Golang integration tests |
| check   | Markdown linting on syntax and dead links |
| validate | Validate the ./dist/index.html file by W3C Validator |
| help    | Shows this help message |

## Workflow

The [`github-workflow.yml`](./github-workflow.yml) file has a symbolic link to
the corresponding file which is located in the
[`workflows`](../.github/workflows/) folder.
