FROM golang:1.10 as builder
RUN go get -u github.com/go-swagger/go-swagger/cmd/swagger
WORKDIR /go/src/github.com/laincloud/console-api
COPY . .
RUN go install github.com/laincloud/console-api/gen/cmd/console-server

FROM debian:stretch
WORKDIR /lain/app
COPY --from=builder /go/bin/swagger .
COPY --from=builder /go/bin/console-server .
COPY --from=builder /go/src/github.com/laincloud/console-api/swagger.json .
EXPOSE 8081
CMD ["/lain/app/console-server", "--host=0.0.0.0", "--port=8080"]