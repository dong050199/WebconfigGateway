package repoimpl

import (
	models "SQLite_Repo_Pattern/model/user"
	repo "SQLite_Repo_Pattern/repository"
	"database/sql"
	"fmt"
)

type UserRepoImpl struct {
	Db *sql.DB
}

func NewUserRepo(db *sql.DB) repo.UserRepo {
	return &UserRepoImpl{
		Db: db,
	}
}

func (u *UserRepoImpl) FindUserByEmail(email string) (bool, error) {
	row, err := u.Db.Query("SELECT email FROM user WHERE email=?", email)
	if err != nil {
		err := models.ERR_USER_NOT_FOUND
		return false, err
	}
	for row.Next() {
		user := models.User{}
		err := row.Scan(&user.Email)
		if err != nil {
			err := models.ERR_USER_NOT_FOUND
			return false, err
		}
		errr := models.ERR_EMAIL_DUPLICATE
		return true, errr
	}
	errr := models.ERR_USER_NOT_FOUND
	return false, errr
}

func (u *UserRepoImpl) CheckLoginInfo(email string, password string) (models.User, error) {
	user := models.User{}
	row, err := u.Db.Query("SELECT email,password FROM user WHERE email=? AND password=?", email, password)
	if err != nil {
		panic(err)
	}
	for row.Next() {
		user := models.User{}
		err := row.Scan(&user.Email, &user.Password)
		if err != nil {
			panic(err)
		}
	}
	return user, nil
}

func (u *UserRepoImpl) Isert(user models.User) error {
	insertStatement := `
	INSERT INTO user (email, password, displayName)
	VALUES ($1,$2,$3)
	`
	_, err := u.Db.Exec(insertStatement, user.Email, user.Password, user.DisplayName)
	if err != nil {
		return err
	}
	fmt.Println("user added", user)
	return nil
}

func (u *UserRepoImpl) Update(user models.User) error {
	stmt, _ := u.Db.Prepare("UPDATE user set password = ? where email = ?")
	_, err := stmt.Exec(user.Password, user.Email)
	if err != nil {
		return err
	}
	fmt.Println("password updated", user)
	return nil
}

func (u *UserRepoImpl) Delete(email string) error {
	_, err := u.Db.Exec(fmt.Sprintf("DELETE FROM user WHERE email = %s", email))
	if err != nil {
		return err
	}
	fmt.Println("user deleted")
	return nil
}
