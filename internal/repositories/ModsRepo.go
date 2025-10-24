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

func NewModRepo(db *sql.DB) *AssemblyRepo {
	return &AssemblyRepo{db: db}
}

func (r *AssemblyRepo) FindMod(name, version, loader string) (*ent.DataMods, error) {
	log.Printf("Проверка наличия мода: %v", name)

	var mod ent.DataMods
	const query = `
		SELECT project_id, name, version, loader, filename, url
		FROM mods
		WHERE (? = '' OR name = ?)
		AND (? = '' OR version = ?)
		AND (? = '' OR loader = ?)
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
		log.Printf("Ошибка при поиске мода в базе: %v", err)
		return nil, err
	}

	log.Print("Мод найден в базе")
	return &mod, nil
}

func (r *AssemblyRepo) SaveMod(Name string, mod ent.DataMods) error {
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

func (r *AssemblyRepo) SearchMods(modName, version, loader string) ([]ent.DataMods, error) {
	const query = `
		SELECT project_id, filename, url, version, loader, name, gameversion_id
		FROM mods
		WHERE (? = '' OR name = ?)
		AND (? = '' OR version = ?)
		AND (? = '' OR loader = ?)
	`

	rows, err := r.db.Query(query, modName, modName, version, version, loader, loader)
	if err != nil {
		return nil, fmt.Errorf("ошибка запроса модов: %w", err)
	}
	defer rows.Close()

	var results []ent.DataMods
	for rows.Next() {
		var dm ent.DataMods
		var m ent.Mods
		if err := rows.Scan(
			&dm.ProjectID,
			&dm.Filename,
			&dm.URL,
			&dm.Version,
			&dm.Loader,
			&m.Name,
			&m.GameVersionID,
		); err != nil {
			return nil, fmt.Errorf("ошибка сканирования модов: %w", err)
		}
		dm.Mods = m
		results = append(results, dm)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка итерации модов: %w", err)
	}

	return results, nil
}
