package scrper

import (
	"Bring-some-water-please/utils"
	"fmt"
)

const Modrinth string = "https://cdn.modrinth.com/data/"

type dataMod struct {
	ModId   string `json:"id"`
	Version string `json:"versions"`
}

type ModEnties struct {
	ID   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	data dataMod
}

type TexurePackEnties struct {
}

func GetModE(ModName string) {
	ModName = utils.SpaceToBars(ModName)

	fmt.Print(ModName)
}
