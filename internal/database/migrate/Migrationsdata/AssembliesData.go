package migrationsdata

const AssembliesTable = `
CREATE TABLE IF NOT EXISTS assemblies (
	id INTEGER PRIMARY KEY NOT NULL,
	creator_id INTEGER,
	name TEXT
);
`
