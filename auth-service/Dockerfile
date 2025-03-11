FROM golang:1.24

WORKDIR /app

COPY . .

RUN go build -o auth-service

EXPOSE 8080

CMD ["./gateway-service"]
