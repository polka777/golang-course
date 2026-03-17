FROM golang:latest

COPY ./ ./

RUN go build -o collector ./cmd/main.go

CMD ["./collector"]