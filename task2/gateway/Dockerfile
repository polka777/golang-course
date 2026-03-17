FROM golang:latest

COPY ./ ./

RUN go build -o gateway ./cmd/main.go

CMD ["./gateway"]