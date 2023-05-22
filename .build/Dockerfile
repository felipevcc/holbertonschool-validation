FROM golang:latest
RUN apt update && apt upgrade -y
RUN apt install -y wget make dpkg curl sudo
RUN wget -O hugo_binary.deb https://github.com/gohugoio/hugo/releases/download/v0.111.3/hugo_extended_0.111.3_linux-amd64.deb
RUN sudo dpkg -i hugo_binary.deb
RUN rm hugo_binary.deb
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin 
RUN curl -fsSL https://deb.nodesource.com/setup_20.x | sudo -E bash - && sudo apt-get install -y nodejs
RUN sudo npm install -g markdown-link-check
RUN sudo npm install -g markdownlint-cli
