package cmd

import (
	"os"

	"okp4/cosmos-faucet/internal/server"
	"okp4/cosmos-faucet/pkg/client"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const (
	FlagAddress       = "address"
	FlagMetrics       = "metrics"
	FlagHealth        = "health"
	FlagCaptchaSecret = "captcha-secret"
)

var serverConfig server.Config

// NewStartCommand returns a CLI command to start the REST api allowing to send tokens.
func NewStartCommand() *cobra.Command {
	var addr string

	startCmd := &cobra.Command{
		Use:   "start",
		Short: "Start the REST api",
		Run: func(cmd *cobra.Command, args []string) {
			faucet, err := client.NewFaucet(config)
			if err != nil {
				log.Fatal().Err(err).Msg("Failed create a new faucet instance")
			}

			defer func(faucet client.Faucet) {
				_ = faucet.Close()
				log.Info().Msg("Server stopped")
			}(faucet)

			serverConfig.Faucet = faucet
			server.NewServer(serverConfig).Start(addr)
		},
	}

	startCmd.Flags().StringVar(&addr, FlagAddress, ":8080", "rest api address")
	startCmd.Flags().BoolVar(&serverConfig.EnableMetrics, FlagMetrics, false, "enable metrics endpoint")
	startCmd.Flags().BoolVar(&serverConfig.EnableHealth, FlagHealth, false, "enable health endpoint")
	startCmd.Flags().StringVar(
		&serverConfig.CaptchaSecret,
		FlagCaptchaSecret,
		os.Getenv("CAPTCHA_SECRET"),
		"set Captcha secret (default from env: CAPTCHA_SECRET)",
	)

	return startCmd
}

func init() {
	rootCmd.AddCommand(NewStartCommand())
}
