package domain

import "github.com/labstack/echo/v4"

type Core struct {
	ID       uint
	Nama     string
	HP       string
	Password string
}

type Repository interface { // Data /Repository (berhubungan dg DB)
	Insert(newUser Core) (Core, error)
	Update(updatedData Core) (Core, error)
	Get(ID uint) (Core, error)
	GetAll() ([]Core, error)
	Login(user Core) (Core, error)
}

type Service interface { // Bisnis logic
	AddUser(newUser Core) (Core, error)
	UpdateProfile(updatedData Core) (Core, error)
	Profile(ID uint) (Core, error)
	ShowAllUser() ([]Core, error)
	Login(user Core) (Core, error)
}

type Handler interface {
	AddUser() echo.HandlerFunc
	ShowAllUser() echo.HandlerFunc
	Login(user Core) (Core, error)
}
