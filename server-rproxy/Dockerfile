FROM golang:1.7

WORKDIR /go/src/github.com/dkassab/go-calc-service/server-rproxy
COPY . .
RUN go get -v

CMD ["server-rproxy"]