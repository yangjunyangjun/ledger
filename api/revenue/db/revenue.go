package db

import (
	"ledger/sdk"
	"time"
)

type Revenue struct {
	Id        int64      `json:"id" gorm:"primary_key"`
	UserId    int64      `json:"user_id"`
	Money     float64    `json:"Money"`
	Type      int64      `json:"type"`
	Remark    string     `json:"remark"`
	CreatedAt time.Time  `json:"created_at" gorm:"default:null"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"default:null"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"default:null"`
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
	if err := sdk.DB.Where("id = ? and user_id = ?", r.Id, r.UserId).Update(&r).Error; err != nil {
		sdk.Log.Errorf("update revenue error: %s", err.Error())
		return err
	}
	return nil
}
