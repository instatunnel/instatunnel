# MCP with InstaTunnel

Use InstaTunnel to expose a local MCP endpoint over HTTPS for remote access.

## 1. Run your local MCP server

Example local endpoint (replace with your setup):

~~~text
http://localhost:8787
~~~

## 2. Expose the port with InstaTunnel

~~~bash
instatunnel 8787 --subdomain my-mcp
~~~

This returns a public URL like:

~~~text
https://my-mcp.instatunnel.my
~~~

## 3. Use the public URL in your MCP config

Point your MCP client/integration to the generated HTTPS URL and path expected by your MCP server.

## Security notes

- Treat your MCP endpoint as sensitive.
- Prefer authenticated endpoints.
- Rotate API keys if exposed.
