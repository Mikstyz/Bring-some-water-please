package converter

import (
	"fmt"

	ent "bring_some_water_please/internal/entities"
	ext "bring_some_water_please/internal/externalapi"
	StringUtils "bring_some_water_please/utils/stringutils"
)

func ConvertExtToEnt(apiM ext.Mods) (ent.DataMods, error) {
	if len(apiM.Files) == 0 {
		return ent.DataMods{}, fmt.Errorf("мод %s не содержит файлов", apiM.Project_id)
	}

	modName := StringUtils.BeforeFirstBars(apiM.Files[0].Filename)

	// Берём первую комбинацию: первый файл, первый loader, первая версия
	file := apiM.Files[0]
	loader := apiM.Loaders[0]
	version := apiM.GameVersions[0]

	data := ent.DataMods{
		Project_id: apiM.Project_id,
		ProjectID:  apiM.Project_id,
		Filename:   file.Filename,
		URL:        file.URL,
		Version:    version,
		Loader:     loader,
		Mods: ent.Mods{
			ProjectID: apiM.Project_id,
			Name:      modName,
		},
	}

	return data, nil
}

