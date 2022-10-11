package repository

import (
	"rent-book/features/book/domain"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title     string
	Author    string
	Publisher string
	Year      string
}

func FromDomain(db domain.Core) Book {
	return Book{
		Model:     gorm.Model{ID: db.ID},
		Title:     db.Title,
		Author:    db.Author,
		Publisher: db.Publisher,
		Year:      db.Year,
	}
}

func ToDomain(b Book) domain.Core {
	return domain.Core{
		ID: b.ID,
		Title: b.Title,
		Author: b.Author,
		Publisher: b.Publisher,
		Year: b.Year,
	}
}

func ToDomainArray(ab []Book) []domain.Core {
	var res []domain.Core
	for _, val := range ab {
		res = append(res, domain.Core{ID: val.ID, Title: val.Title,
			Author: val.Author, Publisher: val.Publisher, Year: val.Year,
		})
	}
	return res
}