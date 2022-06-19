FROM golang:1.18-alpine as builder

ENV GOPATH=/root
ENV PATH="/usr/local/go/bin:${PATH}"

WORKDIR /go

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go mod vendor

RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/currency-alerter

FROM alpine:3.16.0
WORKDIR /app

COPY --from=builder /go/bin/currency-alerter ./

ENTRYPOINT ["/app/currency-alerter"]