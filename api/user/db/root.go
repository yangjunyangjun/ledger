package db

type ImplDb struct {
}

var _ DataBase = (*ImplDb)(nil)

type DataBase interface {
	AddUser(user User) error

	GetByUser(email, userName string) (*User, error)

	QueryUserList(id int64, userName, email string, limit, offset int64) (userList []*User, err error)

	Update(id int64, user User) error

	Delete(user User) error
}
