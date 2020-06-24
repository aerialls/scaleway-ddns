package scaleway

import (
	domainAPI "github.com/scaleway/scaleway-sdk-go/api/domain/v2alpha2"
	"github.com/scaleway/scaleway-sdk-go/scw"
	log "github.com/sirupsen/logrus"
)

// DNS struct to contact Scaleway Domains
type DNS struct {
	logger *log.Logger
	client *scw.Client
}

// NewDNS returns a new DNS instance
func NewDNS(
	logger *log.Logger,
	organizationID string,
	accessKey string,
	secretKey string,
) (*DNS, error) {
	client, err := scw.NewClient(
		scw.WithDefaultOrganizationID(organizationID),
		scw.WithAuth(accessKey, secretKey),
	)

	if err != nil {
		return nil, err
	}

	return &DNS{
		logger: logger,
		client: client,
	}, nil
}

// UpdateRecord updates a DNS record from a specific domain
func (d *DNS) UpdateRecord(domain string, name string, ttl uint32, data string, recordType string) error {
	api := domainAPI.NewAPI(d.client)
	_, err := api.UpdateDNSZoneRecords(&domainAPI.UpdateDNSZoneRecordsRequest{
		DNSZone: domain,
		Changes: []*domainAPI.RecordChange{
			{
				Set: &domainAPI.RecordChangeSet{
					Name: name,
					Type: domainAPI.RecordTypeA,
					Records: []*domainAPI.Record{
						{
							Name: name,
							Data: data,
							Type: d.getRecordTypeFromString(recordType),
							TTL:  ttl,
						},
					},
				},
			},
		},
		ReturnAllRecords: scw.BoolPtr(false),
	})

	if err != nil {
		return err
	}

	return nil
}

// GetRecord returns the value from a DNS record
func (d *DNS) GetRecord(domain string, name string, recordType string) (string, error) {
	api := domainAPI.NewAPI(d.client)
	records, err := api.ListDNSZoneRecords(&domainAPI.ListDNSZoneRecordsRequest{
		DNSZone: domain,
		Name:    name,
		Type:    d.getRecordTypeFromString(recordType),
	})

	if err != nil {
		return "", err
	}

	if records.TotalCount != 1 {
		return "", nil
	}

	return records.Records[0].Data, nil
}

func (d *DNS) getRecordTypeFromString(recordType string) domainAPI.RecordType {
	if recordType == "A" {
		return domainAPI.RecordTypeA
	}

	if recordType == "AAAA" {
		return domainAPI.RecordTypeAAAA
	}

	return domainAPI.RecordTypeUnknown
}
