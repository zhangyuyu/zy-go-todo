package models

import (
	"github.com/astaxie/beego/orm"
)

type Task struct {
	Id    int64  // Unique identifier
	Title string // Description
	Done  bool   // Is this task done?
}

func init() {
	orm.RegisterModel(new(Task))
}
