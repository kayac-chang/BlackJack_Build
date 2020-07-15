FROM node:alpine AS build

WORKDIR /app

COPY ./client .
COPY .env .

RUN apk add --no-cache python make g++

RUN npm ci && npm run build

FROM nginx:alpine

COPY --from=build /app/dist /usr/share/nginx/html
COPY ./client/locales /usr/share/nginx/html/locales

COPY ./nginx/nginx.conf /etc/nginx/nginx.conf
COPY ./nginx/conf.d /etc/nginx/conf.d
COPY ./nginx/include.d /etc/nginx/include.d
COPY ./nginx/sites-enabled.d/client.conf /etc/nginx/sites-enabled.d/client.conf