package postgres

import (
	"example/graph/models"
	"github.com/go-pg/pg/v10"
)

type UsersRepo struct {
	DB *pg.DB
}

func (u *UsersRepo) GetUserById(id string) (*models.User, error) {
	var user models.User
	err := u.DB.Model(&user).Where("id = ?", id).First()
	if err != nil {
		return nil, err
	}
	return &user, nil
}
