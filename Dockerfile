FROM golang:1.17.2-alpine3.14 as builder
COPY go.mod go.sum /go/src/github.com/patoui/realestate/
WORKDIR /go/src/github.com/patoui/realestate
RUN go mod download
COPY . /go/src/github.com/patoui/realestate
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/realestate github.com/patoui/realestate

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/patoui/realestate/build/realestate /usr/bin/realestate
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/realestate"]