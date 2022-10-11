package repository

import (
	"rent-book/features/book/domain"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type bookQuery struct {
	ud *gorm.DB
}

func New(conn *gorm.DB) domain.Repository {
	return &bookQuery {
		ud: conn,
	}
}

func (bk *bookQuery) Insert(newBook domain.Core) (domain.Core, error) {
	var cnv Book
	cnv = FromDomain(newBook)
	if err := bk.ud.Create(&cnv).Error; err != nil {
		return domain.Core{}, err
	}
	newBook = ToDomain(cnv)
	return newBook, nil
}

func (bk *bookQuery) GetAll() ([]domain.Core, error) {
	var cnv []Book
	if err := bk.ud.Find(&cnv).Error ; err != nil {
		log.Error("error di query")
		return nil, err
	}
	res := ToDomainArray(cnv)
	return res, nil
}