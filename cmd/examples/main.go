package main

import (
	"fmt"

	"github.com/bishopfox/knownawsaccountslookup"
)

func main() {
	// Instantiate the struct
	vendors := knownawsaccountslookup.NewVendorMap()

	// Call the function to populate the struct
	vendors.PopulateKnownAWSAccounts()

	// Get the vendor name from the account ID
	vendorName := vendors.GetVendorNameFromAccountID("454464851268")
	fmt.Println(vendorName)

	// Get the account IDs from the vendor name
	accountIds := vendors.GetAccountIdsFromVendorName("Fugue")
	fmt.Println(accountIds)

}
