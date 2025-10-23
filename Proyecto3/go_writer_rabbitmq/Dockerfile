FROM golang:1.25.3-alpine
WORKDIR /app
COPY . .
RUN go build -o writer_rabbitmq .
EXPOSE 8081
CMD ["./writer_rabbitmq"]
