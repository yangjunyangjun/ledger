package db

import (
	"ledger/sdk"
	"time"
)

type Consume struct {
	Id        int64      `json:"id" gorm:"primary_key"`
	UserId    int64      `json:"user_id"`
	Money     float64    `json:"money"`
	Type      int64      `json:"type"`
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
		db.Where("created_at >= ?", startTime)
	}
	if endTime != "" {
		db.Where("created_at < ?", endTime)
	}
	if userId > 0 {
		db.Where("user_id = ?", userId)
	}
	if bType > 0 {
		db.Where("type = ?", bType)
	}
	if err = db.Offset(offset).Limit(limit).Find(&rst).Error; err != nil {
		return nil, err
	}
	return rst, nil
}

func (*ImplDb) UpdateConsume(c Consume) error {
	if err := sdk.DB.Where("id=?", c.Id).Update(c).Error; err != nil {
		return err
	}
	return nil
}

func (*ImplDb) DeleteConsume(c Consume) error {
	if err := sdk.DB.Delete(c).Error; err != nil {
		return err
	}
	return nil
}
