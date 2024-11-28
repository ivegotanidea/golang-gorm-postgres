#!/usr/bin/env bash

BASE_URL=http://localhost:80/api/v1
TELEGRAM_USER_ID=${1:-6794234746}
PASSWORD=${2:-h5sh3d}
PROFILE_ID=${3:-123e4567-e89b-12d3-a456-426614174000}

IMAGES=(
    "/Users/orhanmamedov/Downloads/img_original.png"
    "/Users/orhanmamedov/Downloads/img_original_new_1.jpeg"
)

set -x

access_token=$(curl -i -X POST ${BASE_URL}/auth/bot/login \
                    -H "Content-Type: application/json" \
                    -d "{\"telegramUserId\": \"${TELEGRAM_USER_ID}\", \"password\": \"${PASSWORD}\"}" | \
    grep "Set-Cookie" | grep "access_token" | cut -d '=' -f2 | cut -d ';' -f1)

curl_cmd="curl -X POST ${BASE_URL}/images -F \"profileID=${PROFILE_ID}\" -b \"access_token=$access_token\""

for img_path in "${IMAGES[@]}"; do
    curl_cmd+=" -F \"images=@${img_path}\""
done

eval $curl_cmd