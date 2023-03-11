-- +migrate Up

drop index lots_external_id_idx;
create index lots_external_id_idx on lots (external_id);

-- +migrate Down

drop index lots_external_id_idx;
create unique index lots_external_id_idx on lots (external_id);