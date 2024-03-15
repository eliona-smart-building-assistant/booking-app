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
	"fmt"
	"net/http"
)

type SynchronizationAPIService struct {
}

func NewSynchronizationAPIService() apiserver.SynchronizationAPIServicer {
	return &SynchronizationAPIService{}
}

// SubscribeBookings - Open a WebSocket connection to get informed about newly created bookings.
func (s *SynchronizationAPIService) SubscribeBookings(ctx context.Context, subscribeBookingsRequest apiserver.SubscribeBookingsRequest) (apiserver.ImplResponse, error) {
	// This method should be handled by websocket instead.
	return apiserver.Response(http.StatusTeapot, nil), nil
}

// SyncBookingsPost - Post bookings from external service
func (s *SynchronizationAPIService) SyncBookingsPost(ctx context.Context, createBookingsRequest []apiserver.CreateBookingRequest) (apiserver.ImplResponse, error) {
	for _, req := range createBookingsRequest {
		if err := processBooking(ctx, req); err != nil {
			return apiserver.Response(http.StatusBadRequest, err.Error()), err
		}
	}
	return apiserver.Response(http.StatusCreated, fmt.Sprintf("%v bookings created successfully", len(createBookingsRequest))), nil
}
