##
## Test
##

FROM golang:1.16

ENV SERVICE_NAME article-service
ENV REPOSITORY https://github.com/tonoputra/

WORKDIR "/go/src/${REPOSITORY}/${SERVICE_NAME}"

COPY . .
RUN cd cmd/${SERVICE_NAME} && go test -v
