package db

type ImplDb struct {
}

var _ DataBase = (*ImplDb)(nil)

type DataBase interface {
	//新增支出记录
	AddConsume(c Consume) error
	//查询支出记录
	QueryConsumeList(startTime, endTime string, userId, bType, limit, offset int64) (rst []*Consume, err error)
	// 更新支出记录
	UpdateConsume(c Consume) error
	// 删除支出记录
	DeleteConsume(c Consume) error
	//   计算月或天的消费总额
	GetConsumeSum(userId int64, month, day string) (consumeMoney float64, err error)
	//创建支出类别
	CreateConsumeType(ct ConsumeType) error
	//支出类别列表
	ConsumeTypeList(userId int64) (ConsumeType []*ConsumeType, err error)
	//删除支出类别
	DeleteConsumeType(ct ConsumeType) error

	//新增预算
	AddBudget(budget Budget) error

	//查询预算
	QueryBudget(userId int64, yearMon string) (bg Budget, err error)
	//月消费支出视图
	GetConsumeView(userId int64, month string) (cView []*ConsumeMonView, err error)
	//日消费支出视图
	GetConsumeDayView(userId int64, day string) (cView []*ConsumeMonViewDay, err error)}
