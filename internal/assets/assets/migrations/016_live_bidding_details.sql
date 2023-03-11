-- +migrate Up

alter table lots add column lane text;
alter table lots add column item_num integer;

-- +migrate Down

alter table lots drop column lane;
alter table lots drop column item_num;