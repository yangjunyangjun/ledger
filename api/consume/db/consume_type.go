package db

import (
	"ledger/sdk"
	"time"
)

type ConsumeType struct {
	Id        int64     `json:"id"`
	TypeName  string    `json:"type_name"`
	UserId    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (ConsumeType) TableName() string {
	return "consume_type"
}

func (*ImplDb) CreateConsumeType(ct ConsumeType) error {
	if err := sdk.DB.Create(ct).Error; err != nil {
		sdk.Log.Errorf("create consume type error :%s", err.Error())
		return err
	}
	return nil
}

func (*ImplDb) ConsumeTypeList(userId int64) (ConsumeType []*ConsumeType, err error) {
	if err := sdk.DB.Where("user_id in (0, ?)", userId).Find(&ConsumeType).Error; err != nil {
		sdk.Log.Errorf("query consume type error：%s", err.Error())
		return nil, err
	}
	return ConsumeType, nil
}

func (*ImplDb) DeleteConsumeType(ct ConsumeType) error {
	if err := sdk.DB.Where("user_id = ? and id = ?", ct.UserId, ct.Id).Delete(ct).Error; err != nil {
		sdk.Log.Errorf("delete ConsumeType error : %s", err.Error())
		return err
	}
	return nil
}
