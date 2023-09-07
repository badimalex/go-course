# syntax=docker/dockerfile:1

FROM golang:1.21

WORKDIR /app

COPY . .
RUN go mod download

WORKDIR /app/cmd/server
RUN GOOS=linux go build -o shortener

EXPOSE 8080

CMD ["./shortener"]

# docker build -f ./build/Dockerfile -t shortener .
# docker run -p 8080:8080 shortener
