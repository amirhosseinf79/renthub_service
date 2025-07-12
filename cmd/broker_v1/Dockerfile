FROM golang:1.24-alpine

RUN apk add --no-cache chromium

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o broker1_server ./cmd/broker_v1/worker.go

ENV PATH="/usr/lib/chromium/:${PATH}"
ENV CHROME_BIN=/usr/bin/chromium-browser
ENV ROD_HEADLESS=true

CMD ["./broker1_server"]
