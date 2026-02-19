# InstaTunnel

Instant, secure localhost tunneling with custom subdomains, analytics, and a modern web dashboard.

InstaTunnel is a fast ngrok alternative for sharing local apps, webhooks, and APIs over HTTPS in seconds.

## Why InstaTunnel

- Instant HTTPS tunnels for local apps and APIs
- Custom subdomains on *.instatunnel.my
- API key based auth for secure access
- Web dashboard for tunnel management and analytics
- Fast setup with a single CLI command
- Works well for webhook and MCP endpoint exposure

## Quick start (CLI)

Install globally:

~~~bash
npm install -g instatunnel
~~~

Sign up and get your API key:

- https://instatunnel.my/get-started

Run a tunnel:

~~~bash
instatunnel 3000
~~~

You will get a public URL like:

~~~text
https://your-subdomain.instatunnel.my
~~~

## MCP support

If your MCP server is running locally, you can expose it securely with InstaTunnel.

Example:

~~~bash
instatunnel 8787 --subdomain my-mcp
~~~

Then use the generated HTTPS URL in your MCP client/server configuration.

See full guide: docs/mcp.md

## Use cases

- Share local web apps with teammates
- Test webhooks (Stripe, GitHub, Slack, etc.)
- Expose local MCP endpoints for remote AI tooling
- Demo staging builds without deployment

## Docs

- Online docs: https://instatunnel.my/docs
- Quick start: docs/quickstart.md
- Configuration: docs/configuration.md
- MCP guide: docs/mcp.md
- Troubleshooting: docs/troubleshooting.md

## Project scope

This repository contains the InstaTunnel CLI and public documentation. The managed service, dashboard, and production infrastructure are hosted by InstaTunnel.

## License

MIT License. See LICENSE.

## SEO keywords

ngrok alternative, localhost tunnel, reverse proxy, webhook testing, secure tunnel, custom subdomains, mcp tunnel, model context protocol tunnel, local development tunneling
