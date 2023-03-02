FROM golang AS builder
WORKDIR /go/src/github.com/Krokin/gocloudcamp/
ADD . .
RUN cd Part_Two/cmd/server && CGO_ENABLED=0 GOOS=linux go build -o ../../../server


FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/github.com/Krokin/gocloudcamp/server ./
COPY --from=builder /go/src/github.com/Krokin/gocloudcamp/config.yaml ./
CMD ["./server"]