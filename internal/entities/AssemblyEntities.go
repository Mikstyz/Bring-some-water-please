package entities

//Данные о сборке
type DataAssemblies struct {
	Assemblies   Assemblies
	AssemblyMods AssemblyMods
}

//----------------------------------
//Сущности которые отвечают за данные о Сборке в дб
//----------------------------------
type Assemblies struct {
	Loader string `db:"Loader"`
	Name   string `db:"name"`
}

type AssemblyMods struct {
	ModName    string `db:"mod_name"`
	AssemblyID string `db:"assembly_id"`
}

//----------------------------------
//Сущности которые отвечают за данные о Сборке в проекте
//----------------------------------

type Assembly struct {
	Id   string
	Name string
	Mods []string
}
