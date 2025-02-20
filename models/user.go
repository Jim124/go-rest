package models

import (
	"errors"
	"go-rest/db"
	"go-rest/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	insertSql := `insert into users(email,password)values(?,?)`
	stmt, error := db.DB.Prepare(insertSql)
	if error != nil {
		return error
	}
	defer stmt.Close()
	hashPassword, error := utils.HashPassword(u.Password)
	if error != nil {
		return error
	}
	result, error := stmt.Exec(u.Email, hashPassword)
	if error != nil {
		return error
	}
	id, error := result.LastInsertId()
	u.ID = id
	return error

}

func (u *User) GetUserByEmail() error {
	querySql := `select id, password from users where email=?`
	row := db.DB.QueryRow(querySql, u.Email)
	var hashPassword string
	error := row.Scan(&u.ID, &hashPassword)
	if error != nil {
		return errors.New("credentials invalid ")
	}
	isValid := utils.HashCompareToPassword(u.Password, hashPassword)
	if !isValid {
		return errors.New("credentials invalid")
	}
	return nil
}
