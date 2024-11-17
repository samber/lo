
FROM golang:1.23.3

WORKDIR /go/src/github.com/samber/lo

COPY Makefile go.* ./

RUN make tools
