package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// NewConfig returns a new config object if the file exists
func NewConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to load the config from file %s: %w", path, err)
	}

	config := &Config{}
	*config = DefaultConfig

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	err = config.validate()
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (c *Config) validate() error {
	if c.Interval < IntervalMinValue {
		return fmt.Errorf(
			"update interval should not be below %ds (currently at %d)",
			IntervalMinValue,
			c.Interval,
		)
	}

	scwCfg := c.ScalewayConfig
	if scwCfg.AccessKey == "" || scwCfg.SecretKey == "" || scwCfg.ProjectID == "" {
		return fmt.Errorf(
			"scaleway parameters (access_key, secret_key, project_id) cannot be empty",
		)
	}

	tgCfg := c.TelegramConfig
	if tgCfg.Enabled && (tgCfg.Token == "" || tgCfg.ChatID == 0) {
		return fmt.Errorf(
			"token and chat_id are required for the Telegram notifier",
		)
	}

	return nil
}
