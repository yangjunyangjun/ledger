package db

import (
	"ledger/sdk"
	"time"
)

type RevenueType struct {
	Id        int64     `json:"id" gorm:"primary_key"`
	TypeName  string    `json:"type_name"`
	UserId    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at" gorm:"default:null"`
}

func (RevenueType) TableName() string {
	return "income_type"
}

func (*ImplData) CreateRevenueType(rt RevenueType) error {
	if err := sdk.DB.Create(rt).Error; err != nil {
		sdk.Log.Errorf("create revenue type error :%s", err.Error())
		return err
	}
	return nil
}

func (*ImplData) QueryRevenueTypeList(userId int64) (rtList []*RevenueType, err error) {
	if err := sdk.DB.Where("user_id in (0,?)", userId).Find(&rtList).Error; err != nil {
		sdk.Log.Errorf("query revenue type error：%s", err.Error())
		return rtList, err
	}
	return rtList, nil
}

func (*ImplData) DeleteRevenueType(rt RevenueType) error {
	if err := sdk.DB.Where("id = ? and user_id = ?", rt.Id, rt.UserId).Delete(rt).Error; err != nil {
		sdk.Log.Errorf("delete revenue type error:%s", err.Error())
		return err
	}
	return nil
}
