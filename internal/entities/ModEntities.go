package entities

//----------------------------------
//Сущности которые отвечают за данные о моде
//----------------------------------

//Данные о моде
type Mods struct {
	ProjectID     string `db:"project_id"`
	Name          string `db:"name"`
	GameVersionID int64  `db:"gameversion_id"`
}

//----------------------------------
//Сущности которые отвечают за сборку данных api to sql
//----------------------------------
type DataMods struct {
	Project_id string

	ProjectID string `db:"project_id"`
	Filename  string `db:"filename"`
	URL       string `db:"url"`

	Version string `db:"version"`

	Loader string
	Mods   Mods
}
