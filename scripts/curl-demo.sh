#!/usr/bin/env bash
set -euo pipefail

# Скрипт становится зелёным после выполнения TODO и проверяет внешний contract.
BASE_URL="${BASE_URL:-http://localhost:8080}"
DEMO_TMP="$(mktemp -d)"
trap 'rm -rf "$DEMO_TMP"' EXIT
LAST_BODY="$DEMO_TMP/body"
LAST_HEADERS="$DEMO_TMP/headers"

request() {
  local title="$1"
  local expected="$2"
  shift 2
  local status
  status="$(curl -sS -D "$LAST_HEADERS" -o "$LAST_BODY" -w '%{http_code}' "$@")"
  if [[ "$status" != "$expected" ]]; then
    echo "FAIL: $title: status=$status, want=$expected" >&2
    cat "$LAST_BODY" >&2
    exit 1
  fi
  echo "OK: $title -> $status" >&2
}

register_and_login() {
  local email="$1"
  request "register $email" 201 -X POST "$BASE_URL/register" \
    -H 'Content-Type: application/json' \
    -d "{\"email\":\"$email\",\"password\":\"secret123\"}"
  request "login $email" 200 -X POST "$BASE_URL/login" \
    -H 'Content-Type: application/json' \
    -d "{\"email\":\"$email\",\"password\":\"secret123\"}"
  python3 -c 'import json,sys; print(json.load(sys.stdin)["token"])' <"$LAST_BODY"
}

suffix="${RANDOM}-$$"
request 'protected endpoint without token' 401 "$BASE_URL/notes"
alice_token="$(register_and_login "alice-${suffix}@example.com")"
bob_token="$(register_and_login "bob-${suffix}@example.com")"

request 'Alice creates Note' 201 -X POST "$BASE_URL/notes" \
  -H 'Content-Type: application/json' \
  -H "Authorization: Bearer $alice_token" \
  -d '{"title":"Private note","text":"Alice owns it"}'
note_id="$(python3 -c 'import json,sys; print(json.load(sys.stdin)["id"])' <"$LAST_BODY")"

request 'Bob cannot read Alice Note' 403 "$BASE_URL/notes/$note_id" \
  -H "Authorization: Bearer $bob_token"
request 'Alice deletes her Note' 204 -X DELETE "$BASE_URL/notes/$note_id" \
  -H "Authorization: Bearer $alice_token"
request 'deleted Note is missing' 404 "$BASE_URL/notes/$note_id" \
  -H "Authorization: Bearer $alice_token"

echo 'curl demo: PASS'
