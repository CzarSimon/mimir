FROM nginx:1.13.3-alpine

COPY . /usr/src/proxy
WORKDIR /usr/src/proxy
RUN mv nginx.conf /etc/nginx/nginx.conf

RUN cat /etc/nginx/nginx.conf

EXPOSE 80 443

CMD ["nginx", "-g", "daemon off;"]
