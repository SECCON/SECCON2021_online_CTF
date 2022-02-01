FROM alpine:3

RUN apk add apache2
WORKDIR /var/www/localhost/htdocs/
ADD app/index.html  .
ADD app/flag.min.js .
ADD app/flag.wasm   .

ADD app/init.sh /etc/init.sh
RUN chmod +x /etc/init.sh
