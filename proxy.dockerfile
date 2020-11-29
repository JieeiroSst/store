FROM nginx:1.17-alpine
COPY ./proxy/default-dev.conf /etc/nginx/conf.d/default.conf