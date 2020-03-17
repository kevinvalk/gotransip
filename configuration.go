package gotransip

import (
	"github.com/transip/gotransip/v6/authenticator"
	"io"
	"net/http"
)

const (
	libraryVersion = "6.0.0"
	basePath       = "https://api.transip.nl/v6"
	userAgent      = "go-client-gotransip/" + libraryVersion
)

// APIMode specifies in which mode the API is used. Currently this is only
// supports either readonly or readwrite
type APIMode string

var (
	// APIModeReadOnly specifies that no changes can be made from API calls
	APIModeReadOnly APIMode = "readonly"
	// APIModeReadWrite specifies that changes can be made from API calls
	APIModeReadWrite APIMode = "readwrite"
)

// DemoClientConfiguration is the default configuration to use when testing the demo mode of the transip api
var DemoClientConfiguration = ClientConfiguration{Token: authenticator.DemoToken}

// ClientConfiguration stores the configuration of the API client
type ClientConfiguration struct {
	AccountName      string
	URL              string
	PrivateKeyReader io.Reader
	Token            string
	TestMode         bool
	HTTPClient       *http.Client
	Mode             APIMode
}
