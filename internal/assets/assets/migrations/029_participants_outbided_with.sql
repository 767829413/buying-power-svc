-- +migrate Up

alter table participants add column outbided_with bigint;

-- +migrate Down

alter table participants drop column outbided_with;
