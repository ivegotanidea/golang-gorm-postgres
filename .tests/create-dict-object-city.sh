#!/usr/bin/env bash

BASE_URL=http://localhost:8888/api/v1
TELEGRAM_USER_ID=${1:-6794234746}
PASSWORD=${2:-h5sh3d}

echo "Attempting login with TELEGRAM_USER_ID=${TELEGRAM_USER_ID}"

# Login and retrieve access token
response=$(curl -s -w "%{http_code}" -o /tmp/login_response -X POST ${BASE_URL}/auth/bot/login \
            -H "Content-Type: application/json" \
            -d "{\"telegramUserId\": \"${TELEGRAM_USER_ID}\", \"password\": \"${PASSWORD}\"}")

http_code=$(echo $response | tail -n1)
access_token=$(cat /tmp/login_response | jq -r .access_token)

if [[ "$http_code" -ne 200 || -z "$access_token" || "$access_token" == "null" ]]; then
    echo "Failed to retrieve access token. HTTP Code: $http_code"
    cat /tmp/login_response
    exit 1
fi

echo "Access Token: $access_token"

# Create a new profileTag
create_response=$(curl -s -w "%{http_code}" -o /tmp/create_response --location --request POST "${BASE_URL}/dict/" \
  -H "Authorization: Bearer $access_token" \
  --header 'Content-Type: application/json' \
  --data-raw '{
    "type": "profileTag",
    "data": {
      "name": "los angeles",
      "aliasRu": "Лос Анджелес",
      "aliasEn": "Los Angeles",
      "flags": [
        {
          "enabled": "false",
          "aliasRu": "тестовый флаг",
          "aliasEn": "test flag"
        }
      ]
    }
  }')
http_code=$(echo $create_response | tail -n1)
create_body=$(cat /tmp/create_response)
created_profile_tag_id=$create_body

if [[ "$http_code" -ne 201 && "$http_code" -ne 409 ]]; then
    echo "Failed to create profileTag. HTTP Code: $http_code"
    echo "Response: $create_body"
    exit 1
fi

echo "profileTag created with ID: $created_profile_tag_id"

# Modify the created profileTag
modify_response=$(curl -s -w "%{http_code}" -o /tmp/modify_response --location --request PUT "${BASE_URL}/dict/" \
  -H "Authorization: Bearer $access_token" \
  --header 'Content-Type: application/json' \
  --data-raw "{
    \"type\": \"profileTag\",
    \"id\": $created_profile_tag_id,
    \"data\": {
      \"name\": \"los angeles\",
      \"aliasRu\": \"LA City\",
      \"aliasEn\": \"LAC\"
    }
  }")
http_code=$(echo $modify_response | tail -n1)
modify_body=$(cat /tmp/modify_response)

if [[ "$http_code" -ne 200 && "$http_code" -ne 409 ]]; then
    echo "Failed to modify profileTag HTTP Code: $http_code"
    echo "Response: $modify_body"
    exit 1
fi

echo "City modified successfully."

# Verify modification
verify_response=$(curl -s -w "%{http_code}" -o /tmp/verify_response --location --request GET "${BASE_URL}/dict/?type=profileTag&page=1&limit=500" \
  -H "Authorization: Bearer $access_token" \
  --header 'Content-Type: application/json')

http_code=$(echo $verify_response | tail -n1)
verify_body=$(cat /tmp/verify_response)

modified_profileTag=$(echo "$verify_body" | jq ".data[] | select(.id == $created_profile_tag_id)")

if [[ "$http_code" -ne 200 || -z $modified_profileTag ]]; then
    echo "Failed to verify modified profileTag HTTP Code: $http_code"
    echo "Response: $verify_body"
    exit 1
fi

echo "Verified modified profileTag: $modified_profileTag"

# Delete the profileTag
delete_response=$(curl -s -w "%{http_code}" -o /tmp/delete_response \
  --location --request DELETE "${BASE_URL}/dict/profileTag/$created_profile_tag_id" \
  -H "Authorization: Bearer $access_token" \
  --header 'Content-Type: application/json')

http_code=$(echo $delete_response | tail -n1)
delete_body=$(cat /tmp/delete_response)

if [[ "$http_code" -ne 200 || $(echo "$delete_body" | jq -r '.status') != "success" ]]; then
    echo "Failed to delete profileTag HTTP Code: $http_code"
    echo "Response: $delete_body"
    exit 1
fi

echo "City deleted successfully."

# Verify deletion
verify_delete_response=$(curl -s -w "%{http_code}" -o /tmp/verify_delete_response \
  --location --request GET "${BASE_URL}/dict/?type=profileTag" \
  -H "Authorization: Bearer $access_token" \
  --header 'Content-Type: application/json')
http_code=$(echo $verify_delete_response | tail -n1)
verify_delete_body=$(cat /tmp/verify_delete_response)

deleted_profile_tag=$(echo "$verify_delete_body" | jq ".data[] | select(.id == $created_profile_tag_id)")

if [[ "$http_code" -ne 200 || -n "deleted_profile_tag" ]]; then
    echo "City still exists after deletion. HTTP Code: $http_code"
    echo "Response: $verify_delete_body"
    exit 1
fi

echo "Deletion verified successfully. City no longer exists."