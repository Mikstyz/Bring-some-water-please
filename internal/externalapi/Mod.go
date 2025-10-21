package entities

type Mods struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	GameVersions []string `json:"game_versions"`
	Loaders      []string `json:"loaders"`
	Files        []Files  `json:"files"`
}

type Files struct {
	Filename string `json:"filename"`
	URL      string `json:"url"`
}

var ApiMods []struct {
	Name         string   `json:"name"`
	ID           string   `json:"id"`
	GameVersions []string `json:"game_versions"`
	Loaders      []string `json:"loaders"`
	Files        []struct {
		Filename string `json:"filename"`
		URL      string `json:"url"`
	} `json:"files"`
}
