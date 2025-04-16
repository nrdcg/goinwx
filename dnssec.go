package goinwx

import (
	"errors"
	"time"

	"github.com/fatih/structs"
	"github.com/go-viper/mapstructure/v2"
)

const (
	methodDNSSecAddDNSKey     = "dnssec.adddnskey"
	methodDNSSecDeleteAll     = "dnssec.deleteall"
	methodDNSSecDeleteDNSKey  = "dnssec.deletednskey"
	methodDNSSecDisableDNSSec = "dnssec.disablednssec"
	methodDNSSecEnableDNSSec  = "dnssec.enablednssec"
	methodDNSSecInfo          = "dnssec.info"
	methodDNSSecListKeys      = "dnssec.listkeys"
)

// DNSSecService API access to DNSSEC.
type DNSSecService service

// Add adds one DNSKEY to a specified domain.
func (s *DNSSecService) Add(request *DNSSecAddRequest) (*DNSSecAddResponse, error) {
	if request == nil {
		return nil, errors.New("request can't be nil")
	}

	requestMap := structs.Map(request)

	req := s.client.NewRequest(methodDNSSecAddDNSKey, requestMap)

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result DNSSecAddResponse

	err = mapstructure.Decode(resp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteAll deletes all DNSKEY/DS entries for a domain.
func (s *DNSSecService) DeleteAll(domain string) error {
	req := s.client.NewRequest(methodDNSSecDeleteAll, map[string]interface{}{
		"domainName": domain,
	})

	_, err := s.client.Do(req)
	if err != nil {
		return err
	}

	return nil
}

// DeleteDNSKey deletes one DNSKEY from a specified domain.
func (s *DNSSecService) DeleteDNSKey(key string) error {
	req := s.client.NewRequest(methodDNSSecDeleteDNSKey, map[string]interface{}{
		"key": key,
	})

	_, err := s.client.Do(req)
	if err != nil {
		return err
	}

	return nil
}

// Disable disables automated DNSSEC management for a domain.
func (s *DNSSecService) Disable(domain string) error {
	req := s.client.NewRequest(methodDNSSecDisableDNSSec, map[string]interface{}{
		"domainName": domain,
	})

	_, err := s.client.Do(req)
	if err != nil {
		return err
	}

	return nil
}

// Enable enables automated DNSSEC management for a domain.
func (s *DNSSecService) Enable(domain string) error {
	req := s.client.NewRequest(methodDNSSecEnableDNSSec, map[string]interface{}{
		"domainName": domain,
	})

	_, err := s.client.Do(req)
	if err != nil {
		return err
	}

	return nil
}

// Info gets current DNSSEC information.
func (s *DNSSecService) Info(domains []string) (*DNSSecInfoResponse, error) {
	req := s.client.NewRequest(methodDNSSecInfo, map[string]interface{}{
		"domains": domains,
	})

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result DNSSecInfoResponse

	err = mapstructure.Decode(resp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// List lists domains.
func (s *DNSSecService) List(request *DNSSecServiceListRequest) (*DNSSecServiceList, error) {
	if request == nil {
		return nil, errors.New("request can't be nil")
	}

	requestMap := structs.Map(request)

	req := s.client.NewRequest(methodDNSSecListKeys, requestMap)

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result DNSSecServiceList

	err = mapstructure.Decode(resp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// DNSSecAddRequest API model.
type DNSSecAddRequest struct {
	DomainName      string `structs:"domainName,omitempty"`
	DNSKey          string `structs:"dnskey,omitempty"`
	DS              string `structs:"ds,omitempty"`
	CalculateDigest bool   `structs:"calculateDigest,omitempty"`
	DigestType      int    `structs:"digestType,omitempty"`
}

// DNSSecAddResponse API model.
type DNSSecAddResponse struct {
	DNSKey string `mapstructure:"dnskey"`
	DS     string `mapstructure:"ds"`
}

// DNSSecInfoResponse API model.
type DNSSecInfoResponse struct {
	Data []DNSSecInfo `mapstructure:"data"`
}

// DNSSecInfo API model.
type DNSSecInfo struct {
	Domain       string `mapstructure:"domain"`
	KeyCount     int    `mapstructure:"keyCount"`
	DNSSecStatus string `mapstructure:"dnsSecStatus"`
}

// DNSSecServiceListRequest API model.
type DNSSecServiceListRequest struct {
	DomainName    string `structs:"domainName,omitempty"`
	DomainNameIdn string `structs:"domainNameIdn,omitempty"`
	KeyTag        int    `structs:"keyTag,omitempty"`
	FlagID        int    `structs:"flagId,omitempty"`
	AlgorithmID   int    `structs:"algorithmId,omitempty"`
	PublicKey     string `structs:"publicKey,omitempty"`
	DigestTypeID  int    `structs:"digestTypeId,omitempty"`
	Digest        string `structs:"digest,omitempty"`
	CreatedBefore string `structs:"createdBefore,omitempty"`
	CreatedAfter  string `structs:"createdAfter,omitempty"`
	Status        string `structs:"status,omitempty"`
	Active        int    `structs:"active,omitempty"`
	Page          int    `structs:"page,omitempty"`
	PageLimit     int    `structs:"pagelimit,omitempty"`
}

// DNSSecServiceList API model.
type DNSSecServiceList struct {
	DNSKeys []DNSSecServiceListResponse `mapstructure:"dnskey"`
}

// DNSSecServiceListResponse API model.
type DNSSecServiceListResponse struct {
	OwnerName    string    `mapstructure:"ownerName"`
	ID           int       `mapstructure:"id"`
	DomainID     int       `mapstructure:"domainId"`
	KeyTag       int       `mapstructure:"keyTag"`
	FlagID       int       `mapstructure:"flagId"`
	AlgorithmID  int       `mapstructure:"algorithmId"`
	PublicKey    string    `mapstructure:"publicKey"`
	DigestTypeID int       `mapstructure:"digestTypeId"`
	Digest       string    `mapstructure:"digest"`
	Created      time.Time `mapstructure:"created"`
	Status       string    `mapstructure:"status"`
	Active       int       `mapstructure:"active"`
}
