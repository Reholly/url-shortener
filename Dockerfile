FROM golang:1.22-alpine AS builder

WORKDIR /app

RUN apk --no-cache add bash git make gettext

#dependencies
COPY go.* ./
RUN go mod download

COPY ./ ./

RUN go build -o ./bin/url-shortener cmd/urlshortener/main.go

FROM alpine AS runner

COPY --from=builder /app/bin/url-shortener /
COPY configs/config.yaml /configs/config.yaml
COPY .env /.env

CMD ["/url-shortener"]
