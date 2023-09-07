# syntax=docker/dockerfile:1

FROM golang:1.21

WORKDIR /app

COPY . .
RUN go mod download

WORKDIR /app/cmd/server
RUN GOOS=linux go build -o memcache

EXPOSE 8081

CMD ["./memcache"]

# docker build -f ./build/Dockerfile -t memcache .
# docker run -p 8081:8081 memcache
