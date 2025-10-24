package scrper

import (
	entApi "bring_some_water_please/internal/externalapi"
	Sutils "bring_some_water_please/utils/stringutils"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const Modrinth string = "https://cdn.modrinth.com/data/"

type TexurePackEnties struct {
}

func GetEntitiesModVersion(ModName string) ([]entApi.Mods, error) {
	fmt.Println("====================================================")
	ModName = Sutils.SpaceToBarsAndLower(ModName)
	url := fmt.Sprintf("https://api.modrinth.com/v2/project/%s/version", ModName)

	log.Printf("url: {%s}", url)

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("%s", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("%s", err)
		return nil, fmt.Errorf("%s", resp.Status)
	}

	var apiMods []entApi.Mods
	if err := json.NewDecoder(resp.Body).Decode(&apiMods); err != nil {
		log.Printf("%s", err.Error())
		return nil, err
	}

	fmt.Println("====================================================")
	return apiMods, nil
}

func GetFileInUrlBytes(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("ошибка при запросе URL %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("не удалось скачать файл %s, статус: %d", url, resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ошибка при чтении данных с URL %s: %w", url, err)
	}

	log.Printf("Файл успешно скачан: %s, размер: %d байт", url, len(data))
	return data, nil
}

func View(ModName string) {
	fmt.Print("====================================================")
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
		log.Printf("Project_id: %s", mod.Project_id)
		for j, loader := range mod.Loaders {
			log.Printf("  Loader %d: %s", j+1, loader)
		}

		for k, file := range mod.Files {
			log.Printf("  File %d: %s | URL: %s", k+1, file.Filename, file.URL)
		}
		log.Printf("---------------------------------------------------------------------------")
	}
	fmt.Print("====================================================")
}
