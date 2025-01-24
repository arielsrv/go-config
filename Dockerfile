FROM golang:latest AS build

COPY . /go-config
WORKDIR /go-config

RUN go mod tidy
RUN go test ./...
RUN go build cmd/program.go

FROM ubuntu:latest AS runtime
WORKDIR /app
COPY --from=build /go-config /app

# remote environment
ENV ENV=dev

ENTRYPOINT ["./program"]
