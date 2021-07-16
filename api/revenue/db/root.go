package db

type ImplData struct {
}

var _ DataBase = (*ImplData)(nil)

type DataBase interface {
	//新增收入记录
	CreateRevenue(r Revenue) error
	//收入记录列表
	QueryRevenue(userId int64, startTime, endTime string, tType int64, offset, limit int64) (rst []*Revenue, err error)
	//更新收入记录
	UpdateRevenue(r Revenue) error
	//创建收入类别
	CreateRevenueType(rt RevenueType) error
	//收入类别列表
	QueryRevenueTypeList(userId int64) (rtList []*RevenueType, err error)
	//删除收入类别
	DeleteRevenueType(rt RevenueType) error
}
