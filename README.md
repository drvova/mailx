# Email Service

## API

- Go
- Fiber (API, middleware)
- Gorm (ORM)
- SQLite (Database)
- Swagger (API Documentation)

## App

- TypeScript
- Vue.js
- Vite (Bundler)
- Tailwind (Styling)

## Browser Extension

- [WXT](https://wxt.dev)

## Installation

### Prerequisites

- Go 1.24+
- Bun (or Node.js 18+)

### Quick Start

```bash
./start.sh
```

This creates `.env` files from samples, installs npm deps, and starts both services.

### Manual Setup

```bash
# Config
cp api/.env.sample api/.env
cp app/.env.sample app/.env

# API (creates mailx.db automatically)
cd api && go run ./cmd/

# App (in another terminal)
cd app && npm install && npm run dev
```

### Endpoints

- **App**: http://localhost:3001
- **API**: http://localhost:3000
- **DB**: `api/mailx.db` (SQLite, auto-created on first run)

### Config

Edit `api/.env`:

```ini
DB_PATH=mailx.db
API_PORT=3000
API_ALLOW_ORIGIN=http://localhost:3001
DOMAINS=example1.net,example2.com
SMTP_CLIENT_HOST=smtp.example.net
SMTP_CLIENT_PORT=2525
```

Edit `app/.env`:

```ini
VITE_API_URL=http://localhost:3000
VITE_DOMAINS=example1.net,example2.net
```

> [!TIP]
> For local SMTP testing, use [MailHog](https://github.com/mailhog/MailHog) or [MailTrap](https://mailtrap.io/email-sandbox/).

## Test

```bash
cd api
go test ./... -v
go vet ./...
```

## API Documentation

http://localhost:3000/docs

```bash
cd api
swag init -g cmd/main.go
```
