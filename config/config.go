package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// NewConfig returns a new config object if the file exists
func NewConfig(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to load the config from file (%s)", err)
	}

	cfg := &Config{}
	*cfg = DefaultConfig

	err = yaml.Unmarshal([]byte(data), &cfg)
	if err != nil {
		return nil, err
	}

	err = cfg.validate()
	if err != nil {
		return nil, err
	}

	return cfg, nil
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
	if scwCfg.AccessKey == "" || scwCfg.SecretKey == "" || scwCfg.OrganizationID == "" {
		return fmt.Errorf(
			"scaleway parameters (access_key, secret_key, organization_id) cannot be empty",
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
