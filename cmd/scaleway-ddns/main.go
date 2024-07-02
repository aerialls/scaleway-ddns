package main

import (
	"fmt"
	"os"

	ddnsconfig "github.com/aerialls/scaleway-ddns/config"
	"github.com/aerialls/scaleway-ddns/ddns"
	"github.com/aerialls/scaleway-ddns/notifier"
	"github.com/aerialls/scaleway-ddns/scaleway"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	configFile string
	verbose    bool
	logger     *logrus.Logger
	dryRun     bool
)

var rootCmd = &cobra.Command{
	Use:   "scaleway-ddns",
	Short: "Dynamic DNS service based on Scaleway DNS",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("starting dynamic records for Scaleway DNS")

		config, err := ddnsconfig.NewConfig(configFile)
		if err != nil {
			logger.Fatal(err)
		}

		dns, err := scaleway.NewDNS(
			logger,
			config.ScalewayConfig.ProjectID,
			config.ScalewayConfig.AccessKey,
			config.ScalewayConfig.SecretKey,
		)

		if err != nil {
			logger.Fatal(err)
		}

		container := ddnsconfig.NewContainer(logger, config, dns)

		if config.TelegramConfig.Enabled {
			telegramConfig := config.TelegramConfig
			container.AddNotifier(notifier.NewTelegram(
				telegramConfig.Token,
				telegramConfig.ChatID,
				telegramConfig.Template,
			))
		}

		updater := ddns.NewDynamicDNSUpdater(container, dryRun)
		updater.Start()
	},
}

func init() {
	logger = logrus.New()

	cobra.OnInitialize(initConfig)

	rootCmd.Flags().StringVar(&configFile, "config", "", "config file")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose logging")
	rootCmd.Flags().BoolVarP(&dryRun, "dry-run", "d", false, "don't update DNS records")

	err := rootCmd.MarkFlagRequired("config")
	if err != nil {
		logger.Fatal(err)
	}
}

func initConfig() {
	level := logrus.InfoLevel
	if verbose {
		level = logrus.DebugLevel
	}

	logger.SetLevel(level)
	logger.SetFormatter(&logrus.TextFormatter{})
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
