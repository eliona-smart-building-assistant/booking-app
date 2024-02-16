/*
 * Booking App
 *
 * API to access and configure the Booking App
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package apiserver

type DeleteBookingRequest struct {

	// Device code obtained from authorization
	DeviceCode string `json:"deviceCode"`
}

// AssertDeleteBookingRequestRequired checks if the required fields are not zero-ed
func AssertDeleteBookingRequestRequired(obj DeleteBookingRequest) error {
	elements := map[string]interface{}{
		"deviceCode": obj.DeviceCode,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertDeleteBookingRequestConstraints checks if the values respects the defined constraints
func AssertDeleteBookingRequestConstraints(obj DeleteBookingRequest) error {
	return nil
}
