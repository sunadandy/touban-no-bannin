FROM golang:latest

WORKDIR /go/src/app

COPY . /go/src/app/

RUN apt-get update
RUN apt-get install -y vim
RUN go mod tidy
RUN go build

CMD [ "go","run","main.go" ]