package db

import "time"

type Budget struct {
	Id        int64     `json:"id" gorm:"primary_key"`
	UserId    int64     `json:"user_id"`
	Money     float64   `json:"money"`
	YearMon   time.Time `json:"year_mon" gorm:"default:null"`
	CreatedAt time.Time `json:"created_at" gorm:"default:null"`
}

func (Budget)TableName()string  {
	return "budget"
}

//func