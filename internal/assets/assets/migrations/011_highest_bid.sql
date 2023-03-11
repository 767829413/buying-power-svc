-- +migrate Up

alter table lots add column highest_bid integer;

-- +migrate Down

alter table lots drop column highest_bid;