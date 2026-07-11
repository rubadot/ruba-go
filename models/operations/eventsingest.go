package operations

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type EventsIngestResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	EventsIngestResponse *components.EventsIngestResponse
}

func (e *EventsIngestResponse) GetHTTPMeta() components.HTTPMetadata {
	if e == nil {
		return components.HTTPMetadata{}
	}
	return e.HTTPMeta
}

func (e *EventsIngestResponse) GetEventsIngestResponse() *components.EventsIngestResponse {
	if e == nil {
		return nil
	}
	return e.EventsIngestResponse
}
