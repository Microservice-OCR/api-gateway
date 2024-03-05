FROM golang:1.21.1-alpine AS builder
RUN apk update && apk add --no-cache git

WORKDIR /go/gateway
COPY ./api-gateway .

RUN go install
RUN go build -o /go/gateway/bin/gateway

FROM scratch
COPY --from=builder /go/gateway/bin/gateway .
COPY --from=builder /go/gateway/logs ./logs
CMD ["/gateway"]