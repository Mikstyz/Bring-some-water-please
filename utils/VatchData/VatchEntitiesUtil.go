package vatchdata

import (
	ent "bring_some_water_please/internal/entities"
	"fmt"
)

func VatchConvert(mod ent.DataMods) {
	fmt.Printf("\n\n======================[ %s ]======================\n", mod.Mods.Name)
	fmt.Printf("ProjectID: %s\n", mod.Mods.ProjectID)
	fmt.Println("---------------------------------------------------")

	fmt.Println("+ MOD INFO:")
	fmt.Printf("  • Name      : %s\n", mod.Mods.Name)
	fmt.Printf("  • ProjectID : %s\n", mod.Mods.ProjectID)
	fmt.Printf("  • Loader    : %s\n", mod.Loader)
	fmt.Printf("  • Version   : %s\n", mod.Version)
	fmt.Println()

	fmt.Println("+ FILE INFO:")
	fmt.Printf("  • FileName  : %s\n", mod.Filename)
	fmt.Printf("  • URL       : %s\n", mod.URL)
	fmt.Println("===================================================")
}
