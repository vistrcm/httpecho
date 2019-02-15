FROM golang:1.11 as builder

WORKDIR /go/src/github.com/vistrcm/httpecho
COPY ./ .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -a -o /httpecho .

FROM scratch
ENV HTTP_PORT=80
COPY --from=builder /httpecho /httpecho
ENTRYPOINT ["/httpecho"]
