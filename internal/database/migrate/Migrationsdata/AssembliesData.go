package migrationsdata

const AssembliesTable = `
CREATE TABLE IF NOT EXISTS assemblies (
	id INTEGER PRIMARY KEY NOT NULL,
	loader TEXT NOT NULL,
	creatorid INTEGER NOT NULL,
	assemblyId TEXT NOT NULL,
	name TEXT NOT NULL,
	FOREIGN KEY (creatorid) REFERENCES users (tgId)
);
`
