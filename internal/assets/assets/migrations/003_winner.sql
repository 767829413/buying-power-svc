-- +migrate Up

create table winners (
    lot_id text not null primary key references lots(id),
    invoice_id text,
    invoice_created_at timestamp without time zone,
    container_id text,
    container_link text,
    fee int not null,
    fee_currency text not null,
    transportation_price int not null,
    transportation_price_currency text not null,
    state int not null
);

-- +migrate Down

drop table if exists winners;