FROM golang:1.12

WORKDIR /

COPY . /go/src/github.com/one-connect/biz-chance-api
RUN go get -u github.com/golang/dep/cmd/dep
RUN cd /go/src/github.com/one-connect/biz-chance-api && go build -o bin/biz-chance-api main.go
RUN cd /go/src/github.com/one-connect/biz-chance-api && cp bin/biz-chance-api /usr/local/bin/

CMD ["biz-chance-api"]