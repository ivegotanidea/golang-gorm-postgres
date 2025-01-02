#!/usr/bin/env bash
# Mock Data Setup
BASE_URL="http://s2t5zmt3q3kibulra383.xyz/api/v1"
PROFILES_URL="${BASE_URL}/profiles"

profile_id="fbfe31ab-eb3d-4146-9e63-c97ec72399ce"

TELEGRAM_USER_ID=${1:-6794234746}
PASSWORD=${2:-h5sh3d}

access_token=$(curl -s -X POST ${BASE_URL}/auth/bot/login \
                    -H "Content-Type: application/json" \
                    -d "{\"telegramUserId\": \"${TELEGRAM_USER_ID}\", \"password\": \"${PASSWORD}\"}" | jq -r .access_token)

curl_cmd="curl -s -X GET ${PROFILES_URL}/${profile_id} -b \"access_token=$access_token\""

ids_and_photos=$(eval $curl_cmd | jq -c '.data | {id: .id, photos: .photos}')

echo "$ids_and_photos" | jq -c '. | select(.photos != null)' | while read -r profile; do
  id=$(echo "$profile" | jq -r '.id')
  photos=$(echo "$profile" | jq -c '.photos')

  # Flag to track inactive photos
  all_inactive=true
  inactive_photo_ids=()

  while read -r photo; do
    photo_id=$(echo "$photo" | jq -r '.id')
    photo_url=$(echo "$photo" | jq -r '.url')
    photo_disabled=$(echo "$photo" | jq -r '.disabled')

    if [[ $photo_disabled == "true" ]]; then
      continue
    fi

    # Check status code of photo URL
    status_code=$(curl --connect-timeout 5 --max-time 10 -L -s -o /dev/null -w "%{http_code}" "$photo_url")

    if [[ "$status_code" -eq 404 ]]; then
      echo "Queuing photo $photo_url (ID: $photo_id) for disabling due to 404 response."
      inactive_photo_ids+=("$photo_id")
      echo "$photo_id has been queued, inactive ids array is:"

      echo "${inactive_photo_ids[@]}"

    else
      echo "Photo $photo_id is active"
      all_inactive=false
    fi
  done < <(echo "$photos" | jq -c '.[]')

  # Disable photos in bulk
  if [[ ${#inactive_photo_ids[@]} -gt 0 ]]; then

    echo "Disabling profile $id photos ${#inactive_photo_ids[@]} of ${#photos[@]}..."

    set -x
    bulk_disable_resp_code=$(curl --connect-timeout 5 --max-time 10 -L -w "%{http_code}" \
                             -X POST "${PROFILES_URL}/${id}/photos" \
                             -H "Content-Type: application/json" \
                             -b "access_token=${access_token}" \
                             -d "$(jq -n --argjson photos "$(printf '%s\n' "${inactive_photo_ids[@]}" | jq -R . | jq -s '[.[] | {id: ., disabled: true}]')" '{photos: $photos}')")

    set +x
    if [[ "$bulk_disable_resp_code" -ne 200 ]]; then
      echo "Failed to disable photos for profile $id" && cat /tmp/response && exit 1
    else
      echo "Disabled photos for profile $id: ${inactive_photo_ids[*]}"
    fi
  fi

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

    if [[ "$upd_resp_code" -ne 200 ]]; then
      echo "Can't disable profile $id" && cat /tmp/response && exit 1
    else
      echo "Profile $id has been disabled."
    fi

  fi
done
