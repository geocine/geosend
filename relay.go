// quickpodctl/relay.go
package main

import (
	"fmt"
	"strings"

	"github.com/schollz/croc/v9/src/tcp"
	"github.com/spf13/cobra"
)

var (
	relayHost  string
	relayPorts string
	relayPass  string
	relayDebug bool
)

var RelayCmd = &cobra.Command{
	Use:   "relay",
	Short: "Start your own croc relay",
	Run: func(cmd *cobra.Command, args []string) {
		ports := strings.Split(relayPorts, ",")
		if len(ports) < 1 {
			fmt.Println("You must specify at least one port (comma-separated)")
			return
		}
		debugLevel := "info"
		if relayDebug {
			debugLevel = "debug"
		}
		// Start all ports except the first in goroutines
		for i, port := range ports {
			if i == 0 {
				continue
			}
			go func(portStr string) {
				err := tcp.Run(debugLevel, relayHost, portStr, relayPass, strings.Join(ports[1:], ","))
				if err != nil {
					fmt.Printf("Error starting relay on port %s: %v\n", portStr, err)
				}
			}(port)
		}
		// Start the main relay on the first port
		err := tcp.Run(debugLevel, relayHost, ports[0], relayPass, strings.Join(ports[1:], ","))
		if err != nil {
			fmt.Printf("Error starting relay on port %s: %v\n", ports[0], err)
		}
	},
}

func init() {
	RelayCmd.Flags().StringVar(&relayHost, "host", "", "Host to bind the relay to (default: all interfaces)")
	RelayCmd.Flags().StringVar(&relayPorts, "ports", "9009,9010,9011,9012,9013", "Comma-separated list of ports for the relay")
	RelayCmd.Flags().StringVar(&relayPass, "pass", "pass123", "Password for the relay")
	RelayCmd.Flags().BoolVar(&relayDebug, "debug", false, "Enable debug logging for the relay")
}
