
FROM golang:1.21.11

WORKDIR /go/src/github.com/samber/lo

COPY Makefile go.* ./

RUN make tools
