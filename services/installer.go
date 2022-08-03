package services

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/apikee/installer/database"
)

func AddAlias() {
	alias := os.Args[2]
	path := os.Args[3:]

	if alias == "" || len(path) == 0 {
		panic("missing arguments")
	}

	resAlias, err := database.CreateAlias(alias)
	if err != nil {
		panic(err)
	}

	for _, p := range path {
		if err := database.CreatePath(p, resAlias.ID); err != nil {
			panic(err)
		}
	}

	fmt.Printf("Record %v created\n", alias)
}

func ListAliases() {
	aliases, err := database.FindAllAliases()
	if err != nil {
		panic(err)
	}

	for _, a := range aliases {
		paths, err := database.FindPathsByAliasID(a.ID)
		if err != nil {
			panic(err)
		}

		for _, p := range paths {
			fmt.Printf("Alias: %v Path: %v\n", a.Alias, p.Path)
		}

		fmt.Println()
	}
}

func InstallDependenciesByAlias() {
	alias := os.Args[2]

	paths, err := database.FindPathsByAlias(alias)
	if err != nil {
		panic(err)
	}

	for _, p := range paths {
		fmt.Printf("Installing %v from alias %v\n", p.Path, alias)
		cmd := exec.Command("go", "get", p.Path)
		stdout, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println(string(stdout))
	}
}

func DeleteAlias() {
	alias := os.Args[2]

	if err := database.DeleteAlias(alias); err != nil {
		panic(err)
	}

	fmt.Printf("Alias %v deleted\n", alias)
}
