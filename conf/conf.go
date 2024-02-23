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
	"booking-app/appdb"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var ErrBadRequest = errors.New("bad request")
var ErrNotFound = errors.New("not found")

func InsertEvent(ctx context.Context, title string, description string, organizer string, startTime, endTime time.Time) error {
	var dbEvent appdb.Event
	dbEvent.Title = title
	dbEvent.Description = description
	dbEvent.Organizer = organizer
	dbEvent.StartTime = startTime
	dbEvent.EndTime = endTime
	return dbEvent.InsertG(ctx, boil.Infer())
}

func GetEventByID(ctx context.Context, eventID string) (*appdb.Event, error) {
	event, err := appdb.Events(
		appdb.EventWhere.ID.EQ(eventID),
	).OneG(ctx)
	if err != nil {
		return nil, fmt.Errorf("fetching event: %v", err)
	}
	return event, nil
}

func GetEventsForAsset(ctx context.Context, assetID int32, since, until time.Time) ([]*appdb.Event, error) {
	events, err := appdb.Events(
		qm.InnerJoin("booking.event_resource r on r.event_id = booking.event.id"),
		appdb.EventWhere.CancelledAt.IsNull(),
		appdb.EventWhere.EndTime.GT(since),
		appdb.EventWhere.StartTime.LT(until),
		appdb.EventResourceWhere.AssetID.EQ(assetID),
	).AllG(ctx)
	if err != nil {
		return nil, fmt.Errorf("fetching events for asset: %v", err)
	}
	return events, nil
}

func CancelEvent(ctx context.Context, eventID string) error {
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
