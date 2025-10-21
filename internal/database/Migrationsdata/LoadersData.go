package migrationsdata

const LoadersTable = `
CREATE TABLE IF NOT EXISTS loaders (
	id INTEGER PRIMARY KEY NOT NULL,
	loader TEXT NOT NULL
);
`

const DefaultLoadersInsert = `
INSERT INTO loaders (loader)
SELECT 'fabric'
WHERE NOT EXISTS (SELECT 1 FROM loaders WHERE loader = 'fabric');

INSERT INTO loaders (loader)
SELECT 'forge'
WHERE NOT EXISTS (SELECT 1 FROM loaders WHERE loader = 'forge');
`
