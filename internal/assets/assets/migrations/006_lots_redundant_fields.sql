-- +migrate Up

alter table lots drop column if exists winner_address;

-- +migrate Down

select 1;