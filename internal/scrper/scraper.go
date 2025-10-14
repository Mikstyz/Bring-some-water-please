package scrper

import (
	ent "bring_some_water_please/internal/entities"
	utils "bring_some_water_please/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const Modrinth string = "https://cdn.modrinth.com/data/"

type TexurePackEnties struct {
}

func GetEntitiesModVersion(ModName string) ([]ent.Mods, error) {
	ModName = utils.SpaceToBars(ModName)
	url := fmt.Sprintf("https://api.modrinth.com/v2/project/%s/version", ModName)

	log.Printf("url: {%s}", url)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Ошибка при запросе: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Неудачный ответ от API: %v", resp.Status)
	}

	// Массив версий из API
	var apiMods []struct {
		ID           string   `json:"id"`
		Name         string   `json:"name"`
		GameVersions []string `json:"game_versions"`
		Loaders      []string `json:"loaders"`
		Files        []struct {
			Filename string `json:"filename"`
			URL      string `json:"url"`
		} `json:"files"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&apiMods); err != nil {
		log.Fatalf("Ошибка при парсинге JSON: %v", err)
	}

	var mods []ent.Mods

	for _, m := range apiMods {
		mod := ent.Mods{
			ID:           m.ID,
			Name:         m.Name,
			GameVersions: m.GameVersions,
			Loaders:      m.Loaders,
			Files:        []ent.Files{},
		}

		for _, f := range m.Files {
			mod.Files = append(mod.Files, ent.Files{
				Filename: f.Filename,
				URL:      f.URL,
			})
		}

		mods = append(mods, mod)
	}

	return mods, nil
}

func View(ModName string) {

	mods, error := GetEntitiesModVersion(ModName)

	if error != nil {
		log.Fatal("Ошибка api")
	}

	// Вывод построчно
	for i, mod := range mods {
		log.Printf("---------------------------------------------------------------------------")
		log.Printf("Version %d:", i+1)
		log.Printf("ID: %s", mod.ID)
		log.Printf("Name: %s", mod.Name)

		for j, loader := range mod.Loaders {
			log.Printf("  Loader %d: %s", j+1, loader)
		}

		for k, file := range mod.Files {
			log.Printf("  File %d: %s | URL: %s", k+1, file.Filename, file.URL)
		}
	}
	log.Printf("---------------------------------------------------------------------------")
}
