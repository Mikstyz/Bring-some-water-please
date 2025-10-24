package migrationsdata

const AssemblyModsTable = `
CREATE TABLE IF NOT EXISTS assemblymods (
	id INTEGER PRIMARY KEY NOT NULL,
	mod_name INTEGER,
	assembly_id INTEGER,

	FOREIGN KEY (mod_name) REFERENCES mods(name),
	FOREIGN KEY (assembly_id) REFERENCES assemblies(assemblyId)
);
`
