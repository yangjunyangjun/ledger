package consume

import "ledger/api/consume/db"

var _db db.ImplDb

func AddConsume(c db.Consume) error {
	return _db.AddConsume(c)
}

func QueryByConsume(startTime, endTime string, userId, bType, limit, offset int64) (rst []*db.Consume, err error) {
	return _db.QueryByConsume(startTime, endTime, userId, bType, limit, offset)
}

func UpdateConsume(c db.Consume) error {
	return _db.UpdateConsume(c)
}

func DeleteConsume(c db.Consume) error {
	return _db.DeleteConsume(c)
}

func CreateConsumeType(ct db.ConsumeType) error {
	return _db.CreateConsumeType(ct)
}

func ConsumeTypeList(userId int64) (ConsumeType []*db.ConsumeType, err error) {
	return _db.ConsumeTypeList(userId)
}

func DeleteConsumeType(ct db.ConsumeType) error {
	return _db.DeleteConsumeType(ct)
}
