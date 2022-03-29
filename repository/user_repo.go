package repository

import (
	models "SQLite_Repo_Pattern/model/user"
)

type UserRepo interface {
	FindUserByEmail(email string) (bool, error)
	CheckLoginInfo(email string, password string) (models.User, error)
	Update(user models.User) error
	Delete(email string) error
	Isert(user models.User) error
}
