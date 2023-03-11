# build-stage
FROM cyzhou314/corteza-webapp as webapp
FROM cyzhou314/corteza-server as server
FROM ubuntu:20.04

RUN apt-get -y update \
 && apt-get -y install \
    ca-certificates \
    curl \
 && rm -rf /var/lib/apt/lists/*

ENV STORAGE_PATH "/data"
ENV CORREDOR_ADDR "corredor:80"
ENV HTTP_ADDR "0.0.0.0:80"
ENV HTTP_WEBAPP_ENABLED "true"
ENV HTTP_WEBAPP_BASE_DIR "/corteza/webapp"
ENV PATH "/corteza/bin:${PATH}"

WORKDIR /corteza

VOLUME /data

# COPY --from=build-stage /corteza ./
RUN rm -rf ./webapp
COPY --from=webapp /usr/share/nginx/html ./webapp
COPY --from=server /corteza ./

# HEALTHCHECK --interval=30s --start-period=1m --timeout=30s --retries=3 \
#     CMD curl --silent --fail --fail-early http://127.0.0.1:80/healthcheck || exit 1

EXPOSE 80

ENTRYPOINT ["./bin/corteza-server"]

CMD ["serve-api"]
