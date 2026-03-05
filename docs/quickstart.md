# Quick start

## 1) Install InstaTunnel CLI

```bash
npm install -g instatunnel
```

## 2) Create account or recover access

```bash
instatunnel auth login -e you@example.com
```

- New users receive verification + API key instructions.
- Existing users receive a recovery flow.

## 3) Save your API key

```bash
instatunnel auth set-key "it_your_api_key"
```

## 4) Start a localhost tunnel

```bash
instatunnel 3000 --subdomain myapp
```

Example output URL:

```text
https://myapp.instatunnel.my
```

## 5) Useful next commands

```bash
instatunnel --stats
instatunnel --logs
instatunnel --list
```
