# BareCMS roadmap

## Status

What works today:

- CRUD: sites, collections, entries
- JWT authentication
- PostgreSQL with GORM
- Echo (Go) backend, React + TypeScript frontend
- Docker and docker-compose
- Public read API at `/api/:siteSlug/data`

## Phase 1. Bare basics

The basics every CMS needs. Without these, MCP positioning means nothing.

- [x] **Entry editing workflow** (`feature/content-updates`)
  - Tenant-safe update endpoint with schema validation
  - Editor UI reusing typed controls and the media picker

### Security and tenant isolation

- [x] **Tenant authorization implementation** (`security/tenant-authorization`)
  - Scope site listings to the authenticated user
  - Bind newly created sites to the JWT identity
  - Enforce ownership through site, collection, and entry relationships
- [x] Add cross-tenant authorization regression tests
- [x] Reject insecure default secrets in production (`security/config-hardening`)
- [x] Add authentication rate limiting, request size limits, and security headers

- [x] **Local file and image storage** (critical)
  - [x] Site-owned upload, serving, listing, and deletion endpoints (`feature/local-media-storage`)
  - [x] Files table tracking size, MIME type, original name, and storage name
  - [x] MIME allowlist, opaque filenames, size limits, and traversal protection
  - [x] Persistent Docker upload volume
  - [x] Image upload and existing-media picker in the collection editor (`feature/media-picker-ui`)
  - [x] `UPLOADS_DIR` and `MAX_FILE_SIZE` environment variables
- [ ] **Input validation and structured error responses**
  - [x] Required, schema, primitive type, URL, and date checks for entries (`feature/entry-validation`)
  - [x] Field-level `422 validation_failed` response shape
  - [ ] Length and numeric range constraints in collection schemas
  - Frontend renders field-level errors
- [ ] **API stability**
  - Lock response shapes for sites, collections, entries
  - [x] Bounded pagination and metadata for collection entries (`feature/api-pagination`)
  - [ ] Pagination on remaining authenticated list endpoints
- [ ] **Database integrity and migrations**
  - [x] Transactional site, collection, and account deletion (`reliability/database-integrity`)
  - [x] Cascade media metadata cleanup with best-effort disk cleanup
  - [x] Rollback regression coverage for failed destructive operations
  - [x] Versioned migration ledger, idempotency, and legacy upgrade tests (`reliability/versioned-migrations`)
  - [x] Collection slug uniqueness scoped to its site
  - [ ] Foreign-key constraints
- [ ] **UI polish**
  - Bug-free CRUD for every model
  - Loading, empty, and error states
  - Confirm dialogs on destructive actions
- [ ] **Docker docs**
  - Document every environment variable
  - Provide `.env.production` example
  - [x] Add liveness `/healthz`, database readiness `/readyz`, and container health check (`reliability/health-readiness`)
  - [x] Fail startup when database initialization fails
  - Slim final image with multi-stage build

## Phase 2. MCP-native pivot

The differentiator.

- [ ] **API tokens**
  - Per-token revocation
  - Scoped to a single site
  - `Authorization: Bearer <token>` accepted by the existing API and MCP
- [ ] **MCP server (HTTP transport)** exposing
  - `list_sites`, `list_collections`, `list_entries`
  - `get_entry`, `create_entry`, `update_entry`, `delete_entry`
  - `query_entries` with filters and pagination
  - `upload_file`
- [ ] **Schema introspection endpoint**
  - Lets agents discover collection fields without out-of-band setup
- [ ] **Connect to Claude Desktop docs**
  - Copy-paste config snippet
  - Same setup for Cursor
- [ ] **Submit to MCP registries** (Smithery, Anthropic directory)

## Phase 3. SaaS MVP

First paying customers.

- [ ] **Multi-site per user**. One user owns multiple projects.
- [ ] **Path-based multi-tenancy**. URLs look like `barecms.dev/u/<slug>`. No subdomain routing yet.
- [ ] **Stripe Checkout**. Single $15 plan. Manual upgrade for team customers.
- [ ] **Manual onboarding** for the first 30 customers. Instances created by hand.
- [ ] **Daily DB backup script**. Cron plus S3 upload, kept simple.

## Phase 4. Growth

Only when paid users ask.

- [ ] **External storage**. S3-compatible upload backend.
- [ ] **Multi-user team tier**. Invitations, Admin and Editor roles, gated to the $30/mo plan.
- [ ] **Subdomain routing and automated SSL**. Replaces path-based when it starts feeling limiting.
- [ ] **One starter template**. A polished Astro example, not a marketplace.
- [ ] **Automated provisioning**. Replaces manual onboarding once it becomes painful.

## Out of scope

The previous roadmap included these. They are removed.

- Password reset flow (manual recovery for first 100 users)
- Activity tracking and audit log
- Webhooks (replaced by MCP server)
- Granular permissions and custom roles
- Template marketplace with multiple starters
- Container orchestration (Docker Swarm, Kubernetes). Fly.io machines suffice.
- Editor AI buttons. External agents handle this over MCP.

## Pricing

| Tier | Price | Includes |
|---|---|---|
| Self-hosted | Free | Full features, your infra, no support |
| Solo (cloud) | $15/mo | 3 sites, local storage, single user, MCP access |
| Team (cloud) | $30/mo | Multi-user, S3 storage, more sites, MCP access |

## Success metrics

- **Phase 1**. 10+ stars. First self-hoster reports a clean install.
- **Phase 2**. Listed in MCP registries. First user reports editing via Claude.
- **Phase 3**. First paying customer at $15 MRR.
- **Phase 4**. $500+ MRR before adding team-tier features.

## Philosophy

Bare minimal, on purpose.

- One sentence describes the project: "MCP-native CMS."
- Deploy with a single binary or Docker image.
- Minimal surface, fast, no bloat.
- Users own their content.

Every feature decision passes through these.
