
FROM golang:1.16-alpine AS builder

RUN apk update && apk add gcc libc-dev make git

WORKDIR /build/
COPY . .
ENV CGO_ENABLED=1 GOOS=linux
RUN make build

FROM alpine:3.15

WORKDIR /

COPY --from=builder /build/disposable-email-cache .

EXPOSE 5000
CMD ["/disposable-email-cache"]