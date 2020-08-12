FROM golang:1.14.6-alpine as builder

ENV GOPATH /go

RUN apk update && apk add --no-cache git & mkdir $GOPATH/src/greeting/

WORKDIR $GOPATH/src/greeting/

COPY ./server.go ./server_test.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o $GOPATH/src/greeting/webserver-go-greeting .

FROM scratch

COPY --from=builder /go/src/greeting/webserver-go-greeting /app/
WORKDIR /app
USER 1000
EXPOSE 8000

ENTRYPOINT ["./webserver-go-greeting"]
