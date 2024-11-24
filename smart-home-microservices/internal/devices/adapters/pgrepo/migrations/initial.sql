create table homes
(
    id         varchar(255)
        constraint homes_pkey
            primary key,
    name       varchar(255) not null,
    user_id    varchar(255) not null,
    created_at timestamp    not null
);

create index homes_user_id_index
    on homes (user_id);

create table devices
(
    id         varchar(255)
        constraint devices_pkey
            primary key,
    type       varchar(255) not null,
    name       varchar(255) not null,
    online     boolean      not null,
    on_off     boolean      not null,
    user_id    varchar(255) not null,
    home_id    varchar(255)
        constraint devices_homes_id_fk
            references homes
            on delete cascade,
    created_at timestamp    not null
);

create index devices_user_id_index
    on devices (user_id);
create index devices_home_id_index
    on devices (home_id);

create table devices_created_outbox
(
    device_id varchar(255)
        constraint devices_created_outbox_fk
            references devices
            on delete cascade
);

create table devices_updated_outbox
(
    device_id varchar(255)
        constraint devices_updated_outbox_fk
            references devices
            on delete cascade
);

create table devices_deleted_outbox
(
    device_id varchar(255)
        constraint devices_deleted_outbox_fk
            references devices
            on delete cascade
);
