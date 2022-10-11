package repository

import (
	"rent-book/features/user/domain"

	"gorm.io/gorm"
)

type repoQuery struct {
	db *gorm.DB
}

func New(dbConn *gorm.DB) domain.Repository {
	return &repoQuery{
		db: dbConn,
	}
}

func (rq *repoQuery) Insert(newUser domain.Core) (domain.Core, error) {
	var cnv User
	cnv = FromDomain(newUser)
	if err := rq.db.Create(&cnv).Error; err != nil {
		return domain.Core{}, err
	}
	
	// selesai dari DB
	newUser = ToDomain(cnv)
	return newUser, nil
}
func (rq *repoQuery) Update(updatedData domain.Core) (domain.Core, error) {
	var cnv User
	cnv = FromDomain(updatedData)
	if err := rq.db.Save(&cnv).Error; err != nil {
		return domain.Core{}, err
	}
	// selesai dari DB
	updatedData = ToDomain(cnv)
	return updatedData, nil
}
func (rq *repoQuery) Get(ID uint) (domain.Core, error) {
	var resQry User
	if err := rq.db.First(&resQry, "ID = ?", ID).Error; err != nil {
		return domain.Core{}, err
	}
	// selesai dari DB
	res := ToDomain(resQry)
	return res, nil
}
func (rq *repoQuery) GetAll() ([]domain.Core, error) {
	var resQry []User
	if err := rq.db.Find(&resQry).Error; err != nil {
		return nil, err
	}
	// selesai dari DB
	res := ToDomainArray(resQry)
	return res, nil
}

func (rq *repoQuery) Login(user domain.Core) (domain.Core, error) {
	var dest User
	if err := rq.db.First(&dest, "nama = ? AND password = ?", user.Nama, user.Password).Error ;err != nil {
		return domain.Core{}, err
	}

	res := ToDomain(dest)
	return res, nil
	
}
