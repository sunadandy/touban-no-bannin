FROM golang:latest

WORKDIR /go/src/app

COPY . /go/src/app/

RUN go mod tidy
RUN go build

CMD [ "go","run","main.go" ]