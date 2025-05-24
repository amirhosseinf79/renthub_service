FROM golang:1.24-alpine

RUN apk add --no-cache chromium

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o renthub_server ./cmd/app/main.go

ENV PATH="/usr/lib/chromium/:${PATH}"
ENV CHROME_BIN=/usr/bin/chromium-browser
ENV ROD_HEADLESS=true

EXPOSE 8080

CMD ["./renthub_server"]
