FROM golang:1.22-alpine as builder

WORKDIR /app
RUN apk add --no-cache make nodejs npm git

COPY . ./
RUN make install
RUN make build-prod

FROM scratch
COPY --from=builder /app/bin/garrettladley /garrettladley
COPY --from=builder /app/config/ /config/
ENV APP_ENVIRONMENT production

ENTRYPOINT [ "./garrettladley" ]
