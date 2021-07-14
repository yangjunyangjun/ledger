package db

import (
	"github.com/jinzhu/gorm"
	"ledger/sdk"
)

type Income struct {
	gorm.Model
	UserId int64   `json:"user_id"`
	Money  float64 `json:"Money"`
	Type   int64   `json:"type"`
	Remark string  `json:"remark"`
}

type IncomeType struct {
	Id       int64 `json:"id"`
	TypeName int64 `json:"type_name"`
}

func (*implData) CreateIncome(income Income) error {
	if err := sdk.DB.Create(income).Error; err != nil {
		sdk.Log.Errorf("create income error: %s", err.Error())
		return err
	}
	return nil
}

func (*implData) QueryIncome(userId int64, startTime, endTime string, tType int64) {
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
}

func (*implData) UpdateIncome(income Income) error {
	if err := sdk.DB.Where("id = ?", income.ID).Update(&income).Error; err != nil {
		sdk.Log.Errorf("update income error: %s", err.Error())
		return err
	}
	return nil
}

func (*implData) CreateIncomeType(incomeType IncomeType) error {
	if err := sdk.DB.Create(incomeType).Error; err != nil {
		sdk.Log.Errorf("create income type error :%s",err.Error())
		return err
	}
	return nil
}


