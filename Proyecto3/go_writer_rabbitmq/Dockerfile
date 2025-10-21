FROM golang:1.22-alpine
WORKDIR /app
COPY . .
RUN go build -o writer .
EXPOSE 9200
CMD ["./writer"]
