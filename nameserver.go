package goinwx

import (
	"errors"
	"fmt"
	"time"

	"github.com/fatih/structs"
	"github.com/go-viper/mapstructure/v2"
)

const (
	methodNameserverCheck        = "nameserver.check"
	methodNameserverCreate       = "nameserver.create"
	methodNameserverCreateRecord = "nameserver.createRecord"
	methodNameserverDelete       = "nameserver.delete"
	methodNameserverDeleteRecord = "nameserver.deleteRecord"
	methodNameserverInfo         = "nameserver.info"
	methodNameserverList         = "nameserver.list"
	methodNameserverUpdate       = "nameserver.update"
	methodNameserverUpdateRecord = "nameserver.updateRecord"
)

// NameserverService API access to Nameservers.
type NameserverService service

// Check checks a domain on nameservers.
func (s *NameserverService) Check(domain string, nameservers []string) (*NameserverCheckResponse, error) {
	req := s.client.NewRequest(methodNameserverCheck, map[string]interface{}{
		"domain": domain,
		"ns":     nameservers,
	})

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result NameserverCheckResponse

	err = mapstructure.Decode(resp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Info gets information.
func (s *NameserverService) Info(request *NameserverInfoRequest) (*NameserverInfoResponse, error) {
	req := s.client.NewRequest(methodNameserverInfo, structs.Map(request))

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	result := NameserverInfoResponse{}

	err = mapstructure.Decode(resp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// List lists nameservers for a domain.
// Deprecated: use ListWithParams instead.
func (s *NameserverService) List(domain string) (*NameserverListResponse, error) {
	requestMap := map[string]interface{}{
		"domain": "*",
		"wide":   2,
	}

	if domain != "" {
		requestMap["domain"] = domain
	}

	req := s.client.NewRequest(methodNameserverList, requestMap)

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	result := NameserverListResponse{}

	err = mapstructure.Decode(resp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// ListWithParams lists nameservers for a domain.
func (s *NameserverService) ListWithParams(request *NameserverListRequest) (*NameserverListResponse, error) {
	if request == nil {
		return nil, errors.New("request can't be nil")
	}

	requestMap := structs.Map(request)
	requestMap["wide"] = "2"

	req := s.client.NewRequest(methodNameserverList, requestMap)

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	result := NameserverListResponse{}

	err = mapstructure.Decode(resp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Create creates a nameserver.
func (s *NameserverService) Create(request *NameserverCreateRequest) (int, error) {
	req := s.client.NewRequest(methodNameserverCreate, structs.Map(request))

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string]int

	err = mapstructure.Decode(resp, &result)
	if err != nil {
		return 0, err
	}

	return result["roId"], nil
}

// CreateRecord creates a DNS record.
func (s *NameserverService) CreateRecord(request *NameserverRecordRequest) (int, error) {
	req := s.client.NewRequest(methodNameserverCreateRecord, structs.Map(request))

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string]int

	err = mapstructure.Decode(resp, &result)
	if err != nil {
		return 0, err
	}

	return result["id"], nil
}

// UpdateRecord updates a DNS record.
func (s *NameserverService) UpdateRecord(recID int, request *NameserverRecordRequest) error {
	if request == nil {
		return errors.New("request can't be nil")
	}

	requestMap := structs.Map(request)
	requestMap["id"] = recID

	req := s.client.NewRequest(methodNameserverUpdateRecord, requestMap)

	_, err := s.client.Do(req)
	if err != nil {
		return err
	}

	return nil
}

// DeleteRecord deletes a DNS record.
func (s *NameserverService) DeleteRecord(recID int) error {
	req := s.client.NewRequest(methodNameserverDeleteRecord, map[string]interface{}{
		"id": recID,
	})

	_, err := s.client.Do(req)
	if err != nil {
		return err
	}

	return nil
}

// FindRecordByID search a DNS record by ID.
func (s *NameserverService) FindRecordByID(recID int) (*NameserverRecord, *NameserverDomain, error) {
	listResp, err := s.client.Nameservers.ListWithParams(&NameserverListRequest{})
	if err != nil {
		return nil, nil, err
	}

	for _, domainItem := range listResp.Domains {
		resp, err := s.client.Nameservers.Info(&NameserverInfoRequest{RoID: domainItem.RoID})
		if err != nil {
			return nil, nil, err
		}

		for _, record := range resp.Records {
			if record.ID == recID {
				return &record, &domainItem, nil
			}
		}
	}

	return nil, nil, fmt.Errorf("couldn't find INWX Record for id %d", recID)
}

// NameserverCheckResponse API model.
type NameserverCheckResponse struct {
	Details []string `mapstructure:"details"`
	Status  string   `mapstructure:"status"`
}

// NameserverRecordRequest API model.
type NameserverRecordRequest struct {
	RoID                   int    `structs:"roId,omitempty"`
	Domain                 string `structs:"domain,omitempty"`
	Type                   string `structs:"type"`
	Content                string `structs:"content"`
	Name                   string `structs:"name,omitempty"`
	TTL                    int    `structs:"ttl,omitempty"`
	Priority               int    `structs:"prio,omitempty"`
	URLAppend              bool   `structs:"urlAppend,omitempty"`
	URLRedirectType        string `structs:"urlRedirectType,omitempty"`
	URLRedirectTitle       string `structs:"urlRedirectTitle,omitempty"`
	URLRedirectDescription string `structs:"urlRedirectDescription,omitempty"`
	URLRedirectFavIcon     string `structs:"urlRedirectFavIcon,omitempty"`
	URLRedirectKeywords    string `structs:"urlRedirectKeywords,omitempty"`
}

// NameserverCreateRequest API model.
type NameserverCreateRequest struct {
	Domain                 string   `structs:"domain"`
	Type                   string   `structs:"type"`
	Nameservers            []string `structs:"ns,omitempty"`
	MasterIP               string   `structs:"masterIp,omitempty"`
	Web                    string   `structs:"web,omitempty"`
	Mail                   string   `structs:"mail,omitempty"`
	SoaEmail               string   `structs:"soaEmail,omitempty"`
	URLRedirectType        string   `structs:"urlRedirectType,omitempty"`
	URLRedirectTitle       string   `structs:"urlRedirectTitle,omitempty"`
	URLRedirectDescription string   `structs:"urlRedirectDescription,omitempty"`
	URLRedirectFavIcon     string   `structs:"urlRedirectFavIcon,omitempty"`
	URLRedirectKeywords    string   `structs:"urlRedirectKeywords,omitempty"`
	Testing                bool     `structs:"testing,omitempty"`
}

// NameserverInfoRequest API model.
type NameserverInfoRequest struct {
	Domain   string `structs:"domain,omitempty"`
	RoID     int    `structs:"roId,omitempty"`
	RecordID int    `structs:"recordId,omitempty"`
	Type     string `structs:"type,omitempty"`
	Name     string `structs:"name,omitempty"`
	Content  string `structs:"content,omitempty"`
	TTL      int    `structs:"ttl,omitempty"`
	Priority int    `structs:"prio,omitempty"`
}

// NamserverInfoResponse API model.
// Deprecated: Use NameserverInfoResponse instead.
type NamserverInfoResponse = NameserverInfoResponse

// NameserverInfoResponse API model.
type NameserverInfoResponse struct {
	RoID          int                `mapstructure:"roId"`
	Domain        string             `mapstructure:"domain"`
	Type          string             `mapstructure:"type"`
	MasterIP      string             `mapstructure:"masterIp"`
	LastZoneCheck time.Time          `mapstructure:"lastZoneCheck"`
	SlaveDNS      []SlaveInfo        `mapstructure:"slaveDns"`
	SOASerial     string             `mapstructure:"SOAserial"`
	Count         int                `mapstructure:"count"`
	Records       []NameserverRecord `mapstructure:"record"`
}

// SlaveInfo API model.
type SlaveInfo struct {
	Name string `mapstructure:"name"`
	IP   string `mapstructure:"ip"`
}

// NameserverRecord API model.
type NameserverRecord struct {
	ID                     int    `mapstructure:"id"`
	Name                   string `mapstructure:"name"`
	Type                   string `mapstructure:"type"`
	Content                string `mapstructure:"content"`
	TTL                    int    `mapstructure:"TTL"`
	Priority               int    `mapstructure:"prio"`
	URLAppend              bool   `mapstructure:"urlAppend,omitempty"`
	URLRedirectType        string `mapstructure:"urlRedirectType"`
	URLRedirectTitle       string `mapstructure:"urlRedirectTitle"`
	URLRedirectDescription string `mapstructure:"urlRedirectDescription"`
	URLRedirectKeywords    string `mapstructure:"urlRedirectKeywords"`
	URLRedirectFavIcon     string `mapstructure:"urlRedirectFavIcon"`
}

// NameserverListRequest API model.
type NameserverListRequest struct {
	Domain    string `structs:"domain,omitempty"`
	Wide      int    `structs:"wide,omitempty"`
	Page      int    `structs:"page,omitempty"`
	PageLimit int    `structs:"pagelimit,omitempty"`
}

// NamserverListResponse API model.
// Deprecated: Use NameserverListResponse instead.
type NamserverListResponse = NameserverListResponse

// NameserverListResponse API model.
type NameserverListResponse struct {
	Count   int                `mapstructure:"count"`
	Domains []NameserverDomain `mapstructure:"domains"`
}

// NameserverDomain API model.
type NameserverDomain struct {
	RoID     int    `mapstructure:"roId"`
	Domain   string `mapstructure:"domain"`
	Type     string `mapstructure:"type"`
	MasterIP string `mapstructure:"masterIp"`
	Mail     string `mapstructure:"mail"`
	Web      string `mapstructure:"web"`
	URL      string `mapstructure:"url"`
	Ipv4     string `mapstructure:"ipv4"`
	Ipv6     string `mapstructure:"ipv6"`
}
