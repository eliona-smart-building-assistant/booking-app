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

	"github.com/gorilla/mux"
)

// BookingAPIController binds http requests to an api service and writes the service results to the http response
type BookingAPIController struct {
	service      BookingAPIServicer
	errorHandler ErrorHandler
}

// BookingAPIOption for how the controller is set up.
type BookingAPIOption func(*BookingAPIController)

// WithBookingAPIErrorHandler inject ErrorHandler into controller
func WithBookingAPIErrorHandler(h ErrorHandler) BookingAPIOption {
	return func(c *BookingAPIController) {
		c.errorHandler = h
	}
}

// NewBookingAPIController creates a default api controller
func NewBookingAPIController(s BookingAPIServicer, opts ...BookingAPIOption) Router {
	controller := &BookingAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the BookingAPIController
func (c *BookingAPIController) Routes() Routes {
	return Routes{
		"BookingsBookingIdDelete": Route{
			strings.ToUpper("Delete"),
			"/v1/bookings/{bookingId}",
			c.BookingsBookingIdDelete,
		},
		"BookingsGet": Route{
			strings.ToUpper("Get"),
			"/v1/bookings",
			c.BookingsGet,
		},
		"BookingsPost": Route{
			strings.ToUpper("Post"),
			"/v1/bookings",
			c.BookingsPost,
		},
	}
}

// BookingsBookingIdDelete - Cancel a booking
func (c *BookingAPIController) BookingsBookingIdDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookingIdParam, err := parseNumericParameter[int32](
		params["bookingId"],
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.BookingsBookingIdDelete(r.Context(), bookingIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// BookingsGet - List bookings
func (c *BookingAPIController) BookingsGet(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	var startParam string
	if query.Has("start") {
		param := query.Get("start")

		startParam = param
	} else {
		c.errorHandler(w, r, &RequiredError{Field: "start"}, nil)
		return
	}
	var endParam string
	if query.Has("end") {
		param := query.Get("end")

		endParam = param
	} else {
		c.errorHandler(w, r, &RequiredError{Field: "end"}, nil)
		return
	}
	var assetIdParam int32
	if query.Has("assetId") {
		param, err := parseNumericParameter[int32](
			query.Get("assetId"),
			WithParse[int32](parseInt32),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: err}, nil)
			return
		}

		assetIdParam = param
	} else {
		c.errorHandler(w, r, &RequiredError{Field: "assetId"}, nil)
		return
	}
	result, err := c.service.BookingsGet(r.Context(), startParam, endParam, assetIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// BookingsPost - Create a booking
func (c *BookingAPIController) BookingsPost(w http.ResponseWriter, r *http.Request) {
	createBookingRequestParam := CreateBookingRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&createBookingRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertCreateBookingRequestRequired(createBookingRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertCreateBookingRequestConstraints(createBookingRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.BookingsPost(r.Context(), createBookingRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
