package client

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/instatunnel/cli/pkg/tunnel"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show active tunnels",
	Long:  "Display status of all active tunnels for your account",
	RunE:  runStatus,
}

func init() {
	rootCmd.AddCommand(statusCmd)
}

func runStatus(cmd *cobra.Command, args []string) error {
	client, err := tunnel.NewClient(getAPIKey(), getServerURL())
	if err != nil {
		return err
	}

	tunnels, err := client.GetTunnels()
	if err != nil {
		return err
	}

	if len(tunnels) == 0 {
		fmt.Println("No active tunnels")
		return nil
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tSUBDOMAIN\tPORT\tURL\tSTATUS")
	
	for _, tunnel := range tunnels {
		fmt.Fprintf(w, "%d\t%s\t%d\t%s\t%s\n",
			tunnel.ID, tunnel.Subdomain, tunnel.LocalPort, tunnel.URL, tunnel.Status)
	}
	
	return w.Flush()
}