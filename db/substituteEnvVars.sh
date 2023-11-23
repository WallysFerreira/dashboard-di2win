#!/bin/bash
set -xe

apt-get update
apt-get install gettext-base

envsubst < /dashboard/src/environments/environment.ts > ./src/environments/environment.ts.tmp
mv ./src/environments/environment.ts.tmp ./src/environments/environment.ts

exec "$@"