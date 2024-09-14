FROM golang

WORKDIR /go/src/app

ADD . /go/src/app

RUN go get -u github.com/gin-gonic/gin@v1.10.0
RUN go build -o /go/bin/app

ENTRYPOINT /go/bin/app

ENV PORT 8080
ENV GIN_MODE release

EXPOSE 8080