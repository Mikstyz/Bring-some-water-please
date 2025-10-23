package migrationsdata

const ModsTable = `
CREATE TABLE IF NOT EXISTS mods (
    id INTEGER PRIMARY KEY NOT NULL,
	project_id TEXT NOT NULL,
	
    name TEXT NOT NULL,
    version TEXT NOT NULL,
    loader TEXT NOT NULL,
    
    filename TEXT NOT NULL,
    url TEXT NOT NULL
);

-- Уникальный индекс на loader + version + name
CREATE UNIQUE INDEX IF NOT EXISTS idx_loader_version_name
ON mods(loader, version, name);
`
