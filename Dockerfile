FROM golang:1.18.1-alpine3.14 as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

FROM golang:1.18.1-alpine3.14 as builder
RUN apk update --no-cache && apk add --no-cache git tzdata make
COPY --from=modules /go/pkg /go/pkg
COPY . /build
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 make build

FROM scratch

ARG TZ=Asia/Jakarta
ENV TZ ${TZ}

COPY --from=builder /build/config.yml.example /config.yml
COPY --from=builder /build/migrations /migrations
COPY --from=builder /build/app /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

CMD ["/app"]
