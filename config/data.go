package config

const (
	// IntervalMinValue is the lowest possible value between two updates (in sec)
	IntervalMinValue = 60
)

// Config struct for the configuration file
type Config struct {
	Interval       uint           `yaml:"interval"`
	IPv4Config     IPConfig       `yaml:"ipv4"`
	IPv6Config     IPConfig       `yaml:"ipv6"`
	ScalewayConfig ScalewayConfig `yaml:"scaleway"`
	DomainConfig   DomainConfig   `yaml:"domain"`
	TelegramConfig TelegramConfig `yaml:"telegram"`
}

// IPConfig struct for the required configuration for IPv4 or IPv6
type IPConfig struct {
	URL     string `yaml:"url"`
	Enabled bool   `yaml:"enabled"`
}

// ScalewayConfig struct for the required configuration to use the Scaleway API
type ScalewayConfig struct {
	ProjectID string `yaml:"project_id"`
	AccessKey string `yaml:"access_key"`
	SecretKey string `yaml:"secret_key"`
}

// DomainConfig struct for the domain parameters
type DomainConfig struct {
	Name   string `yaml:"name"`
	Record string `yaml:"record"`
	TTL    uint32 `yaml:"ttl"`
}

// TelegramConfig struct for Telegram notifications after updates
type TelegramConfig struct {
	Enabled  bool   `yaml:"enabled"`
	Token    string `yaml:"token"`
	ChatID   int64  `yaml:"chat_id"`
	Template string `yaml:"template"`
}

var (
	// DefaultIPv4Config is the default configuration for IPv4
	DefaultIPv4Config = IPConfig{
		URL:     "https://api.ipify.org",
		Enabled: true,
	}

	// DefaultIPv6Config is the default configuration for IPv6
	DefaultIPv6Config = IPConfig{
		URL:     "https://api6.ipify.org",
		Enabled: false,
	}

	// DefaultScalewayConfig is the default configuration to use the Scaleway API
	DefaultScalewayConfig = ScalewayConfig{}

	// DefaultDomainConfig is the default domain configuration for common parameters
	DefaultDomainConfig = DomainConfig{
		Record: "ddns",
		TTL:    60,
	}

	// DefaultTelegramConfig is the default configuration to use Telegram notifications
	DefaultTelegramConfig = TelegramConfig{
		Enabled:  false,
		Template: "DNS record *{{ .RecordName }}.{{ .Domain }}* has been updated from *{{ .PreviousIP }}* to *{{ .NewIP }}*",
	}

	// DefaultConfig is the global default configuration.
	DefaultConfig = Config{
		Interval:       300,
		DomainConfig:   DefaultDomainConfig,
		IPv4Config:     DefaultIPv4Config,
		IPv6Config:     DefaultIPv6Config,
		ScalewayConfig: DefaultScalewayConfig,
		TelegramConfig: DefaultTelegramConfig,
	}
)
