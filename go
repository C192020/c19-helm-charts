#!/bin/bash

target_build() {
  cd letsencrypt-issuer
  helm package .
  cd ..
  helm repo index .
}


if type -t "target_$1" &>/dev/null; then
  target_$1 ${@:2}
else
  echo "usage: $0 <target>

target:
    build      - packages and index the helm charts
"
  exit 1
fi