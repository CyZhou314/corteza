# build-stage
FROM node:16.14-alpine as build-stage

ENV PATH /app/node_modules/.bin:$PATH
ENV BUILD_VERSION latest:$APP_VERSION
ENV VUE_APP_VERSION latest:$APP_VERSION

# build the one
WORKDIR /app

RUN apk update && apk add git

RUN npm config set registry https://registry.npm.taobao.org
RUN npm config set disturl https://npm.taobao.org/dist

COPY . .

# build the app
WORKDIR /app/one

RUN yarn install && yarn build

# build the admin
WORKDIR /app/admin

RUN yarn install && yarn build

# build the compose
WORKDIR /app/compose

RUN yarn install && yarn build

# build the discovery
WORKDIR /app/discovery

RUN yarn install && yarn build

# build the privacy
WORKDIR /app/privacy

RUN yarn install && yarn build

# build the reporter
WORKDIR /app/reporter

RUN yarn install && yarn build

# build the workflow
WORKDIR /app/workflow

RUN yarn install && yarn build

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
