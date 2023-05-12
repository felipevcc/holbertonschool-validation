#!/bin/bash
apt update && apt upgrade -y
apt install -y wget make dpkg

# Hugo installation
wget -O hugo_binary.deb https://github.com/gohugoio/hugo/releases/download/v0.111.3/hugo_extended_0.111.3_linux-amd64.deb
dpkg -i hugo_binary.deb
rm hugo_binary.deb

# golangci-lint installation
wget -q -O golangci-lint.deb https://github.com/golangci/golangci-lint/releases/download/v1.52.2/golangci-lint-1.52.2-linux-amd64.deb
dpkg -i golangci-lint.deb
rm golangci-lint.deb

# markdown installation
npm install -g markdownlint-cli markdown-link-check
