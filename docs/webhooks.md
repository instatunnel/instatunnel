# Webhook testing with InstaTunnel

InstaTunnel is commonly used as a localhost webhook testing tool.

## Generic flow

1. Start local app on a port (example 3000).
2. Run tunnel:

```bash
instatunnel 3000 --subdomain webhook-dev
```

3. Set provider webhook URL to:

```text
https://webhook-dev.instatunnel.my/your/webhook/path
```

4. Send provider test event.
5. Verify signature and idempotency handling in your app.

## CLI webhook helpers

```bash
instatunnel webhook init --provider stripe --port 4242 --path /webhook
instatunnel webhook verify --provider github --secret-env GITHUB_WEBHOOK_SECRET
instatunnel webhook test --provider stripe
```

## Provider guides

- Stripe: https://instatunnel.my/docs/webhooks/stripe
- GitHub: https://instatunnel.my/docs/webhooks/github
- Shopify: https://instatunnel.my/docs/webhooks/shopify
- Twilio: https://instatunnel.my/docs/webhooks/twilio
- PayPal: https://instatunnel.my/docs/webhooks/paypal
