-- +migrate Up

alter table winners
    add column delivered_at timestamp,
    add column published    boolean default true,
    add column version      int     default 0,
    add column is_from_event   bool    default true;

create index on winners (lot_id) where not published;

-- +migrate Down

alter table winners
    drop column delivered_at,
    drop column published,
    drop column version;
