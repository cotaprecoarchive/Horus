#!/bin/sh
set -e

has_command() {
  command -v "$@" >/dev/null 2>&1
}

declare -r GHR="https://github.com/CotaPreco/Horus/releases/download/v0.1.1"
declare -r TAR_URL=$GHR/horus-`uname -s`-`uname -m`.tar.gz
declare -r DOWNLOAD_TO="/tmp/Horus-latest.tar.gz"

if has_command wget; then
  curl --silent -o $DOWNLOAD_TO -L $TAR_URL
elif has_command curl; then
  wget --quiet -O $DOWNLOAD_TO $TAR_URL
fi

tar -zxf $DOWNLOAD_TO -C /tmp

MV_CHMODX="mv /tmp/horus /usr/local/bin/horus && chmod +x /usr/local/bin/horus && rm -f $DOWNLOAD_TO"

if has_command sudo; then
  sudo -E sh -c "$MV_CHMODX"
elif has_command su; then
  su -c "$MV_CHMODX"
fi

echo "Done!"
horus -v
