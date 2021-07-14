package revenue

import "ledger/api/revenue/db"

var _db db.ImplData

func CreateRevenue(r db.Revenue) error {
	return _db.CreateRevenue(r)
}

func QueryRevenue(userId int64, startTime, endTime string, tType int64, offset, limit int64) (incomeType []*db.Revenue, err error) {
	return _db.QueryRevenue(userId, startTime, endTime, tType, offset, limit)
}

func UpdateRevenue(r db.Revenue) error {
	return _db.UpdateRevenue(r)
}

func CreateRevenueType(rt db.RevenueType) error {

	return _db.CreateRevenueType(rt)
}

func QueryRevenueTypeList(userId int64) (iList []*db.RevenueType, err error) {
	return _db.QueryRevenueTypeList(userId)

}

func DeleteRevenueType(rt db.RevenueType) error {
	return _db.DeleteRevenueType(rt)
}
