package wsf

import (
	"net/http"
	"net/url"
)

const (
	libraryVersion = "010"

	defaultUserAgent = "go-wsf/" + libraryVersion

	defaultBaseURL = "http://www.wsdot.wa.gov/ferries/api/"
)

// A Client manages communication with the WSF API.
type Client struct {
	httpClient *http.Client // HTTP client used to communicate with the API.

	// Base URL for API requests. Defaults to http://www.wsdot.wa.gov/ferries/api/.
	// BaseURL should always be specified with a trailing backslash.
	BaseURL *url.URL

	UserAgent string

	// Access code for the WSF API, provisioned at http://www.wsdot.wa.gov/traffic/api/.
	// Requests will fail unless this is defined.
	AccessCode string

	common service // Reuse this single struct instead of having to allocate one for each service.

	//Different parts of the WSF API, referred to herin as Services.
	Vessels *VesselsService
}

type service struct {
	client *Client
}

// NewClient returns a new WSF API client. If a nil httpClient is privided then
// then http.DefaultClient will be used.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{httpClient: httpClient, BaseURL, baseURL, UserAgent: defaultUserAgent, AccessCode: ""}
	c.common.client = c
	c.Vessels = (*VesselsService)(&c.common)
}
