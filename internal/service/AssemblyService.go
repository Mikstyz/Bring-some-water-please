package service

import (
	ent "bring_some_water_please/internal/entities"
	repo "bring_some_water_please/internal/repositories"
	strUtil "bring_some_water_please/utils/stringutils"
	"database/sql"
	"fmt"
	"log"
	"time"
)

// =========================================

// Для работы с бд
type AssemblyService struct {
	db *sql.DB
}

func NewAssemblyService(db *sql.DB) *AssemblyService {
	return &AssemblyService{db: db}
}

// =========================================

// Для работы с ссылкой
type buildAssembly struct {
	Id     string
	Name   string
	Loader string
	Mods   []string
}

func NewBuildAssembly(userID string, loader string) *buildAssembly {
	id := fmt.Sprintf("%d-%s", time.Now().UnixNano(), userID)

	return &buildAssembly{
		Id:     id,
		Name:   "Name",
		Loader: loader,
		Mods:   []string{},
	}
}

// =========================================

// =========================================
//                 Данные в сборке
// =========================================

// данные о сборке
func (b *buildAssembly) IsBuildAssembly() ent.Assembly {
	return ent.Assembly{
		Id:   b.Id,
		Name: b.Name,
		Mods: b.Mods,
	}
}

// =========================================
//                 Работа со сборкой
// =========================================

// Обновление названия сборки
func (b *buildAssembly) UpdateNameInAssembly(name string) error {

	if strUtil.RemoveSpaceAndLower(name) == "" {
		log.Printf("Name содержит пустую строку: Name = {%s}", name)
		return fmt.Errorf("name является пустым")
	}

	b.Name = name
	return nil
}

// =========================================
//                 	Работа с модами в сборке
// =========================================

// Добавление мода в сборку
func (b *buildAssembly) AddModInAssembly(ModName string) error {
	if strUtil.RemoveSpaceAndLower(ModName) == "" {
		log.Printf("ModName содержит пустую строку: ModName = {%s}", ModName)
		return fmt.Errorf("ModName является пустым")
	}

	//Проверка есть ли данный мод на Loader пользователя

	for _, elem := range b.Mods {
		if elem == ModName {
			return nil
		}
	}

	b.Mods = append(b.Mods, ModName)

	return nil
}

// Удаление мода из сборки
func (b *buildAssembly) DeleteModInAssembly(ModName string) error {
	newMods := []string{}
	for _, name := range b.Mods {
		if name != ModName {
			newMods = append(newMods, name)
		}
	}
	b.Mods = newMods
	return nil
}

// =========================================
//                 	Работа со сборкой в бд
// =========================================

// Поиск сборки в бд
func (s *AssemblyService) FetchInDnAssembly(AssemblyId string) error {
	return nil
}

// Делает zip архив с модами
func (s *AssemblyService) DownloadInDnAssembly(AssemblyId string) ([]byte, error) {
	// Запрос в бд на сборку
	// получение сборки или ошибка при получении
	// ПОТОМ ВЫВОД []byte - zip архив с модами
	return nil, nil
}

// Сохраняет сборку в бд
func (s *AssemblyService) SaveInDnAssembly(b *buildAssembly, userId int64) error {

	var assembly = ent.Assemblies{
		Loader: b.Loader,
		Name:   b.Name,
	}

	//тут должны открываться транзакция
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("ошибка начала транзакции: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	//Сохранение данных о сборке
	asId, err := repo.SaveInfoAssembly(tx, assembly, userId)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("ошибка при выполнении транзакции сохранения информации о инфе собрки")
	}

	//Сохранение модов сборки
	err = repo.AddModsByAssembly(tx, b.Mods, asId, userId)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("ошибка при выполнении транзакции сохранения модов сборки")
	}

	fmt.Printf("new assembly added, Id: %s\nLoader: %s, Name: %s\n", asId, b.Loader, b.Name)

	//Сохраниение модов сборки
	//Функция которая будет строить строку добавляения всех модов в сборку и принимать в себя открытую транзакцию

	return nil
}

// Обновляет сбору в бд
func (s *AssemblyService) UpdateInDnAssembly(newData *buildAssembly, AssemblyId string) error {
	//обновление сборки в бд
	//Просто перезаписываем фулл инфу (главное чтобы не nil)

	//Обновление модов сборки в бд

	// Получение модов
	// Сравниваем новый лист модов и старый, если в новом листе нет старого мода, то мы его удаляем из бд (по транзакции)
	// Сравниваем новый лист модов и старый, если в новом есть моды которых нет в старом, то мы их добавляем в бд (по транзакции)
	return nil
}
