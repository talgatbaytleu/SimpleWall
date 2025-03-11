FROM golang:1.24

WORKDIR /app

COPY . .

RUN go build -o comment-service

EXPOSE 8080

CMD ["./comment-service"]
