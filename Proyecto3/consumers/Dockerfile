FROM golang:1.22-alpine
WORKDIR /app
COPY . .
RUN go build -o consumer .
EXPOSE 9300
CMD ["./consumer"]
