# Configuration

## Authentication

Set key in local config:

```bash
instatunnel auth set-key "it_your_api_key"
```

Show key:

```bash
instatunnel auth show-key
instatunnel auth show-key --reveal --copy
```

## Core tunnel options

```bash
instatunnel 3000 --subdomain myapp
instatunnel 3000 --region us
instatunnel 3000 --password demo123
instatunnel 3000 --auth user:pass
```

## MCP mode

```bash
instatunnel 8787 --mcp --transport v2 --subdomain mymcp
```

- Use v1 when you want maximum compatibility.
- Use v2 for streaming MCP clients.

## Dashboard + docs

- Dashboard: https://dashboard.instatunnel.my
- Full CLI flags: https://instatunnel.my/docs/cli-flags
