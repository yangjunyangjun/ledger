package db

import (
	"ledger/sdk"
	"time"
)

type Consume struct {
	Id        int64      `json:"id" gorm:"primary_key"`
	UserId    int64      `json:"user_id"`
	Money     float64    `json:"money"`
	TypeId    int64      `json:"type_id"`
	Remark    string     `json:"remark"`
	CreatedAt time.Time  `json:"created_at" gorm:"default:null"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"default:null"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"default:null"`
}

type ConsumeMonView struct {
	Day   string  `json:"day"`
	Money float64 `json:"money"`
}

type ConsumeMonViewDay struct {
	TypeName string  `json:"type_name"`
	Money    float64 `json:"money"`
}

func (Consume) TableName() string {
	return "consume"
}

func (*ImplDb) AddConsume(c Consume) error {
	if err := sdk.DB.Create(&c).Error; err != nil {
		return err
	}
	return nil
}

func (*ImplDb) QueryConsumeList(startTime, endTime string, userId, bType, limit, offset int64) (rst []*Consume, err error) {
	db := sdk.DB
	if startTime != "" {
		db = db.Where("DATE_FORMAT(created_at,'%Y-%m-%d') >= ?", startTime)
	}
	if endTime != "" {
		db = db.Where("DATE_FORMAT(created_at,'%Y-%m-%d') < ?", endTime)
	}
	if userId > 0 {
		db = db.Where("user_id = ?", userId)
	}
	if bType > 0 {
		db = db.Where("type_id = ?", bType)
	}
	if err = db.Order("created_at").Offset(offset).Limit(limit).Find(&rst).Error; err != nil {
		return nil, err
	}
	return rst, nil
}

func (*ImplDb) UpdateConsume(c Consume) error {
	if err := sdk.DB.Where("id=?", c.Id).Update(&c).Error; err != nil {
		return err
	}
	return nil
}

func (*ImplDb) DeleteConsume(c Consume) error {
	if err := sdk.DB.Delete(&c).Error; err != nil {
		return err
	}
	return nil
}

//   计算月或天的消费总额
func (ImplDb) GetConsumeSum(userId int64, month, day string) (consumeMoney float64, err error) {
	db := sdk.DB.Table("consume").Where("user_id = ?", userId)
	if month != "" {
		db = db.Where("DATE_FORMAT(created_at,'%Y-%m') = ?", month)
	} else if day != "" {
		db = db.Where("DATE_FORMAT(created_at,'%Y-%m-%d') = ?", day)
	}
	var moneyList []float64
	if err := db.Pluck("money", &moneyList).Error; err != nil {
		sdk.Log.Error("get consume sum err: ", err)
		return consumeMoney, err
	}
	for _, v := range moneyList {
		consumeMoney += v
	}
	return consumeMoney, nil
}

//月份消费支出视图
func (ImplDb) GetConsumeView(userId int64, month string) (cView []*ConsumeMonView, err error) {
	err = sdk.DB.Table("consume").Select("DATE_FORMAT(created_at,'%Y-%m-%d') as day, SUM(money) as money").
		Where("user_id = ? and DATE_FORMAT(created_at,'%Y-%m') = ?", userId, month).Order("day").
		Group("DATE_FORMAT(created_at,'%Y-%m-%d')").Find(&cView).Error
	return cView, err
}

//天消费支出视图
func (ImplDb) GetConsumeViewDay(userId int64, day string) (cView []*ConsumeMonViewDay, err error) {
	err = sdk.DB.Table("consume").Select("consume_type.type_name as type_name, SUM(consume.money) as money").
		Joins("inner join consume_type on consume.type_id = consume_type.id").
		Where("consume.user_id = ? and DATE_FORMAT(consume.created_at,'%Y-%m-%d') = ?", userId, day).Order("money", true).
		Group("type_name").Find(&cView).Error
	return cView, err
}
