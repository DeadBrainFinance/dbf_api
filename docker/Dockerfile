FROM golang:1.23

ARG API_PORT

WORKDIR /app

EXPOSE ${API_PORT}

COPY . .

RUN go mod tidy

RUN go install github.com/air-verse/air@latest && go install github.com/pressly/goose/v3/cmd/goose@latest
