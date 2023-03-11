-- +migrate Up

create table movements (
    id bigserial primary key,
    identity text not null,
    amount bigint not null,
    currency text not null,
    action int not null,
    created_at timestamp without time zone not null,
    lot_id text
);

insert into movements (identity, amount, currency, action, created_at) select depositor, amount, currency, 1, created_at from deposits where state in (2,7);
insert into movements (identity, amount, currency, action, created_at, lot_id) select depositor, amount, currency, 5, updated_at, lot_id from deposits where state in (7);

-- +migrate Down

drop table movements;
