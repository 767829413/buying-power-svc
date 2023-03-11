-- +migrate Up

alter table lots alter column highest_bid type bigint;

-- +migrate Down

select 1;