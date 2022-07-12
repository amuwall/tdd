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
