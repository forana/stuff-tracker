package model

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"time"
)

func id(s string) uuid.UUID {
	u, _ := uuid.Parse(s)
	return u
}

func Seed(db *gorm.DB) {
	category := &Category{
		ID:   id("0dd33b12-7e6a-4590-ac59-39c4d6a01d7b"),
		Name: "Cats",
		Parent: &Category{
			ID:     id("83505003-e0b6-4e55-8bb2-43e532caddaf"),
			Name:   "Animals",
			Parent: nil,
		},
	}
	db.Create(category)
	db.Create(&Item{
		ID:       id("e1086555-76ca-4801-b865-27ccdb80110e"),
		Name:     "Wilbur",
		Category: category,
		AddedAt:  time.Now(),
	})
	db.Create(&Item{
		ID:       id("4314c8ec-f946-4983-87da-58468a87a9eb"),
		Name:     "Thidwick",
		Category: category,
		AddedAt:  time.Now(),
	})
	db.Create(&Item{
		ID:       id("2c636d06-eac9-45f4-952b-240d2fc53d54"),
		Name:     "Ankle",
		Category: category,
		AddedAt:  time.Now(),
	})
}
