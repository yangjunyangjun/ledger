package bill

import "ledger/api/bill/db"

var _db db.ImplDb

func AddBill(b db.Bill) error {
	return _db.AddBill(b)
}

func QueryByBill(startTime, endTime string, userId, bType, limit, offset int64) (rst []*db.Bill, err error) {
	return _db.QueryByBill(startTime, endTime, userId, bType, limit, offset)
}

func Update(b db.Bill) error {
	return _db.Update(b)
}

func Delete(b db.Bill) error {
	return _db.Delete(b)
}

func CostList() (cost []*db.Cost, err error) {
	return _db.CostList()
}
