# INWX Go API client

This go library implements some parts of the official INWX XML-RPC API.

## API

```go
package main

import (
	"log"

	"github.com/nrdcg/goinwx"
)

func main() {
	client := goinwx.NewClient("username", "password", &goinwx.ClientOptions{Sandbox: true})

	err := client.Account.Login()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := client.Account.Logout(); err != nil {
			log.Printf("inwx: failed to logout: %v", err)
		}
	}()

	var request = &goinwx.NameserverRecordRequest{
		Domain:  "domain.com",
		Name:    "foo.domain.com.",
		Type:    "TXT",
		Content: "aaa",
		Ttl:     300,
	}

	_, err = client.Nameservers.CreateRecord(request)
	if err != nil {
		log.Fatal(err)
	}
}
```

## Supported Features

Full API documentation can be found [here](https://www.inwx.de/en/help/apidoc).

The following parts are implemented:
* Account
  * Login
  * Logout
  * Lock
  * Unlock (with mobile TAN)
* Domains
  * Check
  * Register
  * Delete
  * Info
  * GetPrices
  * List
  * Whois
* Nameservers
  * Check
  * Create
  * Info
  * List
  * CreateRecord
  * UpdateRecord
  * DeleteRecord
  * FindRecordById
* Contacts
  * List 
  * Info
  * Create
  * Update
  * Delete

## Contributions

Your contributions are very appreciated.
