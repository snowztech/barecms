# Self-Hosting BareCMS

Deploy BareCMS on your own server using Docker. This guide provides two deployment options with step-by-step instructions.

## 📑 Table of Contents

- [📋 Prerequisites](#-prerequisites)
- [🚀 Quick Start](#-quick-start)
- [Common Setup Steps](#common-setup-steps)
- [Option 1: Docker Compose (All-in-One)](#option-1-docker-compose-all-in-one)
- [Option 2: Docker Only (External Database)](#option-2-docker-only-external-database)
- [🛠️ Management Commands](#️-management-commands)
- [🌐 Mapping Your Domain & Enabling HTTPS](#-mapping-your-domain--enabling-https)
- [🐛 Troubleshooting](#-troubleshooting)
- [📚 Resources & Support](#-resources--support)

---

## 📋 Prerequisites

Before you begin, ensure your server meets these requirements:

- **Operating System**: Linux server or VPS (Ubuntu 20.04+ recommended)
- **Docker**: [Install Docker](https://docs.docker.com/get-docker/)
- **Docker Compose**: [Install Docker Compose](https://docs.docker.com/compose/install/)
- **Git**: [Install Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
- **Basic Knowledge**: Familiarity with command line operations

---

## 🚀 Quick Start

### Common Setup Steps

These steps are required for both deployment options:

#### 1. Connect to Your Server

```bash
ssh your-user@your-server-ip
```

#### 2. Clone the Repository

```bash
mkdir -p ~/apps
cd ~/apps
git clone https://github.com/snowztech/barecms.git
cd barecms
```

#### 3. Configure Environment Variables

```bash
cp .env.example .env
nano .env
```

**Essential Environment Variables:**

```env
# Security Configuration (Required for both options)
JWT_SECRET=your-super-secret-jwt-key-here

# Application Configuration
PORT=8080

# Persistent uploads
UPLOADS_DIR=/app/uploads
MAX_FILE_SIZE=10485760
# Must be larger than MAX_FILE_SIZE to allow multipart overhead
MAX_REQUEST_BODY=12M
```

The Compose configuration mounts `/app/uploads` in a named `uploads_data`
volume. Back up this volume together with PostgreSQL; database-only backups do
not contain uploaded media.

💡 **Generate a secure JWT secret:**

```bash
openssl rand -base64 32
```

---

## Option 1: Docker Compose (All-in-One)

This option automatically sets up both the application and PostgreSQL database.

### Additional Environment Variables:

```env
# Database Configuration (Docker Compose)
POSTGRES_USER=barecms_user
POSTGRES_PASSWORD=your-secure-password
POSTGRES_DB=barecms_db

DATABASE_URL=postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@postgres:5432/${POSTGRES_DB}
```

### Deploy:

```bash
docker compose up -d
```

### Access:

Visit `http://your-server-ip:8080`

---

## Option 2: Docker Only (External Database)

Use this option if you have an existing PostgreSQL database.

### Additional Environment Variables:

```env
# Database Configuration (External Database)
DATABASE_URL=postgresql://username:password@your-db-host:5432/barecms_db
```

### Deploy:

```bash
# Pull the latest image
docker pull ghcr.io/snowztech/barecms:latest

# Run the container
docker run -d \
  --name barecms \
  -p 8080:8080 \
  --env-file .env \
  --restart unless-stopped \
  ghcr.io/snowztech/barecms:latest
```

### Access:

Visit `http://your-server-ip:8080`

---

## 🛠️ Management Commands

### Docker Compose Commands

```bash
# Start/Stop services
docker compose up -d
docker compose down

# View logs
docker compose logs -f barecms

# View database logs
docker compose logs -f postgres

# Update and restart
git pull
docker compose pull
docker compose up -d

# Backup database
docker compose exec postgres pg_dump -U barecms_user barecms_db > backup_$(date +%Y%m%d_%H%M%S).sql

# Restore database
docker compose exec -T postgres psql -U barecms_user -d barecms_db < backup_file.sql
```

### Docker Commands

```bash
# Start/Stop container
docker start barecms
docker stop barecms

# View logs
docker logs -f barecms

# Update application
docker pull ghcr.io/snowztech/barecms:latest
docker stop barecms
docker rm barecms
docker run -d \
  --name barecms \
  -p 8080:8080 \
  --env-file .env \
  --restart unless-stopped \
  ghcr.io/snowztech/barecms:latest
```

---

## 🌐 Mapping Your Domain & Enabling HTTPS

Make your app accessible via your own domain and secure it with HTTPS.

### 1. Point Your Domain

1. Log in to your domain registrar
2. Create an **A record** for `yourdomain.com` pointing to your server's IP
3. (Optional) Create a **CNAME record** for `www` pointing to `yourdomain.com`

### 2. Set Up HTTPS with a Reverse Proxy

#### Option A: Caddy (Recommended)

Caddy provides automatic HTTPS and simpler configuration.

**Install Caddy:**

```bash
sudo apt install -y debian-keyring debian-archive-keyring apt-transport-https
curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/gpg.key' | sudo gpg --dearmor -o /usr/share/keyrings/caddy-stable-archive-keyring.gpg
echo "deb [signed-by=/usr/share/keyrings/caddy-stable-archive-keyring.gpg] https://dl.cloudsmith.io/public/caddy/stable/deb/debian all main" | sudo tee /etc/apt/sources.list.d/caddy-stable.list
sudo apt update
sudo apt install caddy
```

**Configure Caddy:**

```bash
sudo nano /etc/caddy/Caddyfile
```

Add:

```
yourdomain.com {
    reverse_proxy localhost:8080
}
```

**Reload Caddy:**

```bash
sudo systemctl reload caddy
```

✅ **Caddy will:**

- Automatically request and install an SSL certificate via Let's Encrypt
- Renew it automatically
- Proxy requests to your app running on port 8080

#### Option B: Nginx + Certbot

If you prefer Nginx:

**Install Nginx and Certbot:**

```bash
sudo apt update
sudo apt install nginx certbot python3-certbot-nginx
```

**Configure Nginx:**

```bash
sudo nano /etc/nginx/sites-available/barecms
```

Add:

```nginx
server {
    listen 80;
    server_name yourdomain.com www.yourdomain.com;

    location / {
        proxy_pass http://localhost:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

**Enable the site:**

```bash
sudo ln -s /etc/nginx/sites-available/barecms /etc/nginx/sites-enabled/
sudo nginx -t
sudo systemctl reload nginx
```

**Get SSL certificate:**

```bash
sudo certbot --nginx -d yourdomain.com -d www.yourdomain.com
```

---

## 🐛 Troubleshooting

### Common Issues

#### Application Won't Start

```bash
# Check logs
docker compose logs barecms  # or docker logs barecms

# Verify environment variables
docker compose exec barecms env | grep DATABASE_URL
docker compose exec barecms env | grep JWT_SECRET
```

#### Database Connection Issues

```bash
# Test database connectivity (Docker Compose)
docker compose exec postgres psql -U barecms_user -d barecms_db -c "SELECT 1;"

# Check database status
docker compose ps postgres

# Check database logs
docker compose logs postgres
```

#### Port Already in Use

```bash
# Check what's using port 8080
sudo lsof -i :8080

# Change port in docker-compose.yml
# ports:
#   - "8081:8080"  # Use port 8081 instead
```

#### JWT Secret Issues

```bash
# Ensure JWT_SECRET is set and not empty
docker compose exec barecms env | grep JWT_SECRET

# Generate a new secret if needed
openssl rand -base64 32
```

#### Permission Issues

```bash
# Check container permissions
docker compose exec barecms ls -la /app

# Ensure proper ownership
sudo chown -R 1000:1000 ~/apps/barecms
```

### Getting Help

If you encounter issues:

1. **Check the logs**: Use `docker compose logs` or `docker logs`
2. **Verify configuration**: Ensure all environment variables are set correctly
3. **Database connectivity**: Test database connection manually
4. **Search existing issues** or create a new one on [GitHub](https://github.com/snowztech/barecms/issues)
5. **Join our community** for support

---

## 📚 Resources & Support

### Documentation

- [Docker Documentation](https://docs.docker.com/)
- [Docker Compose Documentation](https://docs.docker.com/compose/)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)
- [Caddy Documentation](https://caddyserver.com/docs/)
- [BareCMS API Documentation](API.md)

### Community & Support

- **GitHub**: [Repository](https://github.com/snowztech/barecms) | [Issues](https://github.com/snowztech/barecms/issues)
- **Discussions**: [GitHub Discussions](https://github.com/snowztech/barecms/discussions)

### Security Best Practices

- **Use strong passwords** for database and JWT secret
- **Keep your system updated**: `sudo apt update && sudo apt upgrade`
- **Use a firewall**: Configure UFW or iptables
- **Regular backups**: Set up automated database backups
- **Monitor logs**: Check application and system logs regularly

---

_Last updated: July 2025_
