-- +migrate Up

alter table lots add column sale_id text;
update lots set sale_id = details->>'platform_id' || '-' || (details->>'branch_id') || '-' || lane || '-' || (details->>'start_time') where lane is not null;

create index on lots (sale_id);

-- +migrate Down

alter table lots drop column sale_id;
