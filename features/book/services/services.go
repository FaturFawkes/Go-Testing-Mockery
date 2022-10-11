package services

import (
	"errors"
	"rent-book/features/book/domain"
	"strings"
	"github.com/labstack/gommon/log"
)

type bookService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Service {
	return &bookService{
		qry: repo,
	}
}

func (bs *bookService) AddBook(newBook domain.Core) (domain.Core, error) {
	res, err := bs.qry.Insert(newBook)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return domain.Core{}, errors.New("rejected from database")
		}

		return domain.Core{}, errors.New("some problem on database")
	}
	return res, nil
}

func (bs *bookService) ShowAllBook() ([]domain.Core, error) {
	res, err := bs.qry.GetAll()
	if err != nil {
		log.Error("Error Show Book Service", err.Error())
		return nil, err
	}
	return res, nil
}
