-- +migrate Up

CREATE OR REPLACE FUNCTION notify_participant_update() RETURNS trigger AS $$ BEGIN IF (NOT NEW.is_local OR NEW.is_synced) THEN RETURN NULL; END IF; PERFORM pg_notify('participant_updated'::text, NEW.id::text); RETURN NULL; END; $$ LANGUAGE plpgsql;




-- +migrate Down

select 1;
