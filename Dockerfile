FROM golang:1.18

WORKDIR "/go/src/simple-webhook/"
COPY * .
RUN go build -o simple-webhook

EXPOSE 8080
CMD ["./simple-webhook"]