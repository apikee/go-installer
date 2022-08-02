package services

import (
	"fmt"
	"installer/database"
	"os"
	"os/exec"
)

func AddAlias() {
	alias := os.Args[2]
	path := os.Args[3]

	if alias == "" || path == "" {
		panic("missing arguments")
	}

	err := database.CreateRecord(alias, path)
	if err != nil {
		panic(err)
	}

	fmt.Println("Record created")
}

func ListAliases() {
	aliases, err := database.FindAllRecords()
	if err != nil {
		panic(err)
	}

	for _, a := range aliases {
		fmt.Printf("ID: %v, Alias: %v, Path: %v\n", a.ID, a.Alias, a.Path)
	}
}

func InstallDependencyByAlias() {
	alias := os.Args[2]

	record, err := database.FindRecordByAlias(alias)
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("go", "get", record.Path)

	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(stdout))
}

func RemoveRecordByAlias() {
	alias := os.Args[2]

	err := database.DeleteRecordByAlias(alias)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Record %v deleted", alias)
}
