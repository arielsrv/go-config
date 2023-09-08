FROM golang:1.21.1-alpine AS build

ADD . /app
WORKDIR /app

RUN go build cmd/program.go

FROM debian:latest AS runtime
WORKDIR /app
COPY --from=build /app /app

# remote environment
ENV ENV=dev

ENTRYPOINT ["./program"]
