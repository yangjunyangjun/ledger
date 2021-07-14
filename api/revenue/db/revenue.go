package db

import (
	"github.com/jinzhu/gorm"
	"ledger/sdk"
)

type Revenue struct {
	gorm.Model
	UserId int64   `json:"user_id"`
	Money  float64 `json:"Money"`
	Type   int64   `json:"type"`
	Remark string  `json:"remark"`
}

func (Revenue) TableName() string {
	return "revenue"
}

func (*ImplData) CreateRevenue(r Revenue) error {
	if err := sdk.DB.Create(r).Error; err != nil {
		sdk.Log.Errorf("create revenue error: %s", err.Error())
		return err
	}
	return nil
}

func (*ImplData) QueryRevenue(userId int64, startTime, endTime string, tType int64, offset, limit int64) (rst []*Revenue, err error) {
	db := sdk.DB
	if userId > 0 {
		db.Where("user_id = ?", userId)
	}
	if startTime != "" {
		db.Where("created_at >= ?", startTime)
	}
	if endTime != "" {
		db.Where("created_at < ?", endTime)
	}
	if tType > 0 {
		db.Where("type = ?", tType)
	}
	if err := db.Offset(offset).Limit(limit).Find(&rst).Error; err != nil {
		sdk.Log.Errorf("query revenue error: %s", err.Error())
		return rst, err
	}
	return rst, nil
}

func (*ImplData) UpdateRevenue(r Revenue) error {
	if err := sdk.DB.Where("id = ? and user_id = ?", r.ID, r.UserId).Update(&r).Error; err != nil {
		sdk.Log.Errorf("update revenue error: %s", err.Error())
		return err
	}
	return nil
}
