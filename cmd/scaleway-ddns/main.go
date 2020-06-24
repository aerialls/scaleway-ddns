package main

import (
	"fmt"
	"os"
	"time"

	"github.com/aerialls/scaleway-ddns/config"
	"github.com/aerialls/scaleway-ddns/scaleway"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	cfgFile string
	verbose bool
	logger  *logrus.Logger
	dryRun  bool
)

var rootCmd = &cobra.Command{
	Use:   "scaleway-ddns",
	Short: "Dynamic DNS service based on Scaleway DNS",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("starting dynamic records for Scaleway DNS")

		cfg, err := config.NewConfig(cfgFile)
		if err != nil {
			logger.Fatal(err)
		}

		dns, err := scaleway.NewDNS(
			logger,
			cfg.ScalewayConfig.OrganizationID,
			cfg.ScalewayConfig.AccessKey,
			cfg.ScalewayConfig.SecretKey,
		)

		if err != nil {
			logger.Fatal(err)
		}

		ticker := time.NewTicker(time.Duration(cfg.Interval) * time.Second)

		for {
			select {
			case <-ticker.C:
				logger.Debugf(
					"updating A/AAAA records for %s.%s",
					cfg.DomainConfig.Record,
					cfg.DomainConfig.Name,
				)

				recordTypes := map[string]config.IPConfig{
					"A":    cfg.IPv4Config,
					"AAAA": cfg.IPv6Config,
				}

				for recordType, recordCfg := range recordTypes {
					err := UpdateDNSRecordFromCurrentIP(
						dns,
						cfg.DomainConfig,
						recordCfg,
						recordType,
						dryRun,
					)

					if err != nil {
						logger.WithError(err).Errorf(
							"unable to update %s record", recordType,
						)
					}
				}
			}
		}
	},
}

func init() {
	logger = logrus.New()

	cobra.OnInitialize(initConfig)

	rootCmd.Flags().StringVar(&cfgFile, "config", "", "config file")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose logging")
	rootCmd.Flags().BoolVarP(&dryRun, "dry-run", "d", false, "don't update DNS records")

	rootCmd.MarkFlagRequired("config")
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
