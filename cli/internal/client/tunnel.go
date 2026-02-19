package client

import (
	"strconv"

	"github.com/spf13/cobra"
)


var tunnelCmd = &cobra.Command{
	Use:   "tunnel [port]",
	Short: "Create a tunnel to expose local port",
	Long:  "Create a tunnel to expose your local port to the internet",
	Args:  cobra.ExactArgs(1),
	RunE:  runTunnel,
}


func init() {
	tunnelCmd.Flags().StringVarP(&subdomain, "subdomain", "s", "", "Custom subdomain (default: random)")
	tunnelCmd.Flags().StringVarP(&region, "region", "r", "us", "Server region")

	rootCmd.AddCommand(tunnelCmd)
}

func runTunnel(cmd *cobra.Command, args []string) error {
	port, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	return startTunnel(port)
}

