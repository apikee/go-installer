package main

import (
	"os"

	"github.com/apikee/installer/database"
	"github.com/apikee/installer/services"
)

func main() {
	database.New()
	action := os.Args[1]

	switch action {
	case "add":
		services.AddAlias()
	case "list":
		services.ListAliases()
	case "install":
		services.InstallDependencyByAlias()
	case "remove":
		services.RemoveRecordByAlias()
	}
}
