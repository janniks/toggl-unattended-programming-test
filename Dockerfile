FROM golang:1.14

WORKDIR /go/src/app
COPY . .

ENV GIN_MODE=release

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["initial"]
