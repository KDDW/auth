CREATE TABLE
  IF NOT EXISTS realms (
    id SERIAL PRIMARY KEY,
    code VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW (),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW ()
  );

CREATE TABLE
  IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    password VARCHAR(511) NOT NULL,
    realm_id INTEGER NOT NULL,
    CONSTRAINT fk_realms FOREIGN KEY (realm_id) REFERENCES realms (id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW (),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW (),
    CONSTRAINT unique_user UNIQUE (email, realm_id)
  );

CREATE OR REPLACE FUNCTION updated_at_function() 
returns TRIGGER language plpgsql AS $fn$ 
	BEGIN 
		new.updated_at = now();
		RETURN new;
	END; 
$fn$;

DROP TRIGGER IF EXISTS updated_at_trigger ON realms;
DROP TRIGGER IF EXISTS updated_at_trigger ON users;

CREATE TRIGGER updated_at_trigger before
UPDATE ON users FOR each ROW EXECUTE FUNCTION updated_at_function();

CREATE TRIGGER updated_at_trigger before
UPDATE ON realms FOR each ROW EXECUTE FUNCTION updated_at_function();

