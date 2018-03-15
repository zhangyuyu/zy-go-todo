package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

func FindTask(id int64) (*Task, error) {
	task := Task{Id: id}

	o := orm.NewOrm()
	err := o.Read(&task)
	return &task, err
}


func FindAllTask() ([]*Task, error) {
	o := orm.NewOrm()
	var tasks []*Task
	_, err := o.QueryTable("task").All(&tasks)
	return tasks, err
}

func CreateTask(title string) (*Task, error) {
	if title == "" {
		return nil, fmt.Errorf("empty title")
	}
	task := Task{0, title, false}

	o := orm.NewOrm()
	_, err := o.Insert(&task)

	return &task, err
}

func UpdateTask(task Task) (*Task, error) {
	o := orm.NewOrm()

	_, err := o.Update(&task, "Done")
	return &task, err
}