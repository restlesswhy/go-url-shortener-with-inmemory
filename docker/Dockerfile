FROM golang:1.17-buster

ENV config=docker
# ENV GOPATH=/
WORKDIR /app

COPY ./ /app
# COPY ./ ./

# RUN apt-get update
# RUN apt-get -y install postgresql-client

# RUN chmod +x wait-for-postgres.sh

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

EXPOSE 5000

ENTRYPOINT CompileDaemon --build="go build cmd/url_shortener/main.go" --command=./main