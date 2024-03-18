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

package eliona

import (
	"fmt"

	"github.com/eliona-smart-building-assistant/go-eliona/asset"
)

const ClientReference string = "booking-app"

type BookableAsset struct {
	Occupancy int `eliona:"occupancy" subtype:"input"`
}

func SetAssetsBooked(booked bool, assetIDs []int32) error {
	occupancy := 0
	if booked {
		occupancy = 1
	}
	for _, assetID := range assetIDs {
		data := asset.Data{
			AssetId:         assetID,
			Data:            BookableAsset{Occupancy: occupancy},
			ClientReference: ClientReference,
		}
		if err := asset.UpsertAssetDataIfAssetExists(data); err != nil {
			return fmt.Errorf("upserting data: %v", err)
		}
	}
	return nil
}
