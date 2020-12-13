FROM golang:1.14

COPY ./entrypoint.sh /entrypoint.sh

RUN go get -d -v github.com/Pawelek242/home_users-api

WORKDIR /go/src/github.com/Pawelek242/home_users-api

CMD ["/entrypoint.sh"]
