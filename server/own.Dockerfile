# bundle web-console
FROM node:16.14-alpine as webconsole-build-stage

WORKDIR /webconsole
COPY ./webconsole ./
# Snapshot is built in development mode and with source map
# RUN yarn install && yarn build --mode dev --sourcemap
RUN yarn install && yarn build


# build server
FROM golang:1.19-buster as server-build-stage

ENV BUILD_OS=linux
ENV BUILD_ARCH=amd64
ENV BUILD_VERSION=latest

WORKDIR /corteza

COPY . ./

COPY --from=webconsole-build-stage /webconsole/dist ./webconsole/dist

WORKDIR /corteza/server/pkg/locale

RUN make src/en src/de src/fr

WORKDIR /corteza

RUN make release-clean release
# RUN go build "-X github.com/CyZhou314/corteza/server/pkg/version.Version=${BUILD_VERSION}" -o build/pkg/corteza-server cmd/corteza/main.go
# RUN go build -o build/corteza-server cmd/corteza/main.go

# deploy stage
FROM ubuntu:20.04

RUN apt-get -y update \
 && apt-get -y install \
    ca-certificates \
 && rm -rf /var/lib/apt/lists/*

ENV STORAGE_PATH "/data"
ENV CORREDOR_ADDR "corredor:80"
ENV HTTP_ADDR "0.0.0.0:80"
ENV HTTP_WEBAPP_ENABLED "false"
ENV PATH "/corteza/bin:${PATH}"

WORKDIR /corteza

VOLUME /data

COPY --from=server-build-stage /corteza/build/pkg/corteza-server ./

HEALTHCHECK --interval=30s --start-period=1m --timeout=30s --retries=3 \
    CMD curl --silent --fail --fail-early http://127.0.0.1:80/healthcheck || exit 1

EXPOSE 80

ENTRYPOINT ["./bin/corteza-server"]

CMD ["serve-api"]
