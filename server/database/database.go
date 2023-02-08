package database

import (
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

type Snippet struct {
	ID       string `gorm:"primaryKey;type:varchar(255);uniqueIndex"`
	Language string `gorm:"type:varchar(255)"`
	Content  string `gorm:"type:text"`
	Created  int64  `gorm:"type:bigint"`
}

func Connect() (DB, error) {
	db, err := gorm.Open(sqlite.Open(os.Getenv("SNIPPY_DB")), &gorm.Config{})
	if err != nil {
		return DB{}, err
	}

	db.AutoMigrate(&Snippet{})

	return DB{db: db}, nil
}

func (db DB) GetSnippet(id string) (Snippet, error) {
	var snippet Snippet
	result := db.db.First(&snippet, "id = ?", id)

	return snippet, result.Error
}

func (db DB) GetSnippets() ([]Snippet, error) {
	var snippets []Snippet
	result := db.db.Limit(15).Order("created DESC").Find(&snippets)

	return snippets, result.Error
}

func (db DB) UpsertSnippet(snippet Snippet) error {
	snippet.Created = time.Now().Unix()
	result := db.db.Save(&snippet)

	return result.Error
}

func (db DB) DeleteSnippet(id string) error {
	result := db.db.Where("id = ?", id).Delete(&Snippet{})

	return result.Error
}
