# InstaTunnel

Instant, secure localhost tunneling with custom subdomains, analytics, and a modern web dashboard.

InstaTunnel is a fast ngrok alternative for sharing local apps, webhooks, and prototypes over HTTPS in seconds.

## Why InstaTunnel

- **Instant HTTPS tunnels** for local apps and APIs
- **Custom subdomains** on *.instatunnel.my
- **API key based auth** for secure access
- **Web dashboard** for tunnel management and analytics
- **Fast setup** with a single CLI command

## Quick start (CLI)

`ash
npm install -g instatunnel
`

Sign up and get your API key:

- https://instatunnel.my/get-started

Then run a tunnel:

`ash
instatunnel 3000
`

You will see a public URL like:

`
https://your-subdomain.instatunnel.my
`

## Use cases

- Share local web apps with teammates
- Test webhooks (Stripe, GitHub, Slack, etc.)
- Demo a staging build without deploying
- Securely expose local APIs during development

## Docs

- Online docs: https://instatunnel.my/docs
- Quick start guide: docs/quickstart.md
- Configuration: docs/configuration.md
- Troubleshooting: docs/troubleshooting.md

## Project scope

This repository contains the **InstaTunnel CLI** and public documentation. The managed service, dashboard, and infrastructure are hosted by InstaTunnel.

## License

MIT License. See LICENSE.

## SEO keywords

ngrok alternative, localhost tunnel, reverse proxy, webhook testing, secure tunnel, custom subdomains, local development, tunneling CLI
