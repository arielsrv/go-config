FROM golang:latest AS build

ADD . /app
WORKDIR /app

RUN go mod tidy
RUN go build cmd/program.go

FROM debian:latest AS runtime
WORKDIR /app
COPY --from=build /app /app

# remote environment
ENV ENV=dev

ENTRYPOINT ["./program"]
