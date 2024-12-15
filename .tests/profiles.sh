#!/usr/bin/env bash

BASE_URL=http://s2t5zmt3q3kibulra383.xyz/api/v1
TELEGRAM_USER_ID=${1:-6794234746}
PASSWORD=${2:-h5sh3d}

# Function to retrieve a new access token
get_access_token() {
  curl -s -X POST "${BASE_URL}/auth/bot/login" \
       -H "Content-Type: application/json" \
       -d "{\"telegramUserId\": \"${TELEGRAM_USER_ID}\", \"password\": \"${PASSWORD}\"}" | jq -r .access_token
}

# Get the initial access token and timestamp
token_timestamp=$(date +%s)
access_token=$(get_access_token)

# Function to check if more than 10 minutes have passed
is_token_expired() {
  local current_time
  current_time=$(date +%s)
  if (( current_time - token_timestamp > 600 )); then
    return 0 # Token expired
  else
    return 1 # Token still valid
  fi
}

city_count=$(curl -s ${BASE_URL}/dict/cities?limit=500 | jq -r '.results | length')

# iterate cities
for ((city_id=1; city_id<=city_count; city_id++)); do

  page_num=1

  # iterate profiles pages
  while true; do
    if is_token_expired; then
      echo "Access token expired. Retrieving a new one..."
      token_timestamp=$(date +%s)
      access_token=$(get_access_token)
    fi

    echo "Iterating profiles on page $page_num for city $city_id"

    curl_cmd="curl -s -X GET ${BASE_URL}/profiles/list\?page\=${page_num}\&limit\=10\&city\=${city_id} -b \"access_token=$access_token\""

    ids_and_photos=$(eval $curl_cmd | jq -c '.data[] | {id: .id, photos: .photos}')

    if [[ -z $ids_and_photos ]]; then
      echo "Profiles ended at page $page_num for city $city_id"
      break
    fi

    # Iterate over profiles photos
    echo "$ids_and_photos" | jq -c '. | select(.photos != null)' | while read -r profile; do
      id=$(echo "$profile" | jq -r '.id')
      photos=$(echo "$profile" | jq -c '.photos')

      # Flag to track inactive photos
      all_inactive=true

      # Iterate over each photo
      echo "$photos" | jq -c '.[]' | while read -r photo; do
        photo_url=$(echo "$photo" | jq -r '.url')
        photo_disabled=$(echo "$photo" | jq -r '.disabled')

        if [[ $photo_disabled == "true" ]]; then
          continue
        fi

        # Check status code of photo URL
        status_code=$(curl --connect-timeout 5 --max-time 10 -L -s -o /dev/null -w "%{http_code}" "$photo_url")

        if [[ "$status_code" -ne 200 ]]; then
          echo "ID: $id, Photo URL: $photo_url, Status Code: $status_code"
          echo "Should delete photo $photo_url of profile $id"
        else
          all_inactive=false
        fi
      done

      # If all photos are inactive, report profile as inactive
      if [[ "$all_inactive" == "true" ]]; then
        echo "Profile with ID: $id is inactive (all photos are inactive)."

        upd_resp_code=$(curl --connect-timeout 5 --max-time 10 -L -s -o /tmp/response -w "%{http_code}" \
                         -X PUT "${BASE_URL}/profiles/update/${id}" \
                         -H "Content-Type: application/json" \
                         -b "access_token=${access_token}" \
                         -d '{
                               "active": false
                             }')

        if [[ "upd_resp_code" -ne 200 ]]; then
          echo "Can't disable profile $id" && cat /tmp/response && exit 1
        else
          echo "Profile $id has been disabled"
        fi

      fi
    done

    page_num=$((page_num + 1))

  done

done

