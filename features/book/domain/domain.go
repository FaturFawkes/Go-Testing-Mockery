package domain

import "github.com/labstack/echo/v4"

type Core struct {
	ID        uint
	Title     string
	Author    string
	Publisher string
	Year      string
}

type Repository interface {
	Insert(newBook Core) (Core, error)
	GetAll() ([]Core, error)
}

type Service interface {
	AddBook(newBook Core) (Core, error)
	ShowAllBook() ([]Core, error)
}

type Handler interface {
	AddBook() echo.HandlerFunc
	ShowAllBook() echo.HandlerFunc
}
