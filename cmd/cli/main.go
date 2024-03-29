package main

import (
	"os"

	"github.com/apikee/go-installer/internal/database"
	"github.com/apikee/go-installer/internal/services"
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
		services.InstallDependenciesByAlias()
	case "delete":
		services.DeleteAlias()
	}
}
