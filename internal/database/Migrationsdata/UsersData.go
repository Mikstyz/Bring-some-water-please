package migrationsdata

const UsersTable = `
CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY NOT NULL,
	tgId INTEGER NOT NULL
);
`
