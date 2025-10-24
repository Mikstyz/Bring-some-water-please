package service

import (
	"bring_some_water_please/internal/entities"
	"bring_some_water_please/internal/repositories"
	scr "bring_some_water_please/internal/scrper"
	VatchData "bring_some_water_please/utils/VatchData"
	converter "bring_some_water_please/utils/converter"
	"database/sql"
	"log"
)

type ModService struct {
	db *sql.DB
}

func NewModSerivce(db *sql.DB) *ModService {
	return &ModService{db: db}
}

// =========================================
//
//	Наличие мода в бд
//
// =========================================
func IsMod(s *ModService, Name, version, loader string) (*entities.DataMods, error) {
	repoMod := repositories.NewModRepo(s.db)

	isMod, err := repoMod.FindMod(Name, version, loader)
	if err != nil {
		return nil, err
	}

	if isMod == nil {
		log.Println("Мод не найден, загружаю с API")

		allDataInMod, err := scr.GetEntitiesModVersion(Name)
		if err != nil {
			log.Printf("[service][mods] ошибка при получении модов с API: %v", err)
			return nil, err
		}

		found := false
		for _, dataInMod := range allDataInMod {
			modEntitis, err := converter.ConvertExtToEnt(dataInMod)
			if err != nil {
				log.Printf("[service][mod] ошибка при конвертации API модели: %v", err)
				continue
			}

			if err := repoMod.SaveMod(Name, modEntitis); err != nil {
				log.Printf("[service][mod] ошибка при сохранении: %v", err)
				continue
			}

			found = true
			log.Print("[service][mod] мод успешно добавлен в БД")
			VatchData.VatchConvert(modEntitis)
		}

		isMod, err = repoMod.FindMod(Name, version, loader)
		if err != nil {
			return nil, err
		}

		if isMod == nil {
			log.Printf("Мод '%s' нет на версии '%s' с loader '%s'\n", Name, version, loader)
		}

		if !found {
			log.Printf("Мод '%s' вообще не найден в API\n", Name)
		}
	}

	return isMod, nil
}
