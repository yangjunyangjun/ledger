package db

import (
	"github.com/jinzhu/gorm"
	"ledger/sdk"
	"time"
)

type User struct {
	Id        int64      `json:"id" gorm:"primary_key"`
	UserName  string     `json:"user_name"`                              //用户昵称
	Password  string     `json:"password"`                               // 密码
	Icon      string     `json:"icon"`                                   // 图像
	Role      int64      `json:"role" gorm:"type:enum(1,2,3);default:1"` //用户角色
	Email     string     `json:"email"`                                  //邮箱
	CreatedAt time.Time  `json:"created_at" gorm:"default:null"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"default:null"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"default:null"`
}

// 指定表名
func (User) TableName() string {
	return "user"
}

// 新增用户
func (*ImplDb) AddUser(user User) error {
	if err := sdk.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

//根据邮箱或用户名获取user
func (*ImplDb) GetByUser(email, userName string) (*User, error) {
	user := new(User)
	if err := sdk.DB.Where("email = ? or user_name=?", email, userName).First(&user).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return user, nil
}

//用户列表
func (*ImplDb) QueryUserList(id int64, userName, email string, limit, offset int64) (userList []*User, err error) {
	db := sdk.DB
	if id > 0 {
		db.Where("id = ?", id)
	}
	if userName != "" {
		db.Where("user_name = ?", userName)
	}
	if email != "" {
		db.Where("email = ?", email)
	}
	if err := db.Limit(limit).Offset(offset).Find(&userList).Error; err != nil {
		return userList, err
	}
	return userList, nil
}
func (*ImplDb) Update(id int64, user User) error {
	if err := sdk.DB.Where("id = ?", id).Update(&user).Error; err != nil {
		return err
	}
	return nil
}

func (*ImplDb) Delete(user User) error {
	if err := sdk.DB.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
