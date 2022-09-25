FROM golang:1.18-alpine3.15 AS builder
ARG GO_MODULE="github.com/discomco/go-cart/examples"
ARG PA_CONTEXT="quadratic-roots"
ARG CQRS_SERVICE="containers/CMD/monolith"
ARG CR_PAT
RUN apk add git

RUN git config --global user.name "${CR_PAT}" && \
    git config --global url.ssh://git@github.com/.insteadOf https://github.com/

WORKDIR /usr/src/app
COPY . /usr/src/app
RUN go install -v ${GO_MODULE}/${PA_CONTEXT}/${CQRS_SERVICE}/cmd@latest

FROM alpine:3.15
ARG CQRS_SERVICE=""

COPY --from=builder /go/bin/cmd /usr/bin/${CQRS_SERVICE}

RUN echo "#!/bin/sh\n" > /entrypoint.sh \
    && echo "/usr/bin/${CQRS_SERVICE} $@\n" >> /entrypoint.sh \
    && chmod ugo+x /entrypoint.sh

ENTRYPOINT ["/bin/sh", "/entrypoint.sh"]