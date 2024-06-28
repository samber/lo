
FROM golang:1.21.10

WORKDIR /go/src/github.com/samber/lo

COPY Makefile go.* ./

RUN make tools
