-- +migrate Up

alter table participants add column published_creation boolean;




-- +migrate Down

alter table participants drop column published_creation;
