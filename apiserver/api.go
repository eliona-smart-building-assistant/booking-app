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
	"context"
	"net/http"
)

// BookingAPIRouter defines the required methods for binding the api requests to a responses for the BookingAPI
// The BookingAPIRouter implementation should parse necessary information from the http request,
// pass the data to a BookingAPIServicer to perform the required actions, then write the service results to the http response.
type BookingAPIRouter interface {
	BookingsBookingIdDeletePost(http.ResponseWriter, *http.Request)
	BookingsBookingIdRegisterGuestPost(http.ResponseWriter, *http.Request)
	BookingsGet(http.ResponseWriter, *http.Request)
	BookingsPost(http.ResponseWriter, *http.Request)
}

// VersionAPIRouter defines the required methods for binding the api requests to a responses for the VersionAPI
// The VersionAPIRouter implementation should parse necessary information from the http request,
// pass the data to a VersionAPIServicer to perform the required actions, then write the service results to the http response.
type VersionAPIRouter interface {
	GetOpenAPI(http.ResponseWriter, *http.Request)
	GetVersion(http.ResponseWriter, *http.Request)
}

// BookingAPIServicer defines the api actions for the BookingAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type BookingAPIServicer interface {
	BookingsBookingIdDeletePost(context.Context, string, DeleteBookingRequest) (ImplResponse, error)
	BookingsBookingIdRegisterGuestPost(context.Context, string, BookingsBookingIdRegisterGuestPostRequest) (ImplResponse, error)
	BookingsGet(context.Context, string, string, string) (ImplResponse, error)
	BookingsPost(context.Context, CreateBookingRequest) (ImplResponse, error)
}

// VersionAPIServicer defines the api actions for the VersionAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type VersionAPIServicer interface {
	GetOpenAPI(context.Context) (ImplResponse, error)
	GetVersion(context.Context) (ImplResponse, error)
}
