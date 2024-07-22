FROM golang:alpine AS builder

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /app

COPY . .

RUN go get -d -v ./...

# Build the Go app
RUN go build -o api .

EXPOSE 8080

CMD ["./api"]