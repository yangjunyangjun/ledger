package db

import (
	"ledger/sdk"
	"time"
)

type Budget struct {
	Id        int64     `json:"id" gorm:"primary_key"`
	UserId    int64     `json:"user_id"`
	Money     float64   `json:"money"`
	YearMon   string    `json:"year_mon"`
	CreatedAt time.Time `json:"created_at" gorm:"default:null"`
}

func (Budget) TableName() string {
	return "budget"
}

func (ImplDb) AddBudget(budget Budget) error {
	if err := sdk.DB.Create(&budget).Error; err != nil {
		sdk.Log.Error("create budget err:", err)
		return err
	}
	return nil
}

func (ImplDb) QueryBudget(userId int64, yearMon string) (bg Budget, err error) {
	if err := sdk.DB.Where("user_id = ? and year_mon = ?", userId, yearMon).First(&bg).Error; err != nil {
		sdk.Log.Error("query budget err: ", err)
		return bg, err
	}
	return bg, nil
}
