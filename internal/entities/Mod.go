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

//----------------------------------
//Сущности которые отвечают за данные о сборке в sql
//----------------------------------

//Данные о сборке
type DataAssemblies struct {
	Assemblies   Assemblies
	AssemblyMods AssemblyMods
}

//----------------------------------
//Сущности которые отвечают за данные о Сборке
//----------------------------------
type Assemblies struct {
	CreatorID int64  `db:"creator_id"`
	Name      string `db:"name"`
}

type AssemblyMods struct {
	ModID      int64 `db:"mod_id"`
	AssemblyID int64 `db:"assembly_id"`
}
