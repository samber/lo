
FROM golang:1.22.6

WORKDIR /go/src/github.com/samber/lo

COPY Makefile go.* ./

RUN make tools
