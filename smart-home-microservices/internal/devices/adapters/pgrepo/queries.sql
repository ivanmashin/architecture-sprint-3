-- name: checkHomeBelongsToUser :one
select exists(select * from homes where user_id = $1 and id = $2);

-- name: ListUserHomes :many
select *
from homes
where user_id = $1;

-- name: ListHomeDevices :many
select *
from devices
where home_id = $1
  and user_id = $2;

-- name: GetDeviceByID :one
select *
from devices
where id = $1
  and user_id = $2;

-- name: CreateDevice :exec
insert into devices (id, type, name, online, on_off, user_id, home_id)
values ($1, $2, $3, $4, $5, $6, $7);

-- name: UpdateDevice :exec
update devices
set name   = $1,
    online = $2,
    on_off = $3
where id = $4
  and user_id = $5;

-- name: DeleteDeviceByID :exec
delete
from devices
where id = $1
  and user_id = $2;

-- name: GetDeviceByIDForUpdate :one
select *
from devices
where id = $1
  and user_id = $2
    for update;

-- name: SaveDeviceCreatedToOutbox :exec
insert into devices_created_outbox (device_id)
values ($1);

-- name: SaveDeviceUpdatedToOutbox :exec
insert into devices_updated_outbox (device_id)
values ($1);

-- name: SaveDeviceDeletedToOutbox :exec
insert into devices_deleted_outbox (device_id)
values ($1);

-- name: GetOutboxMessagesDeviceCreated :many
select *
from devices_created_outbox;

-- name: GetOutboxMessagesDeviceUpdated :many
select *
from devices_updated_outbox;

-- name: GetOutboxMessagesDeviceDeleted :many
select *
from devices_deleted_outbox;

-- name: DeleteOutboxMessagesDeviceCreated :exec
delete
from devices_created_outbox;

-- name: DeleteOutboxMessagesDeviceUpdated :exec
delete
from devices_updated_outbox;

-- name: DeleteOutboxMessagesDeviceDeleted :exec
delete
from devices_deleted_outbox;
