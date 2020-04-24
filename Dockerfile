FROM golang:1.14.1

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go build -v -o app

CMD ["./app"]
