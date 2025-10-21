package migrationsdata

const VersionsTable = `
CREATE TABLE IF NOT EXISTS game_versions (
	id INTEGER PRIMARY KEY NOT NULL,
	version TEXT NOT NULL
);
`
