FROM php:8.4.6-bookworm

WORKDIR /app

RUN apt-get update && apt upgrade -y

RUN apt-get install -y curl procps git zip

ENV TERM=xterm

#RUN /bin/bash -c "$(curl -fsSL https://php.new/install/linux/8.4)"

# Install php dependencies
RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer

RUN composer global require laravel/installer

# Install node
ENV DEBIAN_FRONTEND=noninteractive
ENV NODE_VERSION=18.17.1

RUN apt-get update && \
    apt-get install -y curl ca-certificates build-essential && \
    curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.7/install.sh | bash && \
    . "$HOME/.nvm/nvm.sh" && \
    nvm install $NODE_VERSION && \
    nvm use $NODE_VERSION \
    nvm alias default $NODE_VERSION && \
    cp "$HOME/.nvm/versions/node/v$NODE_VERSION/bin/node" /usr/local/bin/node && \
    cp "$HOME/.nvm/versions/node/v$NODE_VERSION/bin/npm" /usr/local/bin/npm && \
    cp "$HOME/.nvm/versions/node/v$NODE_VERSION/bin/npx" /usr/local/bin/npx

RUN npm install

CMD [ "composer", "run", "dev" ]

# FROM php:8.4.6-bookworm

# WORKDIR /app

# RUN apt-get update && apt upgrade -y && apt-get install -y \
#     git

# RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer

# #CMD [ "executable" ]