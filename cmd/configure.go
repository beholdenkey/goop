package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure AWS SSO",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter AWS SSO URL: ")
		ssoURL, _ := reader.ReadString('\n')
		ssoURL = strings.TrimSpace(ssoURL)

		// Save the SSO URL to the configuration file
		viper.Set("aws.sso_url", ssoURL)

		// Write configuration to file
		if err := viper.WriteConfigAs(viper.ConfigFileUsed()); err != nil {
			if os.IsNotExist(err) {
				if err := viper.SafeWriteConfigAs(os.Getenv("HOME") + "/.goop.yaml"); err != nil {
					fmt.Println("Error saving configuration:", err)
					return
				}
				fmt.Println("Configuration file created and saved.")
			} else {
				fmt.Println("Error writing configuration:", err)
			}
		} else {
			fmt.Println("AWS SSO URL saved successfully.")
		}
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
}
