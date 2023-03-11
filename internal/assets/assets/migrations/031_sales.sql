-- +migrate Up

create table sales (
    id text not null primary key,
    starts_at timestamp without time zone,
    external_id text not null,
    platform text not null
);

create unique index sales_idx on sales (starts_at, external_id, platform);


insert into sales (id, starts_at, platform, external_id) select sale_id, ends_at, split_part(sale_id, '-', 1), split_part(sale_id, '-', 2) || split_part(sale_id, '-', 3) from (select sale_id, ends_at from lots where sale_id is not null group by 1, 2) as data;

update sales set external_id = 'COPART' || external_id where length(external_id) = 4;
update sales set external_id = 'COPART0' || external_id where length(external_id) = 3;
update sales set external_id = 'COPART00' || external_id where length(external_id) = 2;

-- +migrate Down

drop table sales;
drop index sales_idx;
