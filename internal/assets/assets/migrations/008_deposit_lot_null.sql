-- +migrate Up

alter table deposits alter column lot_id drop not null;

-- +migrate Down

select 1;