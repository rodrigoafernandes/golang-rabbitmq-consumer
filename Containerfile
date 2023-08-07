FROM docker.io/library/golang:1.19-bullseye as build
WORKDIR /go/src
COPY . /go/src/
ENV CGO_ENABLED=0
RUN go mod tidy && \
    go build -o golang-rabbitmq-consumer

FROM scratch
WORKDIR /usr/local/bin
COPY --from=build /go/src/golang-rabbitmq-consumer /usr/local/bin/golang-rabbitmq-consumer
ENTRYPOINT ["./golang-rabbitmq-consumer"]