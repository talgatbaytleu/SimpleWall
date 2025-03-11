FROM golang:1.24

WORKDIR /app

COPY . .

RUN go build -o post-service

EXPOSE 8080

CMD ["./post-service"]
