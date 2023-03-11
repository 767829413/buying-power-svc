-- +migrate Up

alter table withdrawals add column if not exists synced boolean not null default false;

-- +migrate Down

alter table withdrawals drop column synced;