package services

import (
	"rent-book/features/user/domain"
	"rent-book/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShowAll(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sukses Get All", func(t *testing.T) {
		repo.On("GetAll").Return([]domain.Core{{ID: uint(1),Nama: "Fatur", HP: "08123", Password: "rohman"}}, nil).Once()
		srv := New(repo)
		res, err := srv.ShowAllUser()
		obj := res[0]
		assert.Nil(t, err)
		assert.NotEmpty(t, obj.ID, "seharusnya ada id yang dikembalikan")
		assert.NotEmpty(t, obj.HP, "seharusnya ada hp yang dikembalikan")
		assert.NotEmpty(t, obj.ID, "seharusnya ada pw yang dikembalikan")
		assert.NotNil(t, res, "tidak ada ID")
		repo.AssertExpectations(t)
	})
}
