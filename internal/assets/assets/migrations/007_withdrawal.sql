-- +migrate Up

create table withdrawals (
    id text not null primary key,
    deposit_id text not null references deposits(id),
    created_at timestamp without time zone not null default now(),
    updated_at timestamp without time zone not null default now(),
    state integer not null,
    synced boolean not null default false
);

-- +migrate Down

drop table withdrawals;