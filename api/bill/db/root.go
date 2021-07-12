package db

type ImplDb struct {
}

type DataBase interface {
	AddBill(b Bill) error
	QueryByBill(startTime, endTime string, userId, bType, limit, offset int64) (rst []*Bill, err error)
	Update(b Bill) error
	Delete(b Bill) error
}
