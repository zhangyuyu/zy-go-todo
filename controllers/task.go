package controllers

import (
	"encoding/json"
	"strconv"
	"github.com/astaxie/beego"
	"github.com/zhangyuyu/zy-go-todo/models"
)

type TaskController struct {
	beego.Controller
}

// Example:
//
//   req: GET /tasks/
//   res: 200 [
//          {"ID": 1, "Title": "Learn Go", "Done": false},
//          {"ID": 2, "Title": "Buy bread", "Done": true}
//        ]
func (this *TaskController) ListTasks() {
	beego.Info("Listing All Tasks")
	tasks, err := models.FindAllTask()
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte(err.Error()))
		return
	}
	this.Data["json"] = tasks
	this.ServeJSON()
}

// Examples:
//
//   req: POST /tasks/ {"Title": ""}
//   res: 400 empty title
//
//   req: POST /tasks/ {"Title": "Buy bread"}
//   res: 200
func (this *TaskController) NewTask() {
	req := struct{ Title string }{}
	beego.Info("Creating a new task", req)
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &req); err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("empty title"))
		return
	}
	task, err := models.CreateTask(req.Title)
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte(err.Error()))
		return
	}
	this.Data["json"] = task
	this.ServeJSON()
}

// Examples:
//
//   req: GET /tasks/1
//   res: 200 {"ID": 1, "Title": "Buy bread", "Done": true}
//
//   req: GET /tasks/42
//   res: 404 task not found
func (this *TaskController) GetTask() {
	id := this.Ctx.Input.Param(":id")
	beego.Info("Getting Task taskId : ", id)
	intid, _ := strconv.ParseInt(id, 10, 64)
	task, err := models.FindTask(intid)
	if err != nil {
		this.Ctx.Output.SetStatus(404)
		this.Ctx.Output.Body([]byte("task not found"))
		return
	}
	this.Data["json"] = task
	this.ServeJSON()
}

// Example:
//
//   req: PUT /tasks/1 {"ID": 1, "Title": "Learn Go", "Done": true}
//   res: 200
//
//   req: PUT /tasks/2 {"ID": 2, "Title": "Learn Go", "Done": true}
//   res: 400 inconsistent task IDs
func (this *TaskController) UpdateTask() {
	beego.Info("Updating task")

	var task models.Task
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &task); err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	_, err := models.UpdateTask(task)

	if err != nil {
		this.Ctx.Output.SetStatus(404)
		this.Ctx.Output.Body([]byte("task not found"))
		return
	}
	this.Data["json"] = task
	this.ServeJSON()
}

// Examples:
//
//   req: DELETE /tasks/1
//   res: 200
//
//   req: DELETE /tasks/42
//   res: 404 task not found
func (this *TaskController) DeleteTask() {
	id := this.Ctx.Input.Param(":id")
	beego.Info("Deleting Task taskId : ", id)
	intid, _ := strconv.ParseInt(id, 10, 64)
	err := models.DeleteTask(intid)
	if err != nil {
		this.Ctx.Output.SetStatus(404)
		this.Ctx.Output.Body([]byte("task not found"))
		return
	}
}