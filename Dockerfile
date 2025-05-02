FROM debian:bookworm

WORKDIR /app

RUN apt-get update && apt upgrade -y
RUN apt-get install -y \
    git

RUN /bin/bash -c "$(curl -fsSL https://php.new/install/linux/8.4)"

RUN composer global require laravel/installer

# FROM php:8.4.6-bookworm

# WORKDIR /app

# RUN apt-get update && apt upgrade -y && apt-get install -y \
#     git

# RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer

# #CMD [ "executable" ]