FROM haifahrul/go-1.16-alpine:latest

ENV SERVICE_NAME article-service
ENV REPOSITORY https://github.com/tonoputra/
ENV GOPRIVATE https://github.com/tonoputra/*

COPY build/assets/sudoers /etc/sudoers

RUN chmod 0400 /etc/sudoers && \
    addgroup tono -g 1000 && \
    addgroup sudo && \
    adduser -h /home/tono -s /bin/sh -u 1000 -G users -D tono && \
    adduser tono tono && \
    adduser tono sudo

USER tono

RUN sudo chown -R tono:tono /go

WORKDIR "/go/src/${REPOSITORY}/${SERVICE_NAME}"

ENTRYPOINT ["sh", "./build/dev/docker-entrypoint.sh"]
