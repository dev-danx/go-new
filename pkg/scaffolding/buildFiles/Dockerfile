FROM golang:1.21.4-bullseye AS builder

COPY . /src/

WORKDIR /src/

RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/app  /src/cmd/main.go
FROM gcr.io/distroless/base-debian11

COPY --from=builder /go/bin/* /go/bin/

ENTRYPOINT ["/go/bin/app"]
EXPOSE 80
EXPOSE 443