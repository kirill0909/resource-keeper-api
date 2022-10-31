FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

# install psql
RUN apt update
RUN apt install postgresql-client -y

# make script executable
RUN chmod +x wait-for-postgres.sh

# download requriments and build app
RUN go mod download
RUN go build -o .bin/main ./cmd/main.go 

CMD ["./.bin/main"]


