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

func (Consume) TableName() string {
	return "consume"
}

func (*ImplDb) AddConsume(c Consume) error {
	if err := sdk.DB.Create(&c).Error; err != nil {
		return err
	}
	return nil
}

func (*ImplDb) QueryByConsume(startTime, endTime string, userId, bType, limit, offset int64) (rst []*Consume, err error) {
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
	if err = db.Offset(offset).Limit(limit).Find(&rst).Error; err != nil {
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
	db := sdk.DB.Table("consume").Select("SUM(money)").Where("user_id = ?", userId)
	if month != "" {
		db = db.Where("DATE_FORMAT(created_at,'%Y-%m') = ?", month)
	} else if day != "" {
		db = db.Where("DATE_FORMAT(created_at,'%Y-%m-%d') = ?", day)
	}
	if err := db.Scan(&consumeMoney).Error; err != nil {
		sdk.Log.Error("get consume sum err: ", err)
		return consumeMoney, err
	}
	return consumeMoney, nil
}
