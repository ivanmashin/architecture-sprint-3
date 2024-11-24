-- name: DeleteDeviceStates :exec
delete
from telemetry_data
where device_id = $1;

-- name: InsertDeviceState :exec
insert into telemetry_data (device_id, timestamp, state_name, state_value)
values (unnest(@devices_ids::text[]),
        unnest(@timestamps::timestamp[]),
        unnest(@state_names::text[]),
        unnest(@state_values::bytea[]));

-- name: GetDeviceTelemetry :many
select *
from telemetry_data
where device_id = $1
order by timestamp desc;
