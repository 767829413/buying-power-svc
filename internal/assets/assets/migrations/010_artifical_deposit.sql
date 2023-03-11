-- +migrate Up

alter table deposits add column is_artificial boolean not null default false;

-- +migrate Down

alter table deposits drop column is_artificial;