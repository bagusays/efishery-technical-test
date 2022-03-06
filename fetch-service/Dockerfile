FROM golang:alpine3.15 AS builder

RUN apk update && apk upgrade && apk add --no-cache git tzdata ca-certificates && \
cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime && \
echo "Asia/Jakarta" > /etc/timezone

WORKDIR /app
COPY . .

RUN go mod tidy && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -installsuffix cgo -ldflags="-s -w" -o engine

FROM scratch

COPY --from=builder /app/engine /app/engine
COPY --from=builder /app/config /app/config
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /etc/localtime /etc/localtime
COPY --from=builder /etc/timezone /etc/timezone
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

WORKDIR /app

EXPOSE 8081

ENTRYPOINT ["/app/engine"]