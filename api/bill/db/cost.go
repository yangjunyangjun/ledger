package db

import "ledger/sdk"

type Cost struct {
	cType string `json:"c_type"`
}

func (Cost) TableName() string {
	return "cost"
}

func (*ImplDb) CostList() (cost []*Cost, err error) {
	if err := sdk.DB.Model(cost).Find(&cost).Error; err != nil {
		return nil, err
	}
	return cost, nil
}
