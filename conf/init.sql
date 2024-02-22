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

create table booking.configuration
(
    id bigserial primary key,
    active boolean default false,
    enable boolean default false,
    project_ids text[],
    user_id text
);

create table booking.event (
    id bigserial primary key,
    config_id bigserial not null references booking.configuration(id),
    title varchar(255) not null,
    description text not null,
    organizer text not null,
    start_time timestamp with time zone not null,
    end_time timestamp with time zone not null,
    created_at timestamp with time zone not null default current_timestamp,
    cancelled_at timestamp with time zone
);

create table booking.event_resource (
    event_id int not null references booking.event(id),
    asset_id int not null,
    primary key (event_id, asset_id)
);

commit;
