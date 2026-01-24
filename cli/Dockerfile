ARG PHP_VERSION=8.4

FROM php:${PHP_VERSION}-cli

# TODO: Correct source label
LABEL org.opencontainers.image.source https://github.com/aragon999/shopware-docker

COPY --from=mlocati/php-extension-installer /usr/bin/install-php-extensions /usr/bin/
COPY --from=ghcr.io/shyim/shopware-docker/sendmail /sendmail /usr/bin/sendmail
COPY --from=ghcr.io/friendsofshopware/shopware-cli /usr/local/bin/shopware-cli /usr/local/bin/shopware-cli
COPY --from=composer:2 /usr/bin/composer /usr/bin/composer

RUN set -x && \
    apt-get update && \
    apt-get install -y --no-install-recommends \
        unzip \
        bash \
        bash-completion \
        wget \
        git \
        zip \
        patch \
        mariadb-client \
        jq \
        make && \
    rm -rf /var/lib/apt/lists/* && \
    install-php-extensions bcmath gd intl mysqli pdo_mysql pdo_pgsql sockets bz2 gmp soap zip gmp pcntl posix redis pcov xsl calendar apcu ffi opcache && \
    useradd dev && \
    mkdir /home/dev && \
    chown -R dev /home/dev /opt/ && \
    cd /tmp && \
    wget https://github.com/Schniz/fnm/releases/download/v1.38.1/fnm-linux.zip -O fnm.zip && \
    unzip fnm.zip && chmod +x fnm && mv fnm /usr/local/bin/fnm && \
    wget https://github.com/mikefarah/yq/releases/latest/download/yq_linux_amd64 -O /usr/local/bin/yq && \
    chmod +x /usr/local/bin/yq && \
    rm -rf /tmp/*

USER dev

ENV HOME /home/dev
ENV PS1 '${debian_chroot:+($debian_chroot)}\[\033[01;31m\]\u\[\033[01;33m\]@\[\033[01;36m\]\h \[\033[01;33m\]\w \[\033[01;35m\]\$ \[\033[00m\]'
ENV SMTP_ADDRESS=smtp:1025

RUN echo "alias ll=\"ls -lha\"" >> ~/.bashrc && \
    echo 'eval $(fnm --fnm-dir=/nvm/.debian-fnm env --shell=bash)' >> ~/.bashrc

COPY rootfs/ /
WORKDIR /var/www/html
STOPSIGNAL SIGKILL

ENTRYPOINT [ "/entrypoint.sh" ]
