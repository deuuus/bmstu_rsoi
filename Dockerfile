FROM postgres:latest
COPY /postgres/10-create-user-and-db.sql /docker-entrypoint-initdb.d/
CMD ["postgres"]

FROM golang:1.18.4-alpine3.16
WORKDIR /app
COPY ./server .
RUN go build -o main ./cmd
EXPOSE 8080

CMD ["./main"]