package wsf

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// VesselsService handles communication with the WSF Vessels API, which includes
// vessel attributes, locations, and other vessel-specific data. For the
// corresponding WSF REST API documentation, see http://www.wsdot.wa.gov/ferries/api/vessels/documentation/rest.html
type VesselsService service

// VesselLocations is an array of VesselLocation, which should include every
// vessel tracked by the endpoint.
type VesselLocations []VesselLocation

// VesselLocation is the location and related data for a single vessel.
type VesselLocation struct {
	VesselID                int
	VesselName              string
	Mmsi                    int `json:",omitempty"`
	DepartingTerminalID     int
	DepartingTerminalName   string
	DepartingTerminalAbbrev string
	ArrivingTerminalID      int    `json:",omitempty"`
	ArrivingTerminalName    string `json:",omitempty"`
	ArrivingTerminalAbbrev  string `json:",omitempty"`
	Latitude                float64
	Longitude               float64
	Speed                   float64
	Heading                 float64
	InService               bool
	AtDock                  bool
	LeftDock                Time   `json:"LeftDock, string, omitempty"`
	Eta                     Time   `json:"Eta, string, omitempty"`
	EtaBasis                string `json:",omitempty"`
	ScheduledDeparture      Time   `json:"ScheduledDeparture, string, omitempty"`
	OpRouteAbbrev           []string
	VesselPositionNum       int `json:",omitempty"`
	SortSeq                 int
	ManagedBy               int  // Enum, 1 for WSF, and 2 for KCM
	TimeStamp               Time `json:"TimeStamp, string"`
}

// Time implements a custom JSON unmarsaller for the specific format of
// non-RFC 3339 time output by the WSF API. Use a type assertion to get to the
// underlying time.Time type.
type Time time.Time

func (s *VesselsService) VesselLocations() *VesselLocations {
	// TODO: GET at endpoint and unmarshall
}

// UnmarshalJSON unmarshalls the time output by the WSF API
func (t *Time) UnmarshalJSON(b []byte) error {
	// Return on "null" data like the standard library does
	if string(b) == "null" {
		return nil
	}

	// First get rid of the \/Date() portion
	truncated := strings.TrimSuffix(strings.TrimPrefix(string(b), "\"\\/Date("), ")\\/\"") // This is because we have a capture group, and we need to find a submatch

	// Then separate the epoch time from the time zone
	timeSplit := strings.Split(truncated, "-")

	// Make sure that there are the correct number of dashes
	if len(timeSplit) > 2 {
		return errors.New("ASP.NET time submatch had too many dashes")
	}

	// Parse the unix time into an int
	i, err := strconv.ParseInt(timeSplit[0], 10, 64)
	if err != nil {
		return err
	}

	parsedTime := time.Unix(0, i*1000000) // i is in milleseconds so we need to convert to nano, hence the multiplication

	*t = Time{parsedTime}
	return nil
}
