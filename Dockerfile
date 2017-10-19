FROM golang:1.9.1 AS build-env

ADD main.go /go/src/github.com/woremacx/go-dump-server/
WORKDIR /go/src/github.com/woremacx/go-dump-server
RUN go get -v \
 && go build -v

FROM alpine:3.6
# glibc
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

RUN mkdir -p /usr/local/bin
COPY --from=build-env /go/src/github.com/woremacx/go-dump-server/go-dump-server /usr/local/bin

CMD ["/usr/local/bin/go-dump-server"]
