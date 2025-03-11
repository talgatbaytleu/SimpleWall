FROM golang:1.24

WORKDIR /app

COPY . .

RUN go build -o wall-service

EXPOSE 8080

CMD ["./wall-service"]
