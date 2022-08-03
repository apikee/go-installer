package database

import (
	"log"
	"os"
	"path/filepath"

	"github.com/apikee/installer/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func New() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	if os.Getenv("NODE_ENV") == "dev" {
		DB, err = gorm.Open(sqlite.Open("./installer.db"), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
	} else {
		DB, err = gorm.Open(sqlite.Open(dir+"/installer.db"), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
	}
	if err != nil {
		log.Fatal(err)
	}

	if err = DB.AutoMigrate(&model.Alias{}, &model.Path{}); err != nil {
		panic(err)
	}
}

func CreateAlias(alias string) (model.Alias, error) {
	resAlias := model.Alias{Alias: alias}

	if err := DB.Create(&resAlias).Error; err != nil {
		return model.Alias{}, err
	}

	return resAlias, nil
}

func CreatePath(path string, aliasID uint) error {
	return DB.Create(&model.Path{Path: path, Alias: aliasID}).Error
}

func FindAllAliases() ([]model.Alias, error) {
	var aliases []model.Alias

	if err := DB.Find(&aliases).Error; err != nil {
		return []model.Alias{}, err
	}

	return aliases, nil
}

func FindPathsByAliasID(aliasID uint) ([]model.Path, error) {
	var paths []model.Path

	if err := DB.Where("alias = ?", aliasID).Find(&paths).Error; err != nil {
		return []model.Path{}, err
	}

	return paths, nil
}

func FindPathsByAlias(alias string) ([]model.Path, error) {
	var a model.Alias

	if err := DB.Where("alias = ?", alias).First(&a).Error; err != nil {
		return []model.Path{}, err
	}

	var paths []model.Path

	if err := DB.Where("alias = ?", a.ID).Find(&paths).Error; err != nil {
		return []model.Path{}, err
	}

	return paths, nil
}

func DeleteAlias(alias string) error {
	var a model.Alias

	if err := DB.Where("alias = ?", alias).First(&a).Error; err != nil {
		return err
	}

	if err := DB.Unscoped().Where("alias = ?", a.ID).Delete(&model.Path{}).Error; err != nil {
		return err
	}

	if err := DB.Unscoped().Delete(&model.Alias{}, a.ID).Error; err != nil {
		return err
	}

	return nil
}
