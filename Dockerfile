FROM golang:alpine as builder

WORKDIR /go/src/github.com/estk/arbi
COPY main.go .
RUN go build -o arbi .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/github.com/estk/arbi/arbi .
EXPOSE 8080
CMD ["./arbi"]
