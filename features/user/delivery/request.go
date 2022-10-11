package delivery

import (
	"rent-book/features/user/domain"
)

type RegisterFormat struct {
	Nama     string `json:"nama" form:"nama"`
	HP       string `json:"hp" form:"hp"`
	Password string `json:"password" form:"password"`
}

type LoginFormat struct {
	Nama     string `json:"nama" form:"nama"`
	Password string `json:"password" form:"password"`
}

func ToDomain(i interface{}) domain.Core {
	switch i.(type) {
	case RegisterFormat:
		cnv := i.(RegisterFormat)
		return domain.Core{Nama: cnv.Nama, HP: cnv.HP, Password: cnv.Password}
	case LoginFormat:
		cnv := i.(LoginFormat)
		return domain.Core{Nama: cnv.Nama, Password: cnv.Password}
	}
	return domain.Core{}
}
