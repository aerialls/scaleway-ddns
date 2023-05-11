package scaleway

import (
	domainAPI "github.com/scaleway/scaleway-sdk-go/api/domain/v2beta1"
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
	projectID string,
	accessKey string,
	secretKey string,
) (*DNS, error) {
	client, err := scw.NewClient(
		scw.WithDefaultProjectID(projectID),
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

// AddRecord adds a new DNS record
func (d *DNS) AddRecord(domain string, name string, ttl uint32, data string, recordType string) error {
	api := domainAPI.NewAPI(d.client)
	_, err := api.UpdateDNSZoneRecords(&domainAPI.UpdateDNSZoneRecordsRequest{
		DNSZone: domain,
		Changes: []*domainAPI.RecordChange{
			{
				Add: &domainAPI.RecordChangeAdd{
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

// UpdateRecord updates an existing DNS record
func (d *DNS) UpdateRecord(domain string, id string, name string, ttl uint32, data string, recordType string) error {
	api := domainAPI.NewAPI(d.client)
	_, err := api.UpdateDNSZoneRecords(&domainAPI.UpdateDNSZoneRecordsRequest{
		DNSZone: domain,
		Changes: []*domainAPI.RecordChange{
			{
				Set: &domainAPI.RecordChangeSet{
					ID: &id,
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

// GetRecord returns a DNS record
func (d *DNS) GetRecord(domain string, name string, recordType string) (*domainAPI.Record, error) {
	api := domainAPI.NewAPI(d.client)
	records, err := api.ListDNSZoneRecords(&domainAPI.ListDNSZoneRecordsRequest{
		DNSZone: domain,
		Name:    name,
		Type:    d.getRecordTypeFromString(recordType),
	})

	if err != nil {
		return nil, err
	}

	if records.TotalCount == 0 {
		return nil, nil
	}

	return records.Records[0], nil
}

func (d *DNS) getRecordTypeFromString(recordType string) domainAPI.RecordType {
	switch recordType {
	case "A":
		return domainAPI.RecordTypeA
	case "AAAA":
		return domainAPI.RecordTypeAAAA
	default:
		return domainAPI.RecordTypeUnknown
	}
}
