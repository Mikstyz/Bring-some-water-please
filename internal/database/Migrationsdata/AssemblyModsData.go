package migrationsdata

const AssemblyModsTable = `
CREATE TABLE IF NOT EXISTS assembly_mods (
	id INTEGER PRIMARY KEY NOT NULL,
	mod_id INTEGER,
	assembly_id INTEGER,
	FOREIGN KEY (mod_id) REFERENCES mods(id),
	FOREIGN KEY (assembly_id) REFERENCES assemblies(id)
);
`
