FROM golang:latest

ENV GO111MODULE on

RUN mkdir /test

COPY ./sever.go /test

CMD ["go", "run", "/test/server.go"]