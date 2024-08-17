# 開発環境
FROM golang:1.22.3-alpine3.19 as dev-app

WORKDIR /go/src/app

RUN apk update && \
    apk add --no-cache git ca-certificates build-base && \
    update-ca-certificates

COPY go.mod go.sum ./
RUN go mod download && go mod verify

ARG CGO_ENABLED=1

RUN go install github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql@latest
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.16.2
RUN go install github.com/air-verse/air@latest
CMD ["air",  "-c", ".air.toml"]
