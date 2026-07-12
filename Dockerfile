# Frontend build stage
FROM node:18-alpine AS frontend
WORKDIR /app
COPY ui/ .
RUN npm install
RUN npm run build

# Backend build stage
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
COPY .env.example .env
COPY --from=frontend /app/dist ./ui/dist
RUN go build -o /app/barecms ./cmd/main.go

# Final stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/barecms .

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=3s --start-period=10s --retries=3 \
  CMD wget -q -O /dev/null http://127.0.0.1:8080/readyz || exit 1

ENTRYPOINT [ "./barecms" ]
