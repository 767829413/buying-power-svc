-- +migrate Up

update lots set external_id = details->>'external_id';

-- +migrate Down

select 1;