-- +migrate Up

alter table participants add column is_local boolean not null default false;
alter table participants add column is_synced boolean not null default true;
alter table participants add column version integer not null default 0;
update participants set updated_at = created_at where updated_at is null;
ALTER TABLE participants ALTER COLUMN updated_at SET NOT NULL;
alter table participants drop column previous_bid;

CREATE OR REPLACE FUNCTION notify_participant_update() RETURNS trigger AS $$ BEGIN IF (NOT NEW.is_local AND NOT NEW.is_synced) THEN RETURN NULL; END IF; PERFORM pg_notify('participant_updated'); RETURN NULL; END; $$ LANGUAGE plpgsql;

CREATE TRIGGER notify_participant_update AFTER INSERT OR UPDATE ON participants FOR EACH STATEMENT EXECUTE PROCEDURE notify_participant_update();




-- +migrate Down

alter table participants drop column is_local;
alter table participants drop column is_synced;
alter table participants drop column version;
ALTER TABLE participants ALTER COLUMN updated_at drop NOT NULL;
alter table participants add column previous_bid integer;
drop function if exists participant_updated;
drop trigger if exists participants_notify_participant_update on participants;
