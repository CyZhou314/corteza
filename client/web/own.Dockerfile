# build-stage
FROM node:16-alpine as build-stage

ENV PATH /app/node_modules/.bin:$PATH

# build the one
WORKDIR /app/one

RUN apk update && apk add --no-cache git

COPY ./one/package.json ./
COPY ./one/yarn.lock ./

RUN yarn install

COPY ./one/ ./

RUN yarn --non-interactive --no-progress --silent --emoji false build --production

# build the admin
WORKDIR /app/admin

COPY ./admin/package.json ./
COPY ./admin/yarn.lock ./

RUN yarn install

COPY ./admin/ ./

RUN yarn build

# build the compose
WORKDIR /app/compose

COPY ./compose/package.json ./
COPY ./compose/yarn.lock ./

RUN yarn install

COPY ./compose/ ./

RUN yarn build

# build the discovery
WORKDIR /app/discovery

COPY ./discovery/package.json ./
COPY ./discovery/yarn.lock ./

RUN yarn install

COPY ./discovery/ ./

RUN yarn build

# build the privacy
WORKDIR /app/privacy

COPY ./privacy/package.json ./
COPY ./privacy/yarn.lock ./

RUN yarn install

COPY ./privacy/ ./

RUN yarn build

# build the reporter
WORKDIR /app/reporter

COPY ./reporter/package.json ./
COPY ./reporter/yarn.lock ./

RUN yarn install

COPY ./reporter/ ./

RUN yarn build

# build the workflow
WORKDIR /app/workflow

COPY ./workflow/package.json ./
COPY ./workflow/yarn.lock ./

RUN yarn install

COPY ./workflow/ ./

RUN yarn build

# add the applications
WORKDIR /app/applications
COPY ./applications/ ./


# deploy stage
FROM nginx:stable-alpine

WORKDIR /usr/share/nginx/html

COPY --from=build-stage /app/one/dist /usr/share/nginx/html
COPY --from=build-stage /app/admin/dist /usr/share/nginx/html/admin
COPY --from=build-stage /app/compose/dist /usr/share/nginx/html/compose
COPY --from=build-stage /app/discovery/dist /usr/share/nginx/html/discovery
COPY --from=build-stage /app/privacy/dist /usr/share/nginx/html/privacy
COPY --from=build-stage /app/reporter/dist /usr/share/nginx/html/reporter
COPY --from=build-stage /app/workflow/dist /usr/share/nginx/html/workflow
COPY --from=build-stage /app/applications /usr/share/nginx/html/applications
COPY nginx.conf /etc/nginx/nginx.conf
# COPY CONTRIBUTING.* DCO LICENSE README.* ./
COPY entrypoint.sh /entrypoint.sh

RUN chmod +x /entrypoint.sh

EXPOSE 80

HEALTHCHECK --interval=30s --start-period=10s --timeout=30s \
  CMD wget --quiet --tries=1 --spider "http://127.0.0.1:80/config.js" || exit 1

ENTRYPOINT ["/entrypoint.sh"]
