-- +migrate Up

alter table participants add column requested_bid integer not null default 0;

-- +migrate Down

alter table participants drop column requested_bid;
