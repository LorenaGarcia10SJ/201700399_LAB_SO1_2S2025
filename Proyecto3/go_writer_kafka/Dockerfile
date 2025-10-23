FROM golang:1.25.3-alpine
WORKDIR /app
COPY . .
RUN go build -o writer_kafka .
EXPOSE 8082
CMD ["./writer_kafka"]
