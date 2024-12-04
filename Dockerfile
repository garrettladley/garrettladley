FROM golang:1.23-alpine as builder

WORKDIR /app
RUN apk add --no-cache git ca-certificates

COPY ./ /app/
RUN  go build -tags prod -o bin/garrettladley cmd/server/main.go

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/bin/garrettladley /garrettladley

EXPOSE 8080
ENTRYPOINT [ "./garrettladley" ]
