# InstaTunnel CLI

[![npm version](https://img.shields.io/npm/v/instatunnel)](https://www.npmjs.com/package/instatunnel)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](./LICENSE)

InstaTunnel is a secure localhost tunnel service for developers who need a fast ngrok alternative for webhook testing, MCP endpoint tunneling, OAuth callback testing, and demo links.

Website: https://instatunnel.my

## Why developers use InstaTunnel

- HTTPS tunnel for localhost in one command
- Custom subdomains on `*.instatunnel.my`
- CLI + web dashboard workflow
- Built-in analytics and request logs
- MCP tunnel support for AI tooling workflows
- Better pricing for frequent webhook and API testing

## Install

```bash
npm install -g instatunnel
```

Verify:

```bash
instatunnel --version
```

Update:

```bash
npm install -g instatunnel@latest
```

## Quick start

1. Create or recover account access:

```bash
instatunnel auth login -e you@example.com
```

2. Save your API key locally:

```bash
instatunnel auth set-key "it_your_api_key"
```

3. Run a localhost tunnel:

```bash
instatunnel 3000 --subdomain myapp
```

4. Optional: confirm saved key:

```bash
instatunnel auth show-key --reveal --copy
```

## MCP tunnel quick start

For MCP streaming endpoints, use transport v2:

```bash
instatunnel 8787 --mcp --transport v2 --subdomain mymcp
```

Use the generated URL in your MCP client config:

```json
{
  "url": "https://mymcp.instatunnel.my/mcp",
  "headers": {
    "Authorization": "Bearer YOUR_MCP_TOKEN"
  }
}
```

## Common use cases

- Local webhook testing for Stripe, GitHub, Shopify, Twilio, PayPal
- OAuth callback development on localhost
- Sharing local QA/demo builds with teammates
- Exposing local MCP servers securely over HTTPS
- Temporary public API endpoint for integration tests

## InstaTunnel vs ngrok

If you are searching for terms like `ngrok alternative`, `localhost tunnel`, `secure tunnel`, `webhook testing tool`, or `MCP tunnel`, InstaTunnel is designed for those exact workflows with a CLI-first experience and lower-cost paid plans.

## Documentation

- Docs hub: https://instatunnel.my/docs
- CLI flags guide: https://instatunnel.my/docs/cli-flags
- Webhook testing guides: https://instatunnel.my/docs/webhooks
- MCP guide: https://instatunnel.my/docs#features-mcp-usage
- Troubleshooting: https://instatunnel.my/docs/troubleshooting
- Pricing: https://instatunnel.my/pricing
- Changelog: ./CHANGELOG.md

Local docs in this repo:

- docs/quickstart.md
- docs/configuration.md
- docs/webhooks.md
- docs/mcp.md
- docs/troubleshooting.md

## Support

- Support email: support@instatunnel.my
- GitHub issues: https://github.com/instatunnel/instatunnel/issues

## Public repo scope

This public repository is focused on the InstaTunnel CLI docs and user-facing setup material. Managed service infrastructure and internal platform services run in private repositories.

## License

MIT. See ./LICENSE
