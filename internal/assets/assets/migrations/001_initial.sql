-- +migrate Up

create table identities (
    id text primary key,
    platform text not null
);

create table deposits (
    id text primary key,
    state int not null,
    amount int not null,
    currency text not null,
    depositor text not null references identities(id),
    created_at timestamp without time zone not null default now(),
    updated_at timestamp without time zone not null
);

-- +migrate Down

drop table deposits;
drop table identities;