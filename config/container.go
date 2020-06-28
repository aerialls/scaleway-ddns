package config

import (
	"github.com/aerialls/scaleway-ddns/scaleway"

	"github.com/sirupsen/logrus"
)

// Notifier interface to represent any notifier
type Notifier interface {
	Notify(domain string, recordName string, recordType string, previousIP string, newIP string) error
}

// Container structure to hold global objects
type Container struct {
	Logger    *logrus.Logger
	Config    *Config
	DNS       *scaleway.DNS
	Notifiers []Notifier
}

// NewContainer returns a new container instance
func NewContainer(
	logger *logrus.Logger,
	config *Config,
	dns *scaleway.DNS,

) *Container {
	return &Container{
		Config:    config,
		Logger:    logger,
		DNS:       dns,
		Notifiers: []Notifier{},
	}
}

// AddNotifier adds a new notifier into the container
func (c *Container) AddNotifier(notifier Notifier) {
	c.Logger.Debugf("New notifier %T added", notifier)
	c.Notifiers = append(c.Notifiers, notifier)
}
