FROM golang:1.22-alpine
WORKDIR /app
COPY . .
RUN go build -o go_client .
EXPOSE 9000
CMD ["./go_client"]
