package db

import (
	"github.com/jinzhu/gorm"
	"ledger/sdk"
)

type Bill struct {
	gorm.Model
	UserId int64   `json:"user_id"`
	Amount float64 `json:"amount"`
	Type   int64   `json:"type"`
	Remark string  `json:"remark"`
}

func (Bill) TableName() string {
	return "bill"
}

func (*ImplDb) AddBill(b Bill) error {
	if err := sdk.DB.Create(&b).Error; err != nil {

		return err
	}
	return nil
}

func (*ImplDb) QueryByBill(startTime, endTime string, userId, bType, limit, offset int64) (rst []*Bill, err error) {
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

func (*ImplDb) Update(b Bill) error {
	if err := sdk.DB.Where("id=?", b.ID).Update(b).Error; err != nil {
		return err
	}
	return nil
}

func (*ImplDb) Delete(b Bill) error {
	if err := sdk.DB.Delete(b).Error; err != nil {
		return err
	}
	return nil
}
