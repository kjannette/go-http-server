# syntax=docker/dockerfile:1

FROM golang:1.25-alpine AS builder

WORKDIR /src

RUN apk add --no-cache ca-certificates git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG APP_VERSION=dev
RUN CGO_ENABLED=0 GOOS=linux go build \
	-ldflags="-s -w -X main.appName=go-http-server -X main.appVersion=${APP_VERSION}" \
	-o /out/server \
	./cmd/app

FROM alpine:3.20

RUN apk add --no-cache ca-certificates wget

WORKDIR /app

RUN addgroup -S app && adduser -S app -G app

COPY --from=builder /out/server /app/server

USER app

EXPOSE 3000

ENTRYPOINT ["/app/server"]
