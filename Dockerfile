
FROM golang:1.18rc1-bullseye

WORKDIR /go/src/github.com/samber/lo

COPY Makefile go.* ./

RUN make tools
