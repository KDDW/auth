DROP TABLE IF EXISTS users;

DROP TABLE IF EXISTS realms;

DROP TRIGGER IF EXISTS updated_at_trigger ON users;

DROP TRIGGER IF EXISTS updated_at_trigger ON realms;

DROP TRIGGER IF EXISTS updated_at_function CASCADE;
