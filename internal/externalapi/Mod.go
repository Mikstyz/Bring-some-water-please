package externalapi

type Mods struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	GameVersions []string `json:"game_versions"`
	Loaders      []string `json:"loaders"`
	Project_id   string   `json:"project_id"`
	Files        []Files  `json:"files"`
}

type Files struct {
	Filename string `json:"filename"`
	URL      string `json:"url"`
}
