FROM golang:1.18-alpine3.15 AS builder
ARG GO_MODULE="github.com/discomco/go-cart/robby"
ARG PA_CONTEXT="execute-game"
ARG CQRS_SERVICE="command"

WORKDIR /usr/src/app
COPY . /usr/src/app
RUN go install -v ${GO_MODULE}/${PA_CONTEXT}/${CQRS_SERVICE}/cmd

FROM alpine:3.15
ARG CQRS_SERVICE=""

COPY --from=builder /go/bin/cmd /usr/bin/${CQRS_SERVICE}

RUN echo "#!/bin/sh\n" > /entrypoint.sh \
    && echo "/usr/bin/${CQRS_SERVICE} $@\n" >> /entrypoint.sh \
    && chmod ugo+x /entrypoint.sh

ENTRYPOINT ["/bin/sh", "/entrypoint.sh"]