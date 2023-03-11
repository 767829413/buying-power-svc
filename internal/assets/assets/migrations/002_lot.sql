-- +migrate Up

create table lots (
    id text not null primary key,
    currency text not null,
    ends_at timestamp without time zone not null,
    state int not null,
    details jsonb not null,
    winner_address text
);

alter table deposits add column lot_id text not null references lots(id);

create table participants
(
    id                bigserial primary key,
    lot_id            text        not null references lots (id),
    account_address   varchar(56) references identities(id),
    state             smallint    not null,

    created_at        timestamp   not null default now(),
    updated_at        timestamp            default null,

    requested_buy_now boolean     not null,
    current_bid       bigint      not null,
    previous_bid      bigint               default null,
    bid_limit         bigint      not null,
    auto_bid_limit    bigint      not null
);

alter table participants
    add constraint participant_lot_account_unique unique (lot_id, account_address);

-- +migrate Down

alter table deposits drop column lot_id;
drop table if exists participants;
drop table if exists lots;