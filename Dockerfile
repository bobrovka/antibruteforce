FROM golang:1.13 as modules

ADD go.mod go.sum /m/
RUN cd /m && go mod download


FROM golang:1.13 AS build

COPY --from=modules /go/pkg /go/pkg
COPY . /code
WORKDIR /code
RUN make build


FROM alpine:3.11

RUN apk add --no-cache bash
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
COPY ./scripts/wait-for-it.sh /scripts/wait-for-it.sh
COPY --from=build /code/bin/abf /app/abf
