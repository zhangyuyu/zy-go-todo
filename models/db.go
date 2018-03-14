package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

func Find(id int64) (*Task, bool) {
	task := Task{Id: id}

	o := orm.NewOrm()
	err := o.Read(&task)
	if err != nil {
		fmt.Println(err)
		return nil, false
	} else {
		return &task, true
	}
}
