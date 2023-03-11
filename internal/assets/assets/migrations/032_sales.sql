-- +migrate Up

delete from sales;


insert into sales (id, starts_at, platform, external_id) select sale_id, ends_at, split_part(sale_id, '-', 1), split_part(sale_id, '-', 2) || split_part(sale_id, '-', 3) from (select sale_id, ends_at from lots where sale_id is not null group by 1, 2) as data;

update sales set external_id = 'COPART' || external_id where length(external_id) = 4;
update sales set external_id = 'COPART0' || external_id where length(external_id) = 3;
update sales set external_id = 'COPART00' || external_id where length(external_id) = 2;

-- +migrate Down

select 1;
