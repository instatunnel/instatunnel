# Troubleshooting

## Invalid API key

- Ensure key starts with `it_`.
- Re-copy key without extra spaces.
- Re-set key:

```bash
instatunnel auth set-key "it_your_api_key"
```

## Existing email login confusion (409)

Use email flag form:

```bash
instatunnel auth login -e you@example.com
```

Existing accounts use recovery flow. New accounts use signup flow.

## Tunnel starts but URL shows local connection refused

Your local app is not running on the tunneled port.

Example: if command is `instatunnel 8787`, ensure local app is actually listening on `localhost:8787`.

## MCP says Pro/Business required

MCP mode requires paid plan support.

Check plan and retry:

```bash
instatunnel --stats
instatunnel 8787 --mcp --transport v2
```

## Need help

- Docs: https://instatunnel.my/docs/troubleshooting
- Support: support@instatunnel.my
