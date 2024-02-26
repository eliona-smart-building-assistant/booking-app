/*
 * Booking App
 *
 * API to access and configure the Booking App
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package apiserver

type CreateBookingRequest struct {

	// The IDs of the assets to be booked.
	AssetIds []int32 `json:"assetIds"`

	// The ID (email) of the organizer.
	OrganizerID string `json:"organizerID"`

	// The start datetime of the booking in ISO 8601 format.
	Start string `json:"start"`

	// The end datetime of the booking in ISO 8601 format.
	End string `json:"end"`

	// The name of the event. (Optional)
	EventName string `json:"eventName,omitempty"`

	// A description of the event or booking. (Optional)
	Description string `json:"description,omitempty"`
}

// AssertCreateBookingRequestRequired checks if the required fields are not zero-ed
func AssertCreateBookingRequestRequired(obj CreateBookingRequest) error {
	elements := map[string]interface{}{
		"assetIds":    obj.AssetIds,
		"organizerID": obj.OrganizerID,
		"start":       obj.Start,
		"end":         obj.End,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertCreateBookingRequestConstraints checks if the values respects the defined constraints
func AssertCreateBookingRequestConstraints(obj CreateBookingRequest) error {
	return nil
}
