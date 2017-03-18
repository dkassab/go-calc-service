FROM golang:1.7

RUN mkdir -p /go/src/app
WORKDIR /go/src/app
COPY . /vendor/github.com/gengo/grpc-gateway/third_party/googleapis/google/api

COPY . /go/src/app

RUN go install -v

CMD ["app"]