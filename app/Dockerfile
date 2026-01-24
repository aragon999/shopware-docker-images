ARG PHP_VERSION=8.4

FROM php:${PHP_VERSION}-fpm-alpine

LABEL org.opencontainers.image.source https://github.com/aragon999/shopware-docker

COPY --from=ghcr.io/shyim/supervisord /usr/local/bin/supervisord /usr/bin/supervisord
COPY --from=ghcr.io/shyim/shopware-docker/sendmail /sendmail /usr/bin/sendmail
COPY --from=mlocati/php-extension-installer /usr/bin/install-php-extensions /usr/bin/

ENV TZ=Europe/Berlin \
    FPM_PM=dynamic \
    FPM_PM_MAX_CHILDREN=5 \
    FPM_PM_START_SERVERS=2 \
    FPM_PM_MIN_SPARE_SERVERS=1 \
    FPM_PM_MAX_SPARE_SERVERS=3 \
    PHP_MAX_UPLOAD_SIZE=128m \
    PHP_MAX_EXECUTION_TIME=300 \
    PHP_MEMORY_LIMIT=512m \
    APP_DOCUMENT_ROOT=/var/www/html/public \
    SMTP_ADDRESS=smtp:1025

RUN set -x && \
    apk add --no-cache \
      nginx \
      shadow \
      gnu-libiconv \
      icu-data-full \
      socat && \
    install-php-extensions bcmath gd intl mysqli pdo_mysql pdo_pgsql sockets bz2 gmp soap zip opcache redis calendar apcu ffi && \
    ln -s /usr/local/bin/php /usr/bin/php && \
    ln -sf /dev/stdout /var/log/nginx/access.log && \
    ln -sf /dev/stderr /var/log/nginx/error.log && \
    rm -rf /var/lib/nginx/tmp && \
    ln -sf /tmp /var/lib/nginx/tmp && \
    mkdir -p /var/tmp/nginx/ || true && \
    chown -R www-data:www-data /var/lib/nginx /var/tmp/nginx/ && \
    chmod 777 -R /var/tmp/nginx/ && \
    rm -rf /tmp/* && \
    chown -R www-data:www-data /var/www && \
    usermod -u 1000 www-data

COPY rootfs/ /
EXPOSE 80
STOPSIGNAL SIGKILL

ENTRYPOINT ["/entrypoint.sh"]
