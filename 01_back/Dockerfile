FROM golang:1.21.1

WORKDIR /go/src/app

COPY . /go/src/app/
# npm run buildで生成した静的コンテンツを予めコピーしておく
COPY ./dist /go/src/app/dist

# RUN apt update
# RUN apt install -y vim
RUN go mod tidy

CMD [ "go","run","main.go" ]