package migrationsdata

const ModsTable = `
CREATE TABLE IF NOT EXISTS mods (
	id INTEGER PRIMARY KEY NOT NULL,
	project_id TEXT NOT NULL,
	name TEXT NOT NULL,
	gameversion_id INTEGER NOT NULL,
	FOREIGN KEY (gameversion_id) REFERENCES game_versions(id)
);
`
