package main

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// Inquery is mel.
type Inquery struct {
	Id      int       `json: "id"`
	Content string    `json: "content"`
	Created time.Time `json: "created"`
}

// Inqueries are models.
type Inqueries []Inquery

func init() {
	orm.RegisterModel(new(Inquery))
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		panic(err)
	}
}
