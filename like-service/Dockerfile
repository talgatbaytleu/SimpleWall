FROM golang:1.24

WORKDIR /app

COPY . .

RUN go build -o like-service  

EXPOSE 8080

CMD ["./like-service"]
