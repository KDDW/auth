CREATE TABLE IF NOT EXISTS realms (
	id SERIAL PRIMARY KEY,
	code VARCHAR(255) NOT NULL UNIQUE,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	email VARCHAR(255) NOT NULL,
	password VARCHAR(511) NOT NULL,
	realm_id INTEGER NOT NULL,
	CONSTRAINT fk_realms FOREIGN KEY (realm_id) REFERENCES realms(id),
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);