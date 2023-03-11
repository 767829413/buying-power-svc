-- +migrate Up

alter table identities add column is_deleted boolean not null default false;

-- +migrate Down

alter table identities drop column is_deleted;
