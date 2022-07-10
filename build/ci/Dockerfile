##
## Build
##

FROM golang:1.16-alpine AS build

ENV SERVICE_NAME article-service
ENV REPOSITORY https://github.com/tonoputra/
ENV GO111MODULE on
ENV GOPRIVATE https://github.com/tonoputra/*

RUN apk update && apk add --no-cache git

WORKDIR "/go/src/${REPOSITORY}/${SERVICE_NAME}"

COPY . .

RUN mv .netrc ~/.netrc && \
    chmod 600 ~/.netrc
RUN go mod vendor
RUN go build -o /usr/bin/${SERVICE_NAME} ./cmd/${SERVICE_NAME}/...

##
## Deploy
##

FROM alpine

ENV SERVICE_NAME article-service
ENV REPOSITORY https://github.com/tonoputra/

COPY --from=build /usr/bin/${SERVICE_NAME} /usr/bin/${SERVICE_NAME}
COPY --from=build /go/src/${REPOSITORY}/${SERVICE_NAME}/rds-combined-ca-bundle.pem /usr/bin/

RUN adduser -D -g '' tono

USER tono

ENTRYPOINT /usr/bin/${SERVICE_NAME}
