FROM golang:1.24

WORKDIR /app

COPY . .

RUN go build -o notification-service

EXPOSE 8080

CMD ["./notification-service"]
