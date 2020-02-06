FROM node:10-alpine AS node-builder

RUN apk update && \
	apk add make

RUN npm install uglify-js -g

COPY static /app/static

RUN cd /app/static && \
	make

FROM golang:1.13 AS go-builder

WORKDIR /go/src/github.com/cyber-republic/go-playground/webapp/

COPY .. .

RUN go install github.com/cyber-republic/go-playground/webapp

FROM debian:buster

ENV HOME_DIR="/home/elauser"

RUN apt-get update && \
  apt-get install -y ca-certificates

COPY --from=go-builder /go/bin/webapp /usr/local/bin/webapp
COPY --from=node-builder /app/static /var/app/static

RUN mkdir -p $HOME_DIR && \
    useradd -d $HOME_DIR elauser && \
    chown -R elauser:elauser $HOME_DIR /var/app && \
    setcap CAP_NET_BIND_SERVICE=+eip /usr/local/bin/webapp

USER elauser

WORKDIR /var/app

EXPOSE 80

CMD [ "webapp", "-allow-share" ]
