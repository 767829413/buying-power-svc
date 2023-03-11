-- +migrate Up

drop index lots_external_id_idx;
create index lots_external_id_idx on lots (external_id text_pattern_ops);

-- +migrate Down

select 1;
