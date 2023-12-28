package knownawsaccountslookup

import (
	"errors"
	"io"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

type Vendors []struct {
	Name     string       `yaml:"name"`
	Source   ListOrString `yaml:"source"`
	Accounts ListOrString `yaml:"accounts"`
}

type ListOrString struct {
	Values []string
}

func (ls *ListOrString) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var single string
	if err := unmarshal(&single); err == nil {
		ls.Values = []string{single}
		return nil
	}

	var slice []string
	if err := unmarshal(&slice); err == nil {
		ls.Values = slice
		return nil
	}

	return errors.New("not a string or list of strings")
}

// Function that uses the constructor pattern. You can use this or instantiate your struct without it.
func NewVendorMap() *Vendors {
	return &Vendors{}
}

// This function downloads the accounts.yaml file from https://raw.githubusercontent.com/fwdcloudsec/known_aws_accounts/main/accounts.yaml and populates the Vendors variable
func (v *Vendors) PopulateKnownAWSAccounts() {
	// Fetch the YAML file from the URL
	resp, err := http.Get("https://raw.githubusercontent.com/fwdcloudsec/known_aws_accounts/main/accounts.yaml")
	if err != nil {
		log.Fatalf("error fetching YAML file: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error reading response body: %v", err)
	}

	// Unmarshal the YAML into the Vendors struct
	err = yaml.Unmarshal(body, v)
	if err != nil {
		log.Fatalf("error unmarshalling YAML: %v", err)
	}

}

// GetVendorNameFromAccountID retrieves the vendor name associated with the given account ID.
// It searches through the Vendors collection and returns the vendor name if a matching account ID is found.
// If no matching account ID is found, it returns an empty string.
func (v *Vendors) GetVendorNameFromAccountID(accountID string) string {

	for _, vendor := range *v {
		for _, account := range vendor.Accounts.Values {
			if account == accountID {
				return vendor.Name
			}
		}
	}
	return ""
}

// GetAccountIdsFromVendorName retrieves the account IDs associated with a specific vendor name.
// It searches through the Vendors collection and returns a slice of account IDs that match the given vendor name.
// If no matching vendor name is found, an empty slice is returned.
func (v *Vendors) GetAccountIdsFromVendorName(vendorName string) []string {
	var accountIds []string
	for _, vendor := range *v {
		if vendor.Name == vendorName {
			for _, account := range vendor.Accounts.Values {
				accountIds = append(accountIds, account)
			}
		}
	}
	return accountIds
}
