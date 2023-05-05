# Post Website

Project for DevOps on Holberton School

## Prerequisites

* [GoHugo](https://gohugo.io/) on extended version - [v0.84](https://github.com/gohugoio/hugo/releases/tag/v0.84.0) or [above](https://github.com/gohugoio/hugo/releases)
* [make](https://www.gnu.org/software/make/) - v3.81 or above

## Lifecycle

| Command | Description |
| ------- | ----------- |
| build   | Generate the website from the markdown and configuration files in the directory dist/ |
| clean   | Cleanup the content of the directory dist/ |
| post    | Create a new blog post whose filename and title come from the environment variables POST_TITLE and POST_NAME |
| help    | Shows this help message |
