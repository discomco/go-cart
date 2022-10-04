FROM golang:1.18-alpine3.15 AS builder
ARG APP
ARG CR_PAT
RUN apk update && \
    apk add git && \
    apk add openssh

#RUN git config --global credential.helper \
#    '!f() { echo username=rgfaber; echo "password=$CR_PAT"; };f'
#    && \
#    git config --global url.ssh://git@github.com/.insteadOf https://github.com/

WORKDIR /usr/src/app
COPY . /usr/src/app

#RUN go install -v ${GO_MODULE}/${PA_CONTEXT}/${CQRS_SERVICE}/cmd@latest

RUN go mod tidy && \
    go build -o /go/bin/cmd/runme ${APP}/cmd  && \
    cp -r ${APP}/config /go/bin/cmd/


FROM alpine:3.15

VOLUME /usr/bin/config/config.yaml

COPY --from=builder /go/bin/cmd/runme /usr/bin/runme
COPY --from=builder /go/bin/cmd/config /usr/bin

RUN echo "#!/bin/sh \n" > /entrypoint.sh \
    && echo "/usr/bin/runme $@\n" >> /entrypoint.sh \
    && chmod ugo+x /entrypoint.sh

RUN ls -la /usr/bin/config


ENTRYPOINT ["/bin/sh", "/entrypoint.sh"]