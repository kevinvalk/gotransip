# gotransip - TransIP API client for Go
[![Go Report Card](https://goreportcard.com/badge/github.com/transip/gotransip)][goreport] [![Documentation](https://godoc.org/github.com/transip/gotransip?status.svg)][doc]

This is the Go client for the [TransIP API][api]. To use it you need an account with [TransIP][transip], enable API usage and setup a private API key.

**NOTE**: We keep the PHP version and golang version close as possible, but if something is not working 100% like you'd expect, please open an issue and of course: you're welcome to [contribute](CONTRIBUTING.md)!

## Example usage
To print a list of your account's VPSes:
```go
package main

import (
	"fmt"
	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/vps"
	"os"
)

func main() {
	// create new TransIP API SOAP client
	file, err := os.Open("/path/to/api/private.key")
	if err != nil {
		panic(err.Error())
	}
	client, err := gotransip.NewClient(gotransip.ClientConfiguration{
		AccountName:      "accountName",
		PrivateKeyReader: file,
	})
	if err != nil {
		panic(err.Error())
	}
	vpsRepo := vps.Repository{Client:client}

	// get list of VPSes
	vpss, err := vpsRepo.GetAll()
	if err != nil {
		panic(err.Error())
	}

	// get list of private networks
	pns, err := vpsRepo.GetPrivateNetworks()
	if err != nil {
		panic(err.Error())
	}

	// print name and description for each VPS
	for _, v := range vpss {
		fmt.Printf("vps: %s (%s)\n", v.Name, v.Description)
	}

	// print name and description for each private network
	for _, pn := range pns {
		fmt.Printf("privatenetwork: %s (%s)\n", pn.Name, pn.Description)
	}

	err = vpsRepo.Order(vps.VpsOrder{
		ProductName:       "vps-bladevps-x8",
		OperatingSystem:   "ubuntu-18.04",
		AvailabilityZone:  "ams0",
		Hostname:          "webserver01",
		Description:       "my-unique-description",
	})
	if err != nil {
		panic(err.Error())
	}
}
```

## Documentation
For detailed descriptions of all functions, check out the [TransIP API documentation][apidoc]. Details about the usage of the Go client can be found on [pkg.go.dev][doc].

[transip]: https://transip.nl/
[api]: https://api.transip.nl/
[doc]: https://pkg.go.dev/github.com/transip/gotransip?tab=doc
[apidoc]: https://api.transip.nl/rest/docs.html
[goreport]: https://goreportcard.com/report/github.com/transip/gotransip