/*
 * Booking App
 *
 * API to access and configure the Booking App
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package apiserver

import (
	"encoding/json"
	"net/http"
	"strings"
)

// SynchronizationAPIController binds http requests to an api service and writes the service results to the http response
type SynchronizationAPIController struct {
	service      SynchronizationAPIServicer
	errorHandler ErrorHandler
}

// SynchronizationAPIOption for how the controller is set up.
type SynchronizationAPIOption func(*SynchronizationAPIController)

// WithSynchronizationAPIErrorHandler inject ErrorHandler into controller
func WithSynchronizationAPIErrorHandler(h ErrorHandler) SynchronizationAPIOption {
	return func(c *SynchronizationAPIController) {
		c.errorHandler = h
	}
}

// NewSynchronizationAPIController creates a default api controller
func NewSynchronizationAPIController(s SynchronizationAPIServicer, opts ...SynchronizationAPIOption) Router {
	controller := &SynchronizationAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the SynchronizationAPIController
func (c *SynchronizationAPIController) Routes() Routes {
	return Routes{
		"SubscribeBookings": Route{
			strings.ToUpper("Get"),
			"/v1/sync/bookings-subscription",
			c.SubscribeBookings,
		},
		"SyncBookingsPost": Route{
			strings.ToUpper("Post"),
			"/v1/sync/bookings",
			c.SyncBookingsPost,
		},
	}
}

// SubscribeBookings - Open a WebSocket connection to get informed about newly created bookings for given assetIDs.
func (c *SynchronizationAPIController) SubscribeBookings(w http.ResponseWriter, r *http.Request) {
	subscribeBookingsRequestParam := SubscribeBookingsRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&subscribeBookingsRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertSubscribeBookingsRequestRequired(subscribeBookingsRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertSubscribeBookingsRequestConstraints(subscribeBookingsRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.SubscribeBookings(r.Context(), subscribeBookingsRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// SyncBookingsPost - Post bookings from external service
func (c *SynchronizationAPIController) SyncBookingsPost(w http.ResponseWriter, r *http.Request) {
	createBookingRequestParam := []CreateBookingRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&createBookingRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	for _, el := range createBookingRequestParam {
		if err := AssertCreateBookingRequestRequired(el); err != nil {
			c.errorHandler(w, r, err, nil)
			return
		}
	}
	result, err := c.service.SyncBookingsPost(r.Context(), createBookingRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
