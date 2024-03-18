--  This file is part of the eliona project.
--  Copyright © 2022 LEICOM iTEC AG. All Rights Reserved.
--  ______ _ _
-- |  ____| (_)
-- | |__  | |_  ___  _ __   __ _
-- |  __| | | |/ _ \| '_ \ / _` |
-- | |____| | | (_) | | | | (_| |
-- |______|_|_|\___/|_| |_|\__,_|
--
--  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING
--  BUT NOT LIMITED  TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
--  NON INFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
--  DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
--  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

create schema if not exists booking;

create table if not exists booking.configuration
(
    id bigserial primary key CHECK (id = 1), -- The app supports only one configuration now.
    start_bookable_hours int NOT NULL,
    start_bookable_mins int NOT NULL,
    end_bookable_hours int NOT NULL,
    end_bookable_mins int NOT NULL,
    CHECK (end_bookable_hours * 60 + end_bookable_mins > start_bookable_hours * 60 + start_bookable_mins)
);

create table if not exists booking.event (
    id bigserial primary key,
    organizer text not null,
    start_time timestamp with time zone not null,
    end_time timestamp with time zone not null,
    created_at timestamp with time zone not null default current_timestamp,
    cancelled_at timestamp with time zone
);

create table if not exists booking.event_resource (
    event_id bigserial not null references booking.event(id),
    asset_id int not null,
    primary key (event_id, asset_id)
);

commit;
