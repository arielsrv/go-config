ARG GO_VERSION
FROM golang:${GO_VERSION}-alpine AS build

ADD . /app
WORKDIR /app

RUN go build cmd/program.go

FROM debian:latest AS runtime
WORKDIR /app
COPY --from=build /app /app

# remote environment
ENV ENV=dev

ENTRYPOINT ["./program"]
