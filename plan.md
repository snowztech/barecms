# BareCMS plan

A bare, MCP-native CMS.

> BareCMS is the minimal CMS your AI agents talk to. It does not add AI buttons to the editor. AI tools (Claude, Cursor, Vikusha) already do that better. BareCMS gives them a clean MCP interface to your content.

## Positioning

Most CMSes are racing to add AI features inside their editors. BareCMS goes the other direction.

- **Bare**. No editor bloat.
- **MCP-native**. Composes with the AI tools users already have.
- **Self-hostable**. Your content, your infra.
- **Solo-friendly pricing**. Built for freelancers and indie developers.

## Phases

1. Bare basics (~4-6 weekends)
2. MCP-native pivot (~1-2 weekends)
3. SaaS MVP (~3-4 weekends)
4. Growth, gated on paid demand

See `roadmap.md` for tasks.

## Out of scope

- Granular role permissions and custom roles
- Audit logs
- Webhooks (replaced by MCP server)
- Template marketplace (one good starter is enough)
- Editor AI buttons (handled by external agents over MCP)
- Page builder and drag-and-drop
- Plugin ecosystem

If a project needs those, Strapi or Sanity are better fits.

## Pricing

| Tier | Price | Includes |
|---|---|---|
| Self-hosted | Free | Full features, your infra, no support |
| Solo (cloud) | $15/mo | 3 sites, local storage, single user, MCP access |
| Team (cloud) | $30/mo | Multi-user, S3 storage, more sites, MCP access |

Solo is the entry tier. Team is the upgrade path for collaboration.

## Distribution

1. Ship the MCP server.
2. Announce on Twitter, Reddit (r/SelfHosted, r/programming), and Hacker News.
3. List in MCP registries (Smithery, Anthropic directory).
4. Publish a comparison post: "BareCMS vs Strapi for AI workflows."

## Philosophy

Bare minimal, on purpose.

- One sentence describes the project: "MCP-native CMS."
- Deploy with a single binary or Docker image.
- Minimal surface, fast, no bloat.
- Users own their content.

Every feature decision passes through these.
