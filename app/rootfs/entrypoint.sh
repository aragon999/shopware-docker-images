#!/usr/bin/env sh

sed -i "s;__DOCUMENT_ROOT__;${APP_DOCUMENT_ROOT};" /etc/nginx/sites-enabled/www.conf

exec /usr/bin/supervisord -c /etc/supervisord.conf
