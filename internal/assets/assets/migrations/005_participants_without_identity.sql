-- +migrate Up

alter table participants drop constraint participants_account_address_fkey;

-- +migrate Down

select 1;