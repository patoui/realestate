FROM golang:1.17.2-alpine3.14

COPY go.mod go.sum /go/src/github.com/patoui/realestate/
WORKDIR /go/src/github.com/patoui/realestate

RUN go mod download

COPY . .

RUN apk add --no-cache ca-certificates && update-ca-certificates
RUN apk add --update npm

ENV GO111MODULE=on
RUN go get github.com/githubnemo/CompileDaemon

EXPOSE 8080 8080

ENTRYPOINT CompileDaemon -build='go build -o build/realestate github.com/patoui/realestate' -command='./build/realestate' -color=true 2>&1
