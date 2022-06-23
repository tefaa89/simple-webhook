FROM golang:1.18

WORKDIR "/go/src/simple-webook/"
COPY * .
RUN go build -o simple-webook

EXPOSE 8080
CMD ["./simple-webook"]