-- +migrate Up

alter table identities add column type int not null default 4;
create index identities_type_idx on identities (type);

-- +migrate Down

drop index identities_type_idx;
alter table identities drop column type;
