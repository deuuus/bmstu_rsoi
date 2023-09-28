FROM golang:1.18.4-alpine3.16

WORKDIR /app

COPY ./server .

RUN go build -o main ./cmd

EXPOSE 8080

CMD ["./main"]