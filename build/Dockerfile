FROM golang:1.17-alpine3.15 as builder
RUN mkdir /resources
ADD . /resources/
WORKDIR /resources

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN go mod download
RUN go build -a -installsuffix cgo -o runner cmd/api/main.go
RUN ls

FROM alpine:3.15 as cacerts
RUN apk --no-cache add ca-certificates

FROM scratch
COPY --from=cacerts  /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /resources/runner .
CMD ["./runner"]