FROM php:8.2-rc-apache-buster AS phpbase 
WORKDIR /var/www/html
RUN docker-php-ext-install mysqli

# This lets us cache the install of mysqli for subsequent rebuilds
FROM phpbase

# These will be overwritten by --build-args
ARG db_container="localhost"
ARG db_password="root"
ARG db_user="root"
ARG db_port="3066"
ARG db="pwpc"

# Environment variables are used by the custom db.php override defined in .docker/php. They tell the container how to connect to the database. 
ENV DOCKER_DB_CONTAINER_NAME=${db_container}
ENV DOCKER_DB_PASS=${db_password}
ENV DOCKER_DB_USER=${db_user}
ENV DOCKER_DB_PORT=${db_port}
ENV DOCKER_DB=${db}

COPY .docker/php/. ./

EXPOSE 80