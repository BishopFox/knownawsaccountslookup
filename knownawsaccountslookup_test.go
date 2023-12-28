package knownawsaccountslookup

import (
	"log"
	"testing"
)

func TestGetVendorNameFromAccountID(t *testing.T) {
	subtests := []struct {
		name           string
		accountId      string
		expectedResult string
	}{
		{
			name:           "test1",
			accountId:      "454464851268",
			expectedResult: "Cloudhealth",
		},
		{
			name:           "test2",
			accountId:      "370134896156",
			expectedResult: "Fugue",
		},
		{
			name:           "test2",
			accountId:      "944830124550",
			expectedResult: "Fugue",
		},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			vendors := NewVendorMap()
			vendors.PopulateKnownAWSAccounts()
			result := vendors.GetVendorNameFromAccountID(subtest.accountId)
			if result != subtest.expectedResult {
				log.Fatal("Vendor name does not match expected name")
			}
		})
	}
}

func TestGetAccountIdsFromVendorName(t *testing.T) {
	subtests := []struct {
		name           string
		vendorName     string
		expectedResult []string
	}{
		{
			name:           "test1",
			vendorName:     "Cloudhealth",
			expectedResult: []string{"454464851268"},
		},
		{
			name:           "test2",
			vendorName:     "Fugue",
			expectedResult: []string{"370134896156", "944830124550"},
		},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			vendors := NewVendorMap()
			vendors.PopulateKnownAWSAccounts()
			result := vendors.GetAccountIdsFromVendorName(subtest.vendorName)
			if result[0] != subtest.expectedResult[0] {
				log.Fatal("AccountID does not match expected AccountID")
			}
		})
	}
}
