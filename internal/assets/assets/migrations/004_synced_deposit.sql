-- +migrate Up

alter table deposits add column synced boolean not null default false;

-- +migrate Down

alter table deposits drop column synced;