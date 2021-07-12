package user

import "ledger/api/user/db"

var _db *db.ImplDb

func AddUser(user db.User) error {
	return _db.AddUser(user)
}

func GetByUser(email, userName string) (*db.User, error) {
	return _db.GetByUser(email, userName)
}

func QueryUserList(id int64, userName, email string, limit, offset int64) (userList []*db.User, err error) {
	return _db.QueryUserList(id, userName, email, limit, offset)
}
