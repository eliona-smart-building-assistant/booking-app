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

package main

import (
	"booking-app/apiserver"
	"booking-app/apiservices"
	"booking-app/conf"
	"booking-app/eliona"
	"context"
	"net/http"

	"github.com/eliona-smart-building-assistant/go-eliona/app"
	"github.com/eliona-smart-building-assistant/go-utils/common"
	"github.com/eliona-smart-building-assistant/go-utils/db"
	utilshttp "github.com/eliona-smart-building-assistant/go-utils/http"
	"github.com/eliona-smart-building-assistant/go-utils/log"
)

func manageOccupancy() {
	common.RunOnce(func() {
		if err := updateOccupancy(); err != nil {
			return // Error is handled in the method itself.
		}
	}, 1)
}

func updateOccupancy() error {
	ctx := context.Background()
	bookedAssets, err := conf.GetBookedAssetIDs(ctx)
	if err != nil {
		log.Error("conf", "getting booked asset IDs: %v", err)
		return err
	}
	if err := eliona.SetAssetsBooked(true, bookedAssets); err != nil {
		log.Error("eliona", "setting booked assets: %v", err)
		return err
	}
	unbookedAssets, err := conf.GetUnbookedAssetIDs(ctx)
	if err != nil {
		log.Error("conf", "getting booked asset IDs: %v", err)
		return err
	}
	if err := eliona.SetAssetsBooked(false, unbookedAssets); err != nil {
		log.Error("eliona", "setting unbooked assets: %v", err)
		return err
	}
	return nil
}

func initialization() {
	ctx := context.Background()

	// Necessary to close used init resources
	conn := db.NewInitConnectionWithContextAndApplicationName(ctx, app.AppName())
	defer conn.Close(ctx)

	// Init the app before the first run.
	app.Init(conn, app.AppName(),
		app.ExecSqlFile("conf/init.sql"),
	)
}

// listenApi starts the API server and listen for requests
func listenApi() {
	err := http.ListenAndServe(":"+common.Getenv("API_SERVER_PORT", "3000"), utilshttp.NewCORSEnabledHandler(
		apiservices.NewRouter(
			apiserver.NewConfigurationAPIController(apiservices.NewConfigurationAPIService()),
			apiserver.NewSynchronizationAPIController(apiservices.NewSynchronizationAPIService()),
			apiserver.NewVersionAPIController(apiservices.NewVersionApiService()),
			apiserver.NewBookingAPIController(apiservices.NewBookingAPIService()),
		)),
	)
	log.Fatal("main", "API server: %v", err)
}
