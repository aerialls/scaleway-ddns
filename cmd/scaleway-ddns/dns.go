package main

import (
	"github.com/aerialls/scaleway-ddns/config"
	"github.com/aerialls/scaleway-ddns/ip"
	"github.com/aerialls/scaleway-ddns/scaleway"
)

// UpdateDNSRecordFromCurrentIP updates the DNS record on Scaleway DNS based
// on the current IP for IPv4 or IPv6
func UpdateDNSRecordFromCurrentIP(
	dns *scaleway.DNS,
	domain config.DomainConfig,
	cfg config.IPConfig,
	recordType string,
	dryRun bool,
) error {
	if !cfg.Enabled {
		logger.Debugf("skipping %s update, disabled in the configuration", recordType)
		return nil
	}

	scalewayIP, err := dns.GetRecord(domain.Name, domain.Record, recordType)
	if err != nil {
		return err
	}

	currentIP, err := ip.GetPublicIP(cfg.URL)
	if err != nil {
		return err
	}

	logger.Debugf("current IP registered in Scaleway is %s", scalewayIP)
	logger.Debugf("current public IP is %s", currentIP)

	if scalewayIP == currentIP {
		logger.Debug("both IPs are identical, nothing to do")
		return nil
	}

	logger.Infof(
		"updating %s record from %s to %s", recordType, scalewayIP, currentIP,
	)

	if dryRun {
		logger.Info("running in dry-run mode, doing nothing")
		return nil
	}

	err = dns.UpdateRecord(
		domain.Name,
		domain.Record,
		domain.TTL,
		currentIP,
		recordType,
	)

	if err != nil {
		return err
	}

	return nil
}
