//  This file is part of the eliona project.
//  Copyright © 2022 LEICOM iTEC AG. All Rights Reserved.
//  ______ _ _
// |  ____| (_)
// | |__  | |_  ___  _ __   __ _
// |  __| | | |/ _ \| '_ \ / _` |
// | |____| | | (_) | | | | (_| |
// |______|_|_|\___/|_| |_|\__,_|
//
//  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING
//  BUT NOT LIMITED  TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
//  NON INFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
//  DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package apiservices

import (
	"booking-app/apiserver"
	"context"
	"errors"
	"net/http"
)

// BookingAPIService is a service that implements the logic for the BookingAPIServicer
// This service should implement the business logic for every endpoint for the BookingAPI API.
// Include any external packages or services that will be required by this service.
type BookingAPIService struct {
}

// NewBookingAPIService creates a default api service
func NewBookingAPIService() apiserver.BookingAPIServicer {
	return &BookingAPIService{}
}

// BookingsBookingIdDeletePost - Cancel a booking
func (s *BookingAPIService) BookingsBookingIdDeletePost(ctx context.Context, bookingId string, deleteBookingRequest apiserver.DeleteBookingRequest) (apiserver.ImplResponse, error) {
	return apiserver.Response(http.StatusNotImplemented, nil), errors.New("BookingsBookingIdDeletePost method not implemented")
}

// BookingsBookingIdRegisterGuestPost - Notify event organizer that a guest came for the event.
func (s *BookingAPIService) BookingsBookingIdRegisterGuestPost(ctx context.Context, bookingId string, bookingsBookingIdRegisterGuestPostRequest apiserver.BookingsBookingIdRegisterGuestPostRequest) (apiserver.ImplResponse, error) {
	// Registering guests will not be implemented yet. It is only used in Portier,
	// which is only in Leicom currently.
	return apiserver.Response(http.StatusNotImplemented, nil), errors.New("BookingsBookingIdRegisterGuestPost method not implemented")
}

// BookingsGet - List bookings
func (s *BookingAPIService) BookingsGet(ctx context.Context, start string, end string, assetId string) (apiserver.ImplResponse, error) {
	return apiserver.Response(http.StatusNotImplemented, nil), errors.New("BookingsGet method not implemented")
}

// BookingsPost - Create a booking
func (s *BookingAPIService) BookingsPost(ctx context.Context, createBookingRequest apiserver.CreateBookingRequest) (apiserver.ImplResponse, error) {
	return apiserver.Response(http.StatusNotImplemented, nil), errors.New("BookingsPost method not implemented")
}
