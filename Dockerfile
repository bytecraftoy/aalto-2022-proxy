# Based on the official Go project example
# https://docs.docker.com/language/golang/build-images/#multi-stage-builds

## Build
FROM docker.io/golang:1.19-bullseye AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /apikeyproxy

## Deploy
FROM docker.io/debian:bullseye-slim

RUN apt-get update \
    && apt-get -y --no-install-recommends install ca-certificates \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

COPY --from=build /apikeyproxy /apikeyproxy

RUN useradd --user-group --home /app app

USER app:app

WORKDIR /app

EXPOSE 8080

ENTRYPOINT ["/apikeyproxy"]
