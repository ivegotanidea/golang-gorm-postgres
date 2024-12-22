# Mock Data Setup
BASE_URL="http://localhost/api/v1"
PROFILES_URL="${BASE_URL}/profiles"

id="78b184ea-70a6-41b3-aea2-aec7b8c2e834"

TELEGRAM_USER_ID=${1:-6794234746}
PASSWORD=${2:-h5sh3d}

access_token=$(curl -s -X POST ${BASE_URL}/auth/bot/login \
                    -H "Content-Type: application/json" \
                    -d "{\"telegramUserId\": \"${TELEGRAM_USER_ID}\", \"password\": \"${PASSWORD}\"}" | jq -r .access_token)

# Example inactive photo IDs
inactive_photo_ids=("e43f1124-f547-4572-ab7f-31fa1f5139cd" "be68cf0a-9927-4591-a95c-e8cf7f6b9137")

set -x
# Generate the curl command with mock data
bulk_disable_resp_code=$(curl --connect-timeout 5 --max-time 10 -L -s -o /tmp/response -w "%{http_code}" \
                                 -X POST "${PROFILES_URL}/${id}/photos" \
                                 -H "Content-Type: application/json" \
                                 -b "access_token=${access_token}" \
                                 -d "$(jq -n --argjson photos "$(printf '%s\n' "${inactive_photo_ids[@]}" | jq -R . | jq -s '[.[] | {id: ., disabled: true}]')" '{photos: $photos}')")

# Debugging: Print response and status code
echo "HTTP Response Code: $bulk_disable_resp_code"
echo "Response Body: $(cat /tmp/response)"

