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

# Create a new city
create_response=$(curl -s -w "%{http_code}" -o /tmp/create_response --location --request POST "${BASE_URL}/dict/" \
  -H "Authorization: Bearer $access_token" \
  --header 'Content-Type: application/json' \
  --data-raw '{
    "type": "city",
    "data": {
      "name": "los angeles",
      "aliasRu": "Лос Анджелес",
      "aliasEn": "Los Angeles"
    }
  }')
http_code=$(echo $create_response | tail -n1)
create_body=$(cat /tmp/create_response)
created_city_id=$create_body

if [[ "$http_code" -ne 201 && "$http_code" -ne 409 ]]; then
    echo "Failed to create city. HTTP Code: $http_code"
    echo "Response: $create_body"
    exit 1
fi

echo "City created with ID: $created_city_id"

# Modify the created city
modify_response=$(curl -s -w "%{http_code}" -o /tmp/modify_response --location --request PUT "${BASE_URL}/dict/" \
  -H "Authorization: Bearer $access_token" \
  --header 'Content-Type: application/json' \
  --data-raw "{
    \"type\": \"city\",
    \"id\": $created_city_id,
    \"data\": {
      \"name\": \"los angeles\",
      \"aliasRu\": \"LA City\",
      \"aliasEn\": \"LAC\"
    }
  }")
http_code=$(echo $modify_response | tail -n1)
modify_body=$(cat /tmp/modify_response)

if [[ "$http_code" -ne 200 && "$http_code" -ne 409 ]]; then
    echo "Failed to modify city. HTTP Code: $http_code"
    echo "Response: $modify_body"
    exit 1
fi

echo "City modified successfully."

# Verify modification
verify_response=$(curl -s -w "%{http_code}" -o /tmp/verify_response --location --request GET "${BASE_URL}/dict/?type=city&page=1&limit=500" \
  -H "Authorization: Bearer $access_token" \
  --header 'Content-Type: application/json')

http_code=$(echo $verify_response | tail -n1)
verify_body=$(cat /tmp/verify_response)

modified_city=$(echo "$verify_body" | jq ".data[] | select(.id == $created_city_id)")

if [[ "$http_code" -ne 200 || -z $modified_city ]]; then
    echo "Failed to verify modified city. HTTP Code: $http_code"
    echo "Response: $verify_body"
    exit 1
fi

echo "Verified modified city: $modified_city"

# Delete the city
delete_response=$(curl -s -w "%{http_code}" -o /tmp/delete_response \
  --location --request DELETE "${BASE_URL}/dict/city/$created_city_id" \
  -H "Authorization: Bearer $access_token" \
  --header 'Content-Type: application/json')

http_code=$(echo $delete_response | tail -n1)
delete_body=$(cat /tmp/delete_response)

if [[ "$http_code" -ne 200 || $(echo "$delete_body" | jq -r '.status') != "success" ]]; then
    echo "Failed to delete city. HTTP Code: $http_code"
    echo "Response: $delete_body"
    exit 1
fi

echo "City deleted successfully."

# Verify deletion
verify_delete_response=$(curl -s -w "%{http_code}" -o /tmp/verify_delete_response \
  --location --request GET "${BASE_URL}/dict/?type=city" \
  -H "Authorization: Bearer $access_token" \
  --header 'Content-Type: application/json')
http_code=$(echo $verify_delete_response | tail -n1)
verify_delete_body=$(cat /tmp/verify_delete_response)

deleted_city=$(echo "$verify_delete_body" | jq ".data[] | select(.id == $created_city_id)")

if [[ "$http_code" -ne 200 || -n "$deleted_city" ]]; then
    echo "City still exists after deletion. HTTP Code: $http_code"
    echo "Response: $verify_delete_body"
    exit 1
fi

echo "Deletion verified successfully. City no longer exists."