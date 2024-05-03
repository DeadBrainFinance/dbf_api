FROM golang:1.22.2

ARG API_PORT

WORKDIR /app

EXPOSE ${API_PORT}

COPY . .

RUN go install github.com/cosmtrek/air@latest && go install github.com/pressly/goose/v3/cmd/goose@latest
