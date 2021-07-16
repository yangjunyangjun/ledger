package db

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestAddConsume(t *testing.T) {
	c := Consume{
		UserId: 5,
		Money:  23.36,
		TypeId: 2,
		Remark: "这是个测试啊",
	}
	err := impl.AddConsume(c)
	fmt.Println(err)
}

func TestQueryByConsume(t *testing.T) {
	consumeList, err := impl.QueryByConsume("2021-07-16", "2021-07-17", 0, 2, 1, 0)
	if err != nil {
		fmt.Println(err)
	}
	bt, err := json.Marshal(consumeList)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bt))
}
