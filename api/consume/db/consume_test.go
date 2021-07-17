package db

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestAddConsume(t *testing.T) {
	for i := 0; i < 30; i++ {
		rand.Seed(time.Now().UnixNano())
		typeId := rand.Intn(32)
		Money:=rand.Float64()*100
		c := Consume{
			UserId: 5,
			Money:  Money,
			TypeId: int64(typeId),
			Remark: "这是个测试啊",
		}
		err := impl.AddConsume(c)
		fmt.Println(err)
	}
}

func TestQueryByConsume(t *testing.T) {
	consumeList, err := impl.QueryConsumeList("2021-07-16", "2021-07-17", 0, 2, 1, 0)
	if err != nil {
		fmt.Println(err)
	}
	bt, err := json.Marshal(consumeList)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bt))
}

func TestConsumeView(t *testing.T) {
	c,err:=impl.GetConsumeView(5,"2021-07")
	if err != nil {
		fmt.Println(err)
	}
	bt, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bt))
}


func TestConsumeViewDay(t *testing.T) {
	c,err:=impl.GetConsumeViewDay(5,"2021-07-17")
	if err != nil {
		fmt.Println(err)
	}
	bt, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bt))
}