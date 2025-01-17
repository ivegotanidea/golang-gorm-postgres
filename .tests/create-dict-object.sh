#!/usr/bin/env bash

BASE_URL=http://localhost:8888/api/v1
TELEGRAM_USER_ID=${1:-6794234746}
PASSWORD=${2:-h5sh3d}

echo "Attempting login with TELEGRAM_USER_ID=${TELEGRAM_USER_ID}"

access_token=$(curl -s -X POST ${BASE_URL}/auth/bot/login \
                    -H "Content-Type: application/json" \
                    -d "{\"telegramUserId\": \"${TELEGRAM_USER_ID}\", \"password\": \"${PASSWORD}\"}" | jq -r .access_token)

if [[ -z "$access_token" || "$access_token" == "null" ]]; then
    echo "Failed to retrieve access token. Verify credentials or server logs."
    exit 1
fi

echo "Access Token: $access_token"

set -x
curl --location --request POST "${BASE_URL}/dict/" \
  -H "Authorization: Bearer $access_token" \
  --header 'Content-Type: application/json' \
  --data-raw '{
    "type": "city",
    "data": {
      "name": "New York",
      "aliasRu": "Нью-Йорк",
      "aliasEn": "New York"
    }
  }'