-- +migrate Up

create index if not exists deposits_lot_idx on deposits(lot_id);
create index if not exists deposits_not_synced_idx on deposits(id) where not synced;
create index if not exists deposits_state_idx on deposits(state,created_at);
create index if not exists deposits_need_refund_idx on deposits (created_at) where state in (2,9);
create index if not exists deposits_depositor_idx on deposits(depositor);
create index if not exists participants_not_synced on participants (id) where not is_synced;
create index if not exists participants_updated_at_idx on participants (updated_at);

-- +migrate Down

drop index if exists deposits_lot_idx;
drop index if exists deposits_not_synced_idx;
drop index if exists deposits_state_idx;
drop index if exists deposits_need_refund_idx;
drop index if exists deposits_depositor_idx;
drop index if exists participants_not_synced;
