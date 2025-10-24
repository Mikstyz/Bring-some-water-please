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
	repoAssembly := repo.NewAssemblyRepo(s.db)

	var assembly = ent.Assemblies{
		Loader: b.Loader,
		Name:   b.Name,
	}

	asId, err := repoAssembly.SaveInfoAssembly(assembly, userId)

	if err != nil {
		log.Printf("ошибка при сохранении сборки: ERROR: %w \n", err)
		return err
	}

	fmt.Printf("new assembly added, Id: %s\nLoader: %s, Name: %s\n", asId, b.Loader, b.Name)

	return nil
}

// Обновляет сбору в бд
func (s *AssemblyService) UpdateInDnAssembly(newData *buildAssembly, AssemblyId string) error {
	//обновление сборки в бд
	return nil
}
