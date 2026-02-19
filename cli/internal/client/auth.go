package client

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/instatunnel/cli/pkg/tunnel"
)

var (
	email string
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authentication commands",
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Register or login with email",
	RunE:  runLogin,
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Remove stored credentials",
	RunE:  runLogout,
}

func init() {
	loginCmd.Flags().StringVarP(&email, "email", "e", "", "Email address")
	loginCmd.MarkFlagRequired("email")

	authCmd.AddCommand(loginCmd)
	authCmd.AddCommand(logoutCmd)
	rootCmd.AddCommand(authCmd)
}

func runLogin(cmd *cobra.Command, args []string) error {
	client, err := tunnel.NewClient("", getServerURL())
	if err != nil {
		return err
	}

	apiKey, err := client.Register(email)
	if err != nil {
		return fmt.Errorf("registration failed: %w", err)
	}

	// Save API key to config
	viper.Set("api_key", apiKey)
	viper.Set("email", email)

	configFile := viper.ConfigFileUsed()
	if configFile == "" {
		home, _ := os.UserHomeDir()
		configFile = home + "/.instatunnel.yaml"
	}

	err = viper.WriteConfigAs(configFile)
	if err != nil {
		return fmt.Errorf("failed to save config: %w", err)
	}

	fmt.Printf("✓ Successfully authenticated as %s\n", email)
	fmt.Printf("✓ API key saved to %s\n", configFile)
	fmt.Println("\nYou can now start tunneling with:")
	fmt.Println("  instatunnel 3000")

	return nil
}

func runLogout(cmd *cobra.Command, args []string) error {
	viper.Set("api_key", "")
	viper.Set("email", "")

	configFile := viper.ConfigFileUsed()
	if configFile == "" {
		home, _ := os.UserHomeDir()
		configFile = home + "/.instatunnel.yaml"
	}

	err := viper.WriteConfigAs(configFile)
	if err != nil {
		return fmt.Errorf("failed to save config: %w", err)
	}

	fmt.Println("✓ Successfully logged out")
	return nil
}