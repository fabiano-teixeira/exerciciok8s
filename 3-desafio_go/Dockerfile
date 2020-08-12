FROM golang:1.14.6-alpine as builder

ENV GOPATH /go

RUN apk update && apk add --no-cache git & mkdir $GOPATH/src/greeting/

WORKDIR $GOPATH/src/greeting/

COPY ./server.go ./server_test.go ./

RUN go get -d -v
RUN go build -o $GOPATH/src/greeting/webserver-go-greeting

FROM scratch AS final

WORKDIR /app

COPY --from=builder /go/src/greeting/webserver-go-greeting .

USER 1000

EXPOSE 8000

ENTRYPOINT ["./webserver-go-greeting"]
