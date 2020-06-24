package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

const (
	IntervalMinValue = 60
)

// Config struct for the configuration file
type Config struct {
	Interval       uint           `yaml:"interval"`
	IPv4Config     IPConfig       `yaml:"ipv4"`
	IPv6Config     IPConfig       `yaml:"ipv6"`
	ScalewayConfig ScalewayConfig `yaml:"scaleway"`
	DomainConfig   DomainConfig   `yaml:"domain"`
}

// IPConfig struct for the required configuration for IPv4 or IPv6
type IPConfig struct {
	URL     string `yaml:"url"`
	Enabled bool   `yaml:"enabled"`
}

// ScalewayConfig struct for the required configuration to use the Scaleway API
type ScalewayConfig struct {
	OrganizationID string `yaml:"organization_id"`
	AccessKey      string `yaml:"access_key"`
	SecretKey      string `yaml:"secret_key"`
}

// DomainConfig struct for the domain parameters
type DomainConfig struct {
	Name   string `yaml:"name"`
	Record string `yaml:"record"`
	TTL    uint32 `yaml:"ttl"`
}

var (
	// DefaultIPv4Config is the default configuration for IPv4
	DefaultIPv4Config = IPConfig{
		URL:     "https://api-ipv4.ip.sb/ip",
		Enabled: true,
	}

	// DefaultIPv6Config is the default configuration for IPv6
	DefaultIPv6Config = IPConfig{
		URL:     "https://api-ipv6.ip.sb/ip",
		Enabled: false,
	}

	// DefaultScalewayConfig is the default configuration to use the Scaleway API
	DefaultScalewayConfig = ScalewayConfig{}

	// DefaultDomainConfig is the default domain configuration for common parameters
	DefaultDomainConfig = DomainConfig{
		Record: "ddns",
		TTL:    60,
	}

	// DefaultConfig is the global default configuration.
	DefaultConfig = Config{
		Interval:       300,
		DomainConfig:   DefaultDomainConfig,
		IPv4Config:     DefaultIPv4Config,
		IPv6Config:     DefaultIPv6Config,
		ScalewayConfig: DefaultScalewayConfig,
	}
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

	return nil
}
