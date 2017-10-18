FROM golang:alpine as builder

WORKDIR /go/src/github.com/estk/arbi
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch
COPY --from=builder /go/src/github.com/estk/arbi/main /
EXPOSE 8080
CMD ["/main"]
