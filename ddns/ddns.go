package ddns

import (
	"time"

	"github.com/aerialls/scaleway-ddns/config"
	"github.com/aerialls/scaleway-ddns/ip"
)

// DynamicDNSUpdater struct
type DynamicDNSUpdater struct {
	container *config.Container
	dryRun    bool
}

// NewDynamicDNSUpdater returns a new DynamicDNSUpdate
func NewDynamicDNSUpdater(
	container *config.Container,
	dryRun bool,
) *DynamicDNSUpdater {
	return &DynamicDNSUpdater{
		container: container,
		dryRun:    dryRun,
	}
}

// Start launches the ticker to update DNS records every interval
func (d *DynamicDNSUpdater) Start() {
	cfg := d.container.Config

	// Start the first update now
	d.doStart()

	ticker := time.NewTicker(time.Duration(cfg.Interval) * time.Second)
	for range ticker.C {
		d.doStart()
	}
}

func (d *DynamicDNSUpdater) doStart() {
	logger := d.container.Logger
	cfg := d.container.Config

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
		err := d.UpdateRecord(
			cfg.DomainConfig,
			recordCfg,
			recordType,
			d.dryRun,
		)

		if err != nil {
			logger.WithError(err).Errorf(
				"unable to update %s record", recordType,
			)
		}
	}
}

// UpdateRecord updates the DNS record on Scaleway DNS based
// on the current IPv4 or IPv6
func (d *DynamicDNSUpdater) UpdateRecord(
	domain config.DomainConfig,
	cfg config.IPConfig,
	recordType string,
	dryRun bool,
) error {
	logger := d.container.Logger
	dns := d.container.DNS

	if !cfg.Enabled {
		logger.Debugf("skipping %s update, disabled in the configuration", recordType)
		return nil
	}

	scalewayRecord, err := dns.GetRecord(domain.Name, domain.Record, recordType)
	if err != nil {
		return err
	}

	scalewayIP := "(empty)"
	if scalewayRecord != nil {
		scalewayIP = scalewayRecord.Data
	}

	currentIP, err := ip.GetPublicIP(cfg.URL)
	if err != nil {
		return err
	}

	logger.Debugf(
		"current IP state (scaleway=%s, current=%s)",
		scalewayIP,
		currentIP,
	)

	if scalewayIP == currentIP {
		logger.Debug("both IPs are identical, nothing to do")
		return nil
	}

	logger.Infof(
		"updating %s record from %s to %s",
		recordType,
		scalewayIP,
		currentIP,
	)

	if dryRun {
		logger.Info("running in dry-run mode, doing nothing")
		return nil
	}

	if scalewayRecord != nil {
		err = dns.UpdateRecord(
			domain.Name,
			scalewayRecord.ID,
			domain.Record,
			domain.TTL,
			currentIP,
			recordType,
		)
	} else {
		err = dns.AddRecord(
			domain.Name,
			domain.Record,
			domain.TTL,
			currentIP,
			recordType,
		)
	}

	if err != nil {
		return err
	}

	notifiers := d.container.Notifiers
	for _, notifier := range notifiers {
		subErr := notifier.Notify(
			domain.Name,
			domain.Record,
			recordType,
			scalewayIP,
			currentIP,
		)

		if subErr != nil {
			logger.WithError(subErr).Errorf(
				"unable to notify the IP change with notifier %T",
				notifier,
			)
		}
	}

	return nil
}
