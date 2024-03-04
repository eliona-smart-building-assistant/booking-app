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

type SynchronizationAPIService struct {
}

func NewSynchronizationAPIService() apiserver.SynchronizationAPIServicer {
	return &SynchronizationAPIService{}
}

// SubscribeBookings - Open a WebSocket connection to get informed about newly created bookings.
func (s *SynchronizationAPIService) SubscribeBookings(ctx context.Context, subscribeBookingsRequest apiserver.SubscribeBookingsRequest) (apiserver.ImplResponse, error) {
	msgChan := getMessageChannelFromContext(ctx)
	for _, assetID := range subscribeBookingsRequest.AssetIDs {
		assetSubscriptions[assetID] = append(assetSubscriptions[assetID], subscriber{msgChan})
	}
	return apiserver.Response(http.StatusOK, nil), nil
}

// SyncBookingsPost - Post bookings from external service
func (s *SynchronizationAPIService) SyncBookingsPost(ctx context.Context, createBookingRequest []apiserver.CreateBookingRequest) (apiserver.ImplResponse, error) {
	// TODO - update SyncBookingsPost with the required logic for this service method.
	// Add api_synchronization_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	// TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	// return Response(400, nil),nil

	return apiserver.Response(http.StatusNotImplemented, nil), errors.New("SyncBookingsPost method not implemented")
}
