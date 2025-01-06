#!/usr/bin/env bash

BASE_URL=http://s2t5zmt3q3kibulra383.xyz/api/v1
TELEGRAM_USER_ID=${1:-6794234746}
PASSWORD=${2:-h5sh3d}
PROFILE_ID=${3:-7e852ab6-878f-4f5b-84d2-2ee2749bfcee}

access_token=$(curl -s -X POST ${BASE_URL}/auth/bot/login \
                    -H "Content-Type: application/json" \
                    -d "{\"telegramUserId\": \"${TELEGRAM_USER_ID}\", \"password\": \"${PASSWORD}\"}" | jq -r .access_token)
set -x

curl_cmd="curl -X GET ${BASE_URL}/profiles/all\?page\=1\&limit\=10\&city\=1 -H \"Authorization: Bearer $access_token\""

eval $curl_cmd