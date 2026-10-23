ARG PHP_VERSION=8.4
ARG SUPERVISORD_IMAGE=ghcr.io/aragon999/shopware-docker-images/supervisord:latest

FROM ${SUPERVISORD_IMAGE} AS supervisord

FROM php:${PHP_VERSION}-cli-alpine

LABEL org.opencontainers.image.source https://github.com/aragon999/shopware-docker

COPY --from=supervisord /usr/local/bin/supervisord /usr/bin/supervisord
COPY --from=mlocati/php-extension-installer /usr/bin/install-php-extensions /usr/bin/

ENV TZ=Europe/Berlin \
    HOME=/var/www/html

RUN set -x && \
    apk add --no-cache \
      shadow \
      gnu-libiconv \
      icu-data-full \
      socat && \
    install-php-extensions bcmath gd intl mysqli pdo_mysql pdo_pgsql sockets bz2 gmp soap zip redis calendar apcu ffi && \
    ln -s /usr/local/bin/php /usr/bin/php && \
    rm -rf /tmp/* && \
    chown -R www-data:www-data /var/www && \
    usermod -u 1000 www-data

COPY rootfs/ /

USER www-data
WORKDIR /var/www/html
STOPSIGNAL SIGKILL

ENTRYPOINT ["/entrypoint.sh"]
