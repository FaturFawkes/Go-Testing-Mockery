package delivery

import "rent-book/features/book/domain"

type AddBookFormat struct {
	Title     string `json:"title" form:"title"`
	Author    string `json:"author" form:"author"`
	Publisher string `json:"publisher" form:"publisher"`
	Year      string `json:"year" form:"year"`
}

type SingleBookFormat struct {
	ID string `json:"id" form:"id" query:"id"`
}

func ToDomain(i interface{}) domain.Core {
	switch i.(type) {
	case AddBookFormat:
		cnv := i.(AddBookFormat)
		return domain.Core{Title: cnv.Title, Author: cnv.Author, Publisher: cnv.Publisher, Year: cnv.Year}
	}
	return domain.Core{}
}
