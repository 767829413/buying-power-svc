-- +migrate Up

alter table winners add column lot_sold_published boolean default true;

-- +migrate Down

alter table winners drop column lot_sold_published;
