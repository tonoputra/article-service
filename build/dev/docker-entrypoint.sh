#!/bin/bash

export DOCKER_HOST="host.docker.internal"
export LOAD_DOT_ENV="true"

ping -q -c1 $DOCKER_HOST > /dev/null 2>&1
if [ $? -ne 0 ]; then
  HOST_IP=$(ip route | awk 'NR==1 {print $3}')
  echo "${HOST_IP}	${DOCKER_HOST}" | sudo tee -a /etc/hosts > /dev/null
fi

sudo chown -R tono:tono /go/pkg
sudo chown -R tono:tono ./vendor

echo "Run 'go mod vendor', please wait..."
go mod vendor

CompileDaemon \
    -exclude-dir="build vendor deployments scripts" \
    -color=true \
    -graceful-kill=true \
    -pattern="^(\.env.+|\.env)|(.+\.go|.+\.c)$" \
    -build="go build -mod=vendor -o $SERVICE_NAME ./cmd/$SERVICE_NAME/..." \
    -command="./${SERVICE_NAME}"
    