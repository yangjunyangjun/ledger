package db

import (
	"encoding/json"
	"fmt"
	"testing"
)

var impl = new(ImplDb)

func TestAddConsumeType(t *testing.T) {

	typeList := []string{"烟酒", "医疗", "学习", "礼物", "维修", "娱乐", "汽车", "宠物", "快递", "彩票", "礼金", "居家"}

	for _, v := range typeList {
		c := ConsumeType{
			UserId:   0,
			TypeName: v,
		}
		err := impl.CreateConsumeType(c)
		fmt.Println(err)

	}
}

func TestConsumeTypeList(t *testing.T) {
	typeList, err := impl.ConsumeTypeList(0)
	if err != nil {
		fmt.Println(err)
	}
	bt, err := json.Marshal(typeList)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bt))

}

func TestDeleteConsumeType(t *testing.T) {
	ct := ConsumeType{
		Id:     1,
		UserId: 0,
	}
	err := impl.DeleteConsumeType(ct)
	fmt.Println(err)

}
