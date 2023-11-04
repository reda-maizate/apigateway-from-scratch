FROM golang:1.21-alpine

ENV GO111MODULE=on

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/main ./cmd/api

EXPOSE 8000

CMD ["./bin/main"]
