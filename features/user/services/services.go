package services

import (
	"errors"
	"rent-book/features/user/domain"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Service {
	return &userService{
		qry: repo,
	}
}

func GenerateToken(id uint) string {
	claim := &jwt.MapClaims{
		"authorized": true,
		"id":         id,
		"exp":        time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	str, err := token.SignedString([]byte("rahasia"))
	if err != nil {
		log.Error(err.Error())
		return ""
	}
	return str
}

func (us *userService) Login(user domain.Core) (domain.Core, error) {
	res, err := us.qry.Login(user)
	if err != nil {
		log.Error(err.Error())
		return domain.Core{}, errors.New("login failed")
	}

	generate, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	// pw = bcrypt.
	if res.Nama != user.Nama && res.Password != string(generate) {
		log.Error(err.Error())
		return domain.Core{}, errors.New("username atau password salah")
	}
	return res, nil
}

func (us *userService) AddUser(newUser domain.Core) (domain.Core, error) {
	generate, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Error(err.Error())
		return domain.Core{}, errors.New("cannot encript password")
	}

	newUser.Password = string(generate)
	res, err := us.qry.Insert(newUser)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return domain.Core{}, errors.New("rejected from database")
		}

		return domain.Core{}, errors.New("some problem on database")
	}

	return res, nil
}

func (us *userService) UpdateProfile(updatedData domain.Core) (domain.Core, error) {
	if updatedData.Password != "" {
		generate, err := bcrypt.GenerateFromPassword([]byte(updatedData.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Error(err.Error())
			return domain.Core{}, errors.New("cannot encript password")
		}
		updatedData.Password = string(generate)
	}

	res, err := us.qry.Update(updatedData)
	if err != nil {
		if strings.Contains(err.Error(), "column") {
			return domain.Core{}, errors.New("rejected from database")
		}
	}

	return res, nil
}

func (us *userService) Profile(ID uint) (domain.Core, error) {
	res, err := us.qry.Get(ID)
	if err != nil {
		log.Error(err.Error())
		if strings.Contains(err.Error(), "table") {
			return domain.Core{}, errors.New("database error")
		} else if strings.Contains(err.Error(), "found") {
			return domain.Core{}, errors.New("no data")
		}
	}

	return res, nil
}

func (us *userService) ShowAllUser() ([]domain.Core, error) {
	res, err := us.qry.GetAll()
	if err != nil {
		log.Error(err.Error())
		if strings.Contains(err.Error(), "table") {
			return nil, errors.New("database error")
		} else if strings.Contains(err.Error(), "found") {
			return nil, errors.New("no data")
		}
	}

	if len(res) == 0 {
		log.Info("no data")
		return nil, errors.New("no data")
	}

	return res, nil
}
