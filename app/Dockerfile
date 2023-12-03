FROM golang:1.20.7 as builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o app

CMD ["./app"]