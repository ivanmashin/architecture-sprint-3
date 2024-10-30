create table telemetry_data
(
    device_id   varchar      not null,
    timestamp   timestamp    not null,
    state_name  varchar(128) not null,
    state_value jsonb        not null
)
