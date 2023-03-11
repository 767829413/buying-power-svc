-- +migrate Up

create index participant_updated_at_idx on participants (state, updated_at);

-- +migrate Down

drop index participant_updated_at_idx;
