FROM golang:1.24

WORKDIR /app

COPY . .

RUN go build -o triple-s 

EXPOSE 8080

CMD ["./triple-s"]
