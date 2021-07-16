package db

import (
	"encoding/json"
	"fmt"
	"testing"
)

var impl = new(ImplData)

func TestCreateRevenueType(t *testing.T) {
	typeName := []string{"工资", "兼职", "理财", "礼金", "稿费"}
	for _, v := range typeName {
		rt := RevenueType{
			UserId:   0,
			TypeName: v,
		}
		if err := impl.CreateRevenueType(rt); err != nil {
			fmt.Println(err)
		}
	}
}

func TestQueryRevenueTypeList(t *testing.T) {
	rtList, err := impl.QueryRevenueTypeList(5)
	if err != nil {
		fmt.Println(err)
	}

	bt, err := json.Marshal(rtList)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bt))
}
