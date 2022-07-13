package dao

import (
	"database/sql"
	"tdd/model"
)

func GetUsers(db *DBManager, page, pageSize uint32) (users []*model.User, err error) {
	err = db.MySQL.DB.Select(
		&users, "SELECT id, username, password FROM user LIMIT ? OFFSET ?", pageSize, (page-1)*pageSize,
	)
	if err == sql.ErrNoRows {
		users = []*model.User{}
		err = nil
	}

	return
}

func GetUserByID(db *DBManager, userID uint32) (user *model.User, err error) {
	user = &model.User{}
	err = db.MySQL.DB.Get(
		user, "SELECT id, username, password FROM user WHERE id=?", userID,
	)
	if err == sql.ErrNoRows {
		user = nil
		err = nil
	}

	return
}
