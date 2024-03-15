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
	"booking-app/conf"
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
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

// BookingsBookingIdDelete - Cancel a booking
func (s *BookingAPIService) BookingsBookingIdDelete(ctx context.Context, bookingId int32) (apiserver.ImplResponse, error) {
	if err := conf.CancelEvent(ctx, int64(bookingId)); errors.Is(err, conf.ErrNotFound) {
		return apiserver.ImplResponse{Code: http.StatusNotFound}, err
	} else if err != nil {
		return apiserver.ImplResponse{Code: http.StatusInternalServerError}, err
	}
	return apiserver.Response(http.StatusOK, nil), nil
}

// BookingsGet - List bookings
func (s *BookingAPIService) BookingsGet(ctx context.Context, start string, end string, assetId int32) (apiserver.ImplResponse, error) {
	since, err := time.Parse(time.RFC3339, start)
	if err != nil {
		return apiserver.Response(http.StatusBadRequest, "Invalid start time format"), fmt.Errorf("error parsing start time: %v", err)
	}
	until, err := time.Parse(time.RFC3339, end)
	if err != nil {
		return apiserver.Response(http.StatusBadRequest, "Invalid end time format"), fmt.Errorf("error parsing end time: %v", err)
	}
	events, err := conf.GetEventsForAsset(ctx, assetId, since, until)
	if err != nil {
		return apiserver.ImplResponse{Code: http.StatusInternalServerError}, err
	}
	return apiserver.Response(http.StatusOK, events), nil
}

// BookingsPost - Creates a new booking
func (s *BookingAPIService) BookingsPost(ctx context.Context, req apiserver.CreateBookingRequest) (apiserver.ImplResponse, error) {
	if err := processBooking(ctx, req); err != nil {
		return apiserver.Response(http.StatusBadRequest, err.Error()), err
	}

	for _, assetID := range req.AssetIds {
		for _, subscriber := range assetSubscriptions[assetID] {
			subscriber.conn.WriteJSON(req)
		}
	}

	return apiserver.Response(http.StatusCreated, "Booking created successfully"), nil
}

func processBooking(ctx context.Context, req apiserver.CreateBookingRequest) error {
	startTime, err := time.Parse(time.RFC3339, req.Start)
	if err != nil {
		return fmt.Errorf("invalid start time format: %w", err)
	}
	endTime, err := time.Parse(time.RFC3339, req.End)
	if err != nil {
		return fmt.Errorf("invalid end time format: %w", err)
	}
	if !endTime.After(startTime) {
		return errors.New("end time must be after start time")
	}

	if err = conf.InsertEvent(ctx, req.AssetIds, req.EventName, req.Description, "OrganizerName", startTime, endTime); err != nil {
		return fmt.Errorf("failed to insert event: %w", err)
	}
	return nil
}
