package migrationsdata

const FilesTable = `
CREATE TABLE IF NOT EXISTS files (
	id INTEGER PRIMARY KEY NOT NULL,
	mod_id INTEGER NOT NULL,
	filename TEXT NOT NULL,
	url TEXT NOT NULL,
	FOREIGN KEY (mod_id) REFERENCES mods(id)
);
`
