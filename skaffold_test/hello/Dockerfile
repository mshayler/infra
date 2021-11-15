FROM golang:1.12.9-alpine3.10 as builder

COPY cmd/main.go .

RUN go build -o /testapp main.go

# ---

FROM alpine:3.12.1

ENV BUILD_PACKAGES \
  bash

# General packages
RUN set -x && \
    \
    echo "==> Upgrading apk and system..."  && \
    apk update && apk upgrade && \
    \
    echo "==> Adding packages..."  && \
    apk add --no-cache ${BUILD_PACKAGES} && \
    rm -rf /var/cache/apk/*

COPY --from=builder /testapp /usr/local/bin

RUN chmod +x /usr/local/bin/testapp

EXPOSE 8081

ENTRYPOINT ["testapp"]
