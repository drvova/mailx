#!/usr/bin/env bash
set -euo pipefail

RED='\033[0;31m'; GREEN='\033[0;32m'; YELLOW='\033[1;33m'; BOLD='\033[1m'; RESET='\033[0m'
info()  { echo -e "${GREEN}-->${RESET} $*"; }
fail()  { echo -e "${RED}ERROR:${RESET} $*"; exit 1; }

for cmd in go bun; do
  command -v "$cmd" &>/dev/null || fail "'$cmd' not found. Install it first."
done

# Create .env files if missing
[ -f api/.env ]   || { info "Creating api/.env";   cp api/.env.sample api/.env; }
[ -f app/.env ]   || { info "Creating app/.env";   cp app/.env.sample app/.env; }

DB_HOST=$(grep '^DB_HOSTS=' api/.env | cut -d= -f2)
DB_PORT=$(grep '^DB_PORT=' api/.env | cut -d= -f2)
DB_NAME=$(grep '^DB_NAME=' api/.env | cut -d= -f2)

# Install app dependencies
[ -d app/node_modules ] || { info "Installing app deps"; (cd app && bun install); }

# Start API
info "Starting API on :3000"
(cd api && go run ./cmd/) &
API_PID=$!

for i in $(seq 1 30); do
  curl -sf http://localhost:3000/healthcheck >/dev/null 2>&1 && break
  sleep 1
done
info "API ready"

# Start App
info "Starting App on :3001"
(cd app && bun run dev) &
APP_PID=$!

echo ""
echo -e "${BOLD}Running:${RESET}"
echo -e "  API  ${GREEN}http://localhost:3000${RESET}  (pid $API_PID)"
echo -e "  App  ${GREEN}http://localhost:3001${RESET}  (pid $APP_PID)"
echo -e "  DB   ${GREEN}MySQL${RESET} ${DB_HOST}:${DB_PORT}/${DB_NAME}"
echo ""
echo -e "Press ${YELLOW}Ctrl+C${RESET} to stop."
echo ""

cleanup() {
  info "Shutting down..."
  kill "$API_PID" "$APP_PID" 2>/dev/null
  wait "$API_PID" "$APP_PID" 2>/dev/null
}
trap cleanup EXIT INT TERM

wait
