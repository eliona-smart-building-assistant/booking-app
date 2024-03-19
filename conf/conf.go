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

package conf

import (
	"booking-app/apiserver"
	"booking-app/appdb"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var ErrBadRequest = errors.New("bad request")
var ErrNotFound = errors.New("not found")

func GetConfig(ctx context.Context) (apiserver.Configuration, error) {
	dbConfig, err := appdb.Configurations().OneG(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return apiserver.Configuration{}, ErrNotFound
	}
	if err != nil {
		return apiserver.Configuration{}, fmt.Errorf("fetching config from database: %v", err)
	}
	return apiConfigFromDbConfig(*dbConfig), nil
}

func UpsertConfig(ctx context.Context, config apiserver.Configuration) (apiserver.Configuration, error) {
	dbConfig := dbConfigFromApiConfig(config)
	dbConfig.ID = 1
	if err := dbConfig.UpsertG(ctx, true, []string{"id"}, boil.Blacklist("id"), boil.Infer()); err != nil {
		return apiserver.Configuration{}, fmt.Errorf("upserting DB config: %v", err)
	}
	return config, nil
}

func dbConfigFromApiConfig(apiConfig apiserver.Configuration) appdb.Configuration {
	return appdb.Configuration{
		ID:                 1,
		StartBookableHours: apiConfig.DayStartHours,
		StartBookableMins:  apiConfig.DayStartMins,
		EndBookableHours:   apiConfig.DayEndHours,
		EndBookableMins:    apiConfig.DayEndMins,
	}
}

func apiConfigFromDbConfig(dbConfig appdb.Configuration) apiserver.Configuration {
	return apiserver.Configuration{
		DayStartHours: dbConfig.StartBookableHours,
		DayStartMins:  dbConfig.StartBookableMins,
		DayEndHours:   dbConfig.EndBookableHours,
		DayEndMins:    dbConfig.EndBookableMins,
	}
}

func InsertEvent(ctx context.Context, assetIDs []int32, organizer string, startTime, endTime time.Time) error {
	dbEvent := &appdb.Event{
		Organizer: organizer,
		StartTime: startTime,
		EndTime:   endTime,
		CreatedAt: time.Now(),
	}

	// Using a transaction to ensure atomicity of the insert operations
	tx, err := boil.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := dbEvent.Insert(ctx, tx, boil.Infer()); err != nil {
		return err
	}

	for _, assetID := range assetIDs {
		dbEventResource := &appdb.EventResource{
			EventID: dbEvent.ID,
			AssetID: assetID,
		}

		if err := dbEventResource.Insert(ctx, tx, boil.Infer()); err != nil {
			return err
		}
	}

	return tx.Commit()
}

func GetEventsForAsset(ctx context.Context, assetID int32, since, until time.Time) ([]apiserver.Booking, error) {
	events, err := appdb.Events(
		qm.InnerJoin("booking.event_resource r on r.event_id = booking.event.id"),
		qm.Where("booking.event.cancelled_at IS NULL"),
		qm.And("r.asset_id = ?", assetID),
		qm.And("booking.event.start_time <= ?", until),
		qm.And("booking.event.end_time >= ?", since),
	).AllG(ctx)
	if err != nil {
		return nil, fmt.Errorf("fetching events for asset: %v", err)
	}
	var apiBookings []apiserver.Booking
	for _, e := range events {
		b, err := dbEventToAPIBooking(ctx, *e)
		if err != nil {
			return nil, fmt.Errorf("dbEvent to API booking: %v", err)
		}
		apiBookings = append(apiBookings, b)
	}
	return apiBookings, nil
}

func dbEventToAPIBooking(ctx context.Context, e appdb.Event) (apiserver.Booking, error) {
	res, err := e.EventResources().AllG(ctx)
	if err != nil {
		return apiserver.Booking{}, fmt.Errorf("fetching resources for event %v: %v", e.ID, err)
	}
	var assetIDs []int32
	for _, r := range res {
		assetIDs = append(assetIDs, r.AssetID)
		fmt.Println(assetIDs)
	}
	return apiserver.Booking{
		Id:          int32(e.ID),
		AssetIds:    assetIDs,
		Start:       e.StartTime,
		End:         e.EndTime,
		OrganizerID: e.Organizer,
	}, nil
}

// This would be better, to have a static type checking. But there is a runtime error:
// "failed to assign all query results to Event slice: bind failed to execute query: pq: invalid reference to FROM-clause entry for table \"event_resource\""
// func GetEventsForAsset2(ctx context.Context, assetID int32, since, until time.Time) ([]*appdb.Event, error) {
// 	events, err := appdb.Events(
// 		qm.Select(),
// 		qm.InnerJoin("booking.event_resource r on r.event_id = booking.event.id"),
// 		appdb.EventWhere.CancelledAt.IsNull(),
// 		appdb.EventWhere.EndTime.GT(since),
// 		appdb.EventWhere.StartTime.LT(until),
// 		appdb.EventResourceWhere.AssetID.EQ(assetID),
// 	).AllG(ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("fetching events for asset: %v", err)
// 	}
// 	return events, nil
// }

func CancelEvent(ctx context.Context, eventID int64) error {
	event, err := appdb.Events(
		appdb.EventWhere.ID.EQ(eventID),
	).OneG(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return ErrNotFound
	}
	if err != nil {
		return fmt.Errorf("fetching event to cancel: %v", err)
	}

	event.CancelledAt = null.TimeFrom(time.Now())

	if _, err := event.UpdateG(ctx, boil.Whitelist(appdb.EventColumns.CancelledAt)); err != nil {
		return fmt.Errorf("updating event to mark as cancelled: %v", err)
	}

	return nil
}

// func GetCurrentEvents(ctx context.Context) ([]apiserver.Booking, error) {
// 	since := time.Now()
// 	until := since.Add(30 * time.Minute)
// 	events, err := appdb.Events(
// 		appdb.EventWhere.CancelledAt.IsNull(),
// 		appdb.EventWhere.EndTime.GTE(since),
// 		appdb.EventWhere.StartTime.LTE(until),
// 	).AllG(ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("fetching events for asset: %v", err)
// 	}
// 	var apiBookings []apiserver.Booking
// 	for _, e := range events {
// 		b, err := dbEventToAPIBooking(ctx, *e)
// 		if err != nil {
// 			return nil, fmt.Errorf("dbEvent to API booking: %v", err)
// 		}
// 		apiBookings = append(apiBookings, b)
// 	}
// 	return apiBookings, nil
// }

type AssetID struct {
	AssetID int32 `boil:"asset_id"`
}

func GetBookedAssetIDs(ctx context.Context) ([]int32, error) {
	since := time.Now()
	until := since.Add(30 * time.Minute)
	var assetIDs []*AssetID
	if err := queries.
		RawG(`
			SELECT DISTINCT er.asset_id
			FROM booking.event_resource er
			JOIN booking.event e ON e.id = er.event_id
			WHERE e.cancelled_at IS NULL
				AND e.start_time <= $1
				AND e.end_time >= $2
			`, until, since).
		BindG(ctx, &assetIDs); err != nil {
		return nil, fmt.Errorf("querying booked asset IDs: %v", err)
	}

	ids := make([]int32, len(assetIDs))
	for i, assetID := range assetIDs {
		ids[i] = assetID.AssetID
	}
	return ids, nil
}

func GetUnbookedAssetIDs(ctx context.Context) ([]int32, error) {
	since := time.Now()
	until := since.Add(30 * time.Minute)
	var assetIDs []*AssetID
	if err := queries.
		RawG(`
			SELECT DISTINCT er.asset_id
			FROM booking.event_resource er
			JOIN booking.event e ON e.id = er.event_id
			WHERE e.cancelled_at IS NOT NULL
				OR e.start_time > $1
				OR e.end_time < $2
			`, until, since).
		BindG(ctx, &assetIDs); err != nil {
		return nil, fmt.Errorf("querying unbooked asset IDs: %v", err)
	}
	ids := make([]int32, len(assetIDs))
	for i, assetID := range assetIDs {
		ids[i] = assetID.AssetID
	}
	return ids, nil
}
