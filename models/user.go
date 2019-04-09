package models

import (
	"comma-rpc-services/dao"
	sq "github.com/Masterminds/squirrel"
	"comma-rpc-services/thrift/schedule"
)

func GetUser(id int64) (*schedule.User, error) {
	u := schedule.User{}

	err := dao.Db.Select("id, username, phone").
		From("t_user").
		Where(sq.Eq{"user_status": VALID, "id": id}).
		QueryRow().
		Scan(&u.ID, &u.UserName, &u.Phone)

	return &u, err
}