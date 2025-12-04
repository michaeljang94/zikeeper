#!/bin/bash

docker compose -f docker-compose-certbot.yaml run --rm certbot certonly --webroot --webroot-path /var/www/certbot/  -d zikeeper.com -d api.zikeeper.com -d www.zikeeper.com