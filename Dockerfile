FROM golang:1.14-alpine as build

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# RUN apk --no-cache add ca-certificates \
#     && apk --no-cache add --virtual build-deps git

COPY ./go.mod ./go.sum ./*.go /go/src/github.com/ymyzk/bq-globalip/
WORKDIR /go/src/github.com/ymyzk/bq-globalip

RUN go build -o /bin/bq-globalip

FROM scratch

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /bin/bq-globalip /usr/bin/bq-globalip

ENTRYPOINT ["/usr/bin/bq-globalip"]
