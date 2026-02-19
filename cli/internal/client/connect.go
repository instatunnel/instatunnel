package client

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/instatunnel/cli/pkg/tunnel"
)

var connectCmd = &cobra.Command{
	Use:   "connect [port]",
	Short: "Connect to an existing tunnel created via web dashboard",
	Long: `Connect to an existing tunnel that was created via the web dashboard.
	
You can specify either a subdomain or a custom domain using the respective flags.

Examples:
  instatunnel connect 3000 --subdomain myapp
  instatunnel connect 3000 --domain myapp.com
  instatunnel connect 8000 -s myproject`,
	Args: cobra.ExactArgs(1),
	RunE: runConnect,
}

func init() {
	connectCmd.Flags().StringVarP(&subdomain, "subdomain", "s", "", "Subdomain of existing tunnel")
	connectCmd.Flags().StringVarP(&customDomain, "domain", "d", "", "Custom domain of existing tunnel")
	
	rootCmd.AddCommand(connectCmd)
}

func runConnect(cmd *cobra.Command, args []string) error {
	// Validate that either subdomain or domain is provided
	if subdomain == "" && customDomain == "" {
		return fmt.Errorf("Either --subdomain or --domain must be specified")
	}
	
	if subdomain != "" && customDomain != "" {
		return fmt.Errorf("Cannot specify both --subdomain and --domain at the same time")
	}

	port, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("Invalid port number: %s", args[0])
	}

	var tunnelURL string
	
	if customDomain != "" {
		tunnelURL = customDomain
		fmt.Printf("🔗 Connecting to existing custom domain tunnel: %s\n", customDomain)
	} else {
		tunnelURL = fmt.Sprintf("%s.instatunnel.my", subdomain)
		fmt.Printf("🔗 Connecting to existing tunnel: %s\n", tunnelURL)
	}
	
	fmt.Printf("🚀 Forwarding to localhost:%d\n", port)
	
	apiKey := getAPIKey()
	if apiKey == "" {
		fmt.Printf("💡 Running in anonymous mode (no API key found)\n")
	}
	
	client, err := tunnel.NewClient(apiKey, getServerURL())
	if err != nil {
		return fmt.Errorf("❌ Failed to connect: %v", err)
	}
	
	// Connect to existing tunnel
	if customDomain != "" {
		return client.ConnectToExistingCustomDomainTunnel(customDomain, port)
	} else {
		return client.ConnectToExistingTunnel(subdomain, port)
	}
}
