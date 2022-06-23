FROM golang:1.18

WORKDIR "/Users/mostafa.abdelkhalek/Development/go/src/simple_webhook_service"

COPY go.mod .
COPY main.go .

RUN go build -o service


EXPOSE 8080
CMD ["./service"]