# build stage
FROM golang:1.20.1-alpine3.17 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz

# run stage
FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate
COPY /api/db/migration ./migration
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
RUN chmod +x /app/start.sh
RUN chmod +x /app/wait-for.sh

EXPOSE 8080

CMD [ "/app/main" ]