-- +migrate Up

alter table lots add column external_id text;
update lots set external_id = details->>'external_id';
create unique index lots_external_id_idx on lots (external_id);

-- +migrate Down

drop index lots_external_id_idx;
alter table lots drop column external_id;