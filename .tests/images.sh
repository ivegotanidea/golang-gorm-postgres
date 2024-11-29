#!/usr/bin/env bash

BASE_URL=http://s2t5zmt3q3kibulra383.xyz/api/v1
TELEGRAM_USER_ID=${1:-6794234746}
PASSWORD=${2:-h5sh3d}
PROFILE_ID=${3:-7e852ab6-878f-4f5b-84d2-2ee2749bfcee}

IMAGES=(
    "/Users/orhanmamedov/Downloads/img_original.png"
    "/Users/orhanmamedov/Downloads/img_original_new_1.jpeg"
)

set -x

access_token=$(curl -s -X POST ${BASE_URL}/auth/bot/login \
                    -H "Content-Type: application/json" \
                    -d "{\"telegramUserId\": \"${TELEGRAM_USER_ID}\", \"password\": \"${PASSWORD}\"}" | jq -r .access_token)

curl_cmd="curl -X POST ${BASE_URL}/images -F \"profileID=${PROFILE_ID}\" -b \"access_token=$access_token\""

for img_path in "${IMAGES[@]}"; do
    curl_cmd+=" -F \"images=@${img_path}\""
done

eval $curl_cmd