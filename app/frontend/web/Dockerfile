FROM nginx:1.13.8-alpine

WORKDIR /etc/nginx
COPY /build html/
COPY nginx.conf nginx.conf
RUN nginx -t
