package db


type ImplDb struct {
}

type DataBase interface {
	//新增支出记录
	AddConsume(c Consume) error
	//查询支出记录
	QueryByConsume(startTime, endTime string, userId, bType, limit, offset int64) (rst []*Consume, err error)
	// 更新支出记录
	UpdateConsume(c Consume) error
	// 删除支出记录
	DeleteConsume(c Consume) error
	//创建支出类别
	CreateConsumeType(ct ConsumeType) error
	//支出类别列表
	ConsumeTypeList(userId int64) (ConsumeType []*ConsumeType, err error)
	//删除支出类别
	DeleteConsumeType(ct ConsumeType) error
}


