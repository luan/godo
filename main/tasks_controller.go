package main

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/luan/godo/godo"
	"net/http"
	"strconv"
)

type TasksController struct{}

func (c *TasksController) List(r render.Render) {
	tasks, _ := godo.NewTaskManager().FindAll()
	r.HTML(200, "tasks", tasks)
}

func (c *TasksController) Create(req *http.Request, r render.Render) {
	godo.NewTaskManager().Add(req.FormValue("name"))
	r.Redirect("/tasks", 301)
}

func (c *TasksController) Update(params martini.Params, req *http.Request, r render.Render) {
	id, _ := strconv.Atoi(params["id"])

	task := &godo.Task{}
	godo.NewTaskManager().Find(id, task)

	task.Status = req.FormValue("status")
	godo.NewTaskManager().Update(task)

	r.Redirect("/tasks", 301)
}
