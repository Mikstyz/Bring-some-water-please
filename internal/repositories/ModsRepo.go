package repositories

import (
	ent "bring_some_water_please/internal/entities"
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

type ModRepo struct {
	db *sql.DB
}

func NewModRepo(db *sql.DB) *ModRepo {
	return &ModRepo{db: db}
}

func (r *ModRepo) FindMod(name, version, loader string) (*ent.DataMods, error) {
	log.Printf("Проверка наличия мода: %v", name)

	var mod ent.DataMods
	query := `
    SELECT project_id, name, version, loader, filename, url
    FROM mods
    WHERE (name = ? OR ? IS NULL)
    AND (version = ? OR ? IS NULL)
    AND (loader = ? OR ? IS NULL)
    LIMIT 1
	`

	err := r.db.QueryRow(query,
		name, name,
		version, version,
		loader, loader,
	).Scan(
		&mod.Mods.ProjectID, // project_id
		&mod.Mods.Name,      // name
		&mod.Version,        // version
		&mod.Loader,         // loader
		&mod.Filename,       // filename
		&mod.URL,            // url
	)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Print("Мод не найден в базе")
			return nil, nil
		}
		log.Print("Ошибка при поиске мода в базе")
		return nil, err
	}

	log.Print("Мод найден в базе")
	return &mod, nil
}

func (r *ModRepo) SaveMod(Name string, mod ent.DataMods) error {
	const Query = `
	INSERT OR IGNORE INTO mods
	(name, version, loader, project_id, filename, url)
	VALUES (?, ?, ?, ?, ?, ?)
	`

	_, err := r.db.Exec(
		Query,
		Name,
		mod.Version,
		mod.Loader,
		mod.ProjectID,
		mod.Filename,
		mod.URL,
	)

	if err != nil {
		return fmt.Errorf("ошибка при сохранении мода в бд: %w", err)
	}

	return nil
}

func (r *ModRepo) LoadMod(modName string, version string, loader string) error {

	const Query string = ""
	return nil
}
