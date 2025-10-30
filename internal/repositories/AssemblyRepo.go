package repositories

import (
	ent "bring_some_water_please/internal/entities"
	unCash "bring_some_water_please/utils/UniqueCash"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type AssemblyRepo struct {
	db *sql.DB
}

func NewAssemblyRepo(db *sql.DB) *AssemblyRepo {
	return &AssemblyRepo{db: db}
}

//=================================================
//    Работа с сборкой
//=================================================

// Проверяет, владеет ли пользователь сборкой
func (r *AssemblyRepo) IsCreatorAssembly(userID int64, assemblyID string) error {
	const query = `
		SELECT name FROM assemblies
		WHERE assemblyid = ? AND creatorid = ?
	`

	var name string
	err := r.db.QueryRow(query, assemblyID, userID).Scan(&name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("пользователь %d не владеет сборкой %s или сборки не существует", userID, assemblyID)
		}
		return fmt.Errorf("ошибка при проверке владельца сборки: %w", err)
	}

	fmt.Print("пользователь явялется владельцем сборки")
	return nil
}

// Получение сборки по Id
func (r *AssemblyRepo) FetchInfoByIdAssembly(AssemblyId string) (*ent.Assemblies, error) {
	var assembly ent.Assemblies

	const query = `
        SELECT loader, name
        FROM assemblies
        WHERE assemblyid = ?
    `

	err := r.db.QueryRow(query, AssemblyId).Scan(
		&assembly.Loader,
		&assembly.Name,
	)
	if err != nil {
		return nil, err
	}

	return &assembly, nil
}

// Сохранение инфы о сборке в бд
func (r *AssemblyRepo) SaveInfoAssembly(assembly ent.Assemblies, userid int64) (Assemblyid string, err error) {

	const Query = `
	INSERT INTO assemblies
	(assemblyId, name, loader)
	VALUES
	(?, ?, ?)
	`

	Assemblyid = unCash.NewCash(userid)

	_, err = r.db.Exec(
		Query,
		Assemblyid,
		assembly.Name,
		assembly.Loader,
	)

	if err != nil {
		return "", fmt.Errorf("ошибка при сохранении сборки в бд: %w", err)
	}

	return Assemblyid, nil
}

// Обновление сборки в бд
func (r *AssemblyRepo) UpdateInfoAssembly(NameAssembly string, assemblyId string, userId int64) error {

	if err := r.IsCreatorAssembly(userId, assemblyId); err != nil {
		return err
	}

	const query = `
        UPDATE assemblies
        SET name = ?
        WHERE assemblyId = ? AND creatorId = ?
    `

	_, err := r.db.Exec(query, NameAssembly, assemblyId, userId)
	if err != nil {
		return fmt.Errorf("ошибка при обновлении сборки: %w", err)
	}

	log.Printf("сборка успешно обновлена")
	return nil
}

//=================================================
//    ПУбличная работа со сборками
//=================================================

//поиск сборок по userid

//=================================================
//    Работа с модами сборки
//=================================================

// Удалить моды из сборки
func (r *AssemblyRepo) RemoveModsByAssembly(mods []string, assemblyId string, userId int64) error {
	// Проверка, владеет ли пользователь сборкой
	if err := r.IsCreatorAssembly(userId, assemblyId); err != nil {
		return err
	}

	// Начинаем транзакцию
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("ошибка начала транзакции: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Удаляем каждый мод по одному
	for _, mod := range mods {
		_, err = tx.Exec(
			`DELETE FROM assemblymods WHERE assembly_id = ? AND mod_name = ?`,
			assemblyId,
			mod,
		)
		if err != nil {
			return fmt.Errorf("ошибка удаления мода %s: %w", mod, err)
		}
	}

	// Коммитим транзакцию
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("ошибка коммита транзакции: %w", err)
	}

	return nil
}

// Добавить моды в сборку
func (r *AssemblyRepo) AddModsByAssembly(mods []string, assemblyId string, userId int64) error {
	// Проверка, владеет ли пользователь сборкой
	if err := r.IsCreatorAssembly(userId, assemblyId); err != nil {
		return err
	}

	// Начинаем транзакцию
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("ошибка начала транзакции: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Вставляем моды по одному
	for _, mod := range mods {
		_, err = tx.Exec(
			`INSERT INTO assemblymods (mod_name, assembly_id) VALUES (?, ?)`,
			mod,
			assemblyId,
		)
		if err != nil {
			return fmt.Errorf("ошибка вставки мода %s: %w", mod, err)
		}
	}

	// Коммитим транзакцию
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("ошибка коммита транзакции: %w", err)
	}

	return nil
}

// Получить все моды в сборке
func (r *AssemblyRepo) FetchModsByAssembly(assemblyId string) ([]string, error) {
	const query = `
		SELECT mod_name
		FROM assemblymods
		WHERE assembly_id = ?
	`

	rows, err := r.db.Query(query, assemblyId)
	if err != nil {
		return nil, fmt.Errorf("ошибка запроса модов: %w", err)
	}
	defer rows.Close()

	var modsName []string
	for rows.Next() {
		var mod string
		if err := rows.Scan(&mod); err != nil {
			return nil, fmt.Errorf("ошибка чтения модов: %w", err)
		}
		modsName = append(modsName, mod)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка итерации модов: %w", err)
	}

	return modsName, nil
}
