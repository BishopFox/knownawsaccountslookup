# Known AWS Accounts Lookup

This go module provides two lookup functions for the data in https://raw.githubusercontent.com/fwdcloudsec/known_aws_accounts/, which is owned and operated by the `fwdcloudsec` team. 

Data source of truth: https://raw.githubusercontent.com/fwdcloudsec/known_aws_accounts/main/accounts.yaml.

## Functions

* [GetVendorNameFromAccountID](#GetVendorNameFromAccountID)
* [GetAccountIdsFromVendorName](#GetAccountIdsFromVendorName)

## Examples

### GetVendorNameFromAccountID

```go
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
}
```


### GetAccountIdsFromVendorName

```go
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

	// Get the account IDs from the vendor name
	accountIds := vendors.GetAccountIdsFromVendorName("Fugue")
	fmt.Println(accountIds)
}
```