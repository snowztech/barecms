FROM node:22-alpine3.22 AS frontend
WORKDIR /src/ui
COPY ui/package.json ui/package-lock.json ./
RUN npm ci --audit=false
COPY ui/ ./
RUN npm run build

FROM golang:1.24-alpine3.22 AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
COPY --from=frontend /src/ui/dist ./ui/dist
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/barecms ./cmd/main.go

FROM alpine:3.22
RUN apk add --no-cache ca-certificates tzdata \
    && addgroup -S -g 10001 barecms \
    && adduser -S -D -H -u 10001 -G barecms barecms \
    && mkdir -p /app/uploads \
    && chown -R barecms:barecms /app

WORKDIR /app
COPY --from=builder --chown=barecms:barecms /out/barecms ./barecms

USER barecms
EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=3s --start-period=10s --retries=3 \
  CMD wget -q -O /dev/null http://127.0.0.1:8080/readyz || exit 1

ENTRYPOINT ["./barecms"]
