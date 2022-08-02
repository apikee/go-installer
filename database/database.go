package database

import (
	"installer/model"
	"log"
	"os"
	"path/filepath"

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

	DB, err = gorm.Open(sqlite.Open(dir+"/installer.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal(err)
	}

	DB.AutoMigrate(&model.Record{})
}

func CreateRecord(alias string, path string) error {
	err := DB.Create(&model.Record{
		Alias: alias,
		Path:  path,
	}).Error

	return err
}

func FindAllRecords() ([]model.Record, error) {
	var records []model.Record

	err := DB.Find(&records).Error

	return records, err
}

func FindRecordByAlias(alias string) (model.Record, error) {
	var record model.Record

	err := DB.Where("alias = ?", alias).First(&record).Error

	return record, err
}

func DeleteRecordByAlias(alias string) error {
	err := DB.Where("alias = ?", alias).Delete(&model.Record{}).Error

	return err
}
