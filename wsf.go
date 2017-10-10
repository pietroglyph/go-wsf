package wsf

import "net/url"

const (
	libraryVersion = "010"

	defaultUserAgent = "go-wsf/" + libraryVersion

	defaultBaseURL = "http://www.wsdot.wa.gov/ferries/api/"
)

// A Client manages communication with the WSF API
type Client struct {
	// Base URL for API requests. Defaults to http://www.wsdot.wa.gov/ferries/api/.
	// BaseURL should always be specified with a trailing backslash.
	BaseURL *url.URL

	UserAgent string

	// Services for the different parts of the WSF API
	// TODO: Fares, Schedules, Terminals, and Vessels API.
}
