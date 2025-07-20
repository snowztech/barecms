# BareCMS Deployment Files

Production-ready deployment configuration for BareCMS.

## Files

- `docker-compose.yml` - Production Docker Compose setup
- `.env.template` - Environment variables template

## Quick Deploy

```bash
mkdir barecms-app && cd barecms-app

curl -sSL https://raw.githubusercontent.com/snowztech/barecms/main/deploy/.env.template | \
  sed "s/gen_jwt_secret/$(openssl rand -base64 32)/" > .env

curl -o docker-compose.yml \
  https://raw.githubusercontent.com/snowztech/barecms/main/deploy/docker-compose.yml

docker compose up -d
```

## Environment Variables

Edit `.env` to customize:
- `POSTGRES_PASSWORD` - Change from default
- `DATABASE_URL` - Update if using external database
- `PORT` - Change application port

## Management

```bash
# Update
docker compose pull && docker compose up -d

# Logs
docker compose logs -f barecms

# Backup database
docker compose exec postgres pg_dump -U barecms_user barecms_db > backup.sql
```

For full documentation, see the [main README](../README.md).
