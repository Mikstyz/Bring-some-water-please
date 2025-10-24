package service

import (
	"bring_some_water_please/internal/entities"
	scr "bring_some_water_please/internal/scrper"
	VatchData "bring_some_water_please/utils/VatchData"
	csFolder "bring_some_water_please/utils/folderCash"
	strUtil "bring_some_water_please/utils/stringutils"
	"database/sql"
	"log"
)

type DownloadModService struct {
	db *sql.DB
}

func NewDownloadModSerivce(db *sql.DB) *DownloadModService {
	return &DownloadModService{db: db}
}

func (s *DownloadModService) DownloadMod(Name string, Version string, Loader string) ([]byte, error) {
	r := NewModSerivce(s.db)

	ismod, err := IsMod(r, Name, Version, Loader)
	if err != nil {
		return nil, err
	}

	if ismod == nil {
		log.Printf("Мод '%s' не найден, пропускаем обработку\n", Name)
		return nil, nil
	}

	log.Println(" + Мод найден")

	VatchData.VatchConvert(*ismod)

	//=============================================================
	//======================Скачивание мода========================
	//=============================================================

	byteFile, err := scr.GetFileInUrlBytes(ismod.URL)

	if err != nil {
		return nil, err
	}

	return byteFile, nil
}

func (s *DownloadModService) DownloadMods(mods []entities.ModFile, userId int64) (DownloadPath string, err error) {
	DownloadPath, err = csFolder.NewCashFolder(userId)
	if err != nil {
		return "", err
	}

	for _, mod := range mods {
		data, err := s.DownloadMod(mod.Name, mod.Version, mod.Loader)

		if err != nil {
			return "", err
		}

		err = csFolder.AddFileInCashFolder(data, DownloadPath, (strUtil.SpaceToBarsAndLower(mod.Name) + "_" + mod.Loader + "_" + mod.Version))

		if err != nil {
			return "", nil
		}
	}

	return DownloadPath, nil
}
