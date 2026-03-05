# MCP with InstaTunnel

Use InstaTunnel to expose a local MCP server over HTTPS.

## 1) Start your local MCP server

Example local endpoint:

```text
http://localhost:8787
```

## 2) Authenticate CLI

```bash
instatunnel auth login -e you@example.com
instatunnel auth set-key "it_your_api_key"
```

## 3) Start MCP tunnel

```bash
instatunnel 8787 --mcp --transport v2 --subdomain mymcp
```

## 4) Configure your MCP client

```json
{
  "url": "https://mymcp.instatunnel.my/mcp",
  "headers": {
    "Authorization": "Bearer YOUR_MCP_TOKEN"
  }
}
```

## v1 vs v2

- `v1`: default, best compatibility for normal HTTP/webhook flows.
- `v2`: use for streaming MCP behavior and long-lived responses.

## Security notes

- Treat MCP endpoints as sensitive.
- Use bearer token auth where possible.
- Rotate API keys/tokens if exposed.
