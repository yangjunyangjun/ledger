package consume

import "ledger/api/consume/db"

var _db db.ImplDb

func AddConsume(c db.Consume) error {
	return _db.AddConsume(c)
}

func QueryConsumeList(startTime, endTime string, userId, bType, limit, offset int64) (rst []*db.Consume, err error) {
	return _db.QueryConsumeList(startTime, endTime, userId, bType, limit, offset)
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

func AddBudget(budget db.Budget) error {
	return _db.AddBudget(budget)
}

func QueryBudget(userId int64, yearMon string) (budgetMoney, consumeMoney float64, err error) {
	budget, err := _db.QueryBudget(userId, yearMon)
	if err != nil {
		return 0, 0, err
	}
	consumeMoney, err = _db.GetConsumeSum(userId, yearMon, "")
	if err != nil {
		return 0, 0, err
	}
	return budget.Money, consumeMoney, nil
}

func GetConsumeView(userId int64, month string) (consume []*db.ConsumeMonView, err error) {
	return _db.GetConsumeView(userId, month)
}
