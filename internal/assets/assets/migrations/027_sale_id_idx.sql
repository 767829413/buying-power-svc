-- +migrate Up

drop index lots_sale_id_idx;
create index lots_sale_id_idx on lots (sale_id text_pattern_ops) where sale_id is not null;

-- +migrate Down

select 1;
