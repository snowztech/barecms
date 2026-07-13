# BareCMS Deployment Files

Production-ready deployment configuration for BareCMS.

## Files

- `docker-compose.yml` - Production Docker Compose setup
- `.env.template` - Environment variables template

## Quick Deploy

```bash
mkdir barecms-app && cd barecms-app

curl -sSLo .env https://raw.githubusercontent.com/snowztech/barecms/main/deploy/.env.template

curl -o docker-compose.yml \
  https://raw.githubusercontent.com/snowztech/barecms/main/deploy/docker-compose.yml

JWT_SECRET=$(openssl rand -hex 32)
DB_PASSWORD=$(openssl rand -hex 24)
sed -i "s/change_jwt_secret/$JWT_SECRET/" .env
sed -i "s/change_database_password/$DB_PASSWORD/g" .env

docker compose up -d
docker compose ps
curl --fail http://localhost:8080/readyz
```

## Environment Variables

| Variable | Purpose |
|---|---|
| `JWT_SECRET` | Token-signing secret; use at least 32 random characters |
| `POSTGRES_USER` | PostgreSQL user created by the database container |
| `POSTGRES_PASSWORD` | PostgreSQL password; must match `DATABASE_URL` |
| `POSTGRES_DB` | PostgreSQL database name |
| `DATABASE_URL` | Complete PostgreSQL connection string |
| `PORT` | Host port published for BareCMS |
| `ENV` | Use `production` to enable production validation and HSTS |
| `DEBUG` | Application debug flag; keep `false` in production |
| `AUTH_RATE_LIMIT_PER_MINUTE` | Per-IP login and registration burst/refill limit |
| `MAX_FILE_SIZE` | Maximum individual upload size in bytes |
| `MAX_REQUEST_BODY` | Global body limit; must exceed `MAX_FILE_SIZE` for multipart overhead |
| `UPLOADS_DIR` | Media path inside the container; keep `/app/uploads` with this Compose file |

## Management

```bash
# Update
docker compose pull && docker compose up -d

# Logs
docker compose logs -f barecms

# Backup database
docker compose exec postgres pg_dump -U barecms_user barecms_db > backup.sql

# Backup uploaded media
docker compose exec -T barecms tar -czf - -C /app uploads > uploads.tar.gz

# Restore database into an empty database
docker compose exec -T postgres psql -U barecms_user barecms_db < backup.sql

# Restore uploaded media
docker compose exec -T barecms tar -xzf - -C /app < uploads.tar.gz
```

Back up both PostgreSQL and uploads. Verify backups periodically by restoring
them into a separate test deployment and checking `/readyz` plus public media.

For full documentation, see the [main README](../README.md).
