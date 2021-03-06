#!/bin/bash

# update
apt-get update -y

# install latest version of docker the lazy way
curl -sSL https://get.docker.com | sh

# make it so you don't need to sudo to run docker commands
usermod -aG docker ubuntu

# install docker-compose
curl -L https://github.com/docker/compose/releases/download/1.21.2/docker-compose-$(uname -s)-$(uname -m) -o /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose

# copy the docker-compose into /srv/docker
# if you change this, change the systemd service file to match
# WorkingDirectory=[whatever you have below]
mkdir /srv/docker
curl -o /srv/docker/docker-compose.yml https://raw.githubusercontent.com/cyber-republic/go-playground/master/webapp/prod/docker-compose.yml

# copy .env.sample file
curl -o /srv/docker/.env https://raw.githubusercontent.com/cyber-republic/go-playground/master/webapp/prod/.env.sample

# download the latest version of nginx.tmpl
curl -o /srv/docker/nginx.tmpl https://raw.githubusercontent.com/jwilder/nginx-proxy/master/nginx.tmpl
# copy in systemd unit file and register it so our compose file runs
# on system restart
curl -o /etc/systemd/system/docker-compose-app.service https://raw.githubusercontent.com/cyber-republic/go-playground/master/webapp/prod/docker-compose-app.service
systemctl enable docker-compose-app

# start docker and enable docker
systemctl start docker
systemctl enable docker

# start up the application via docker-compose
cd /srv/docker
source .env
docker network create $NETWORK $NETWORK_OPTIONS
docker-compose up -d
