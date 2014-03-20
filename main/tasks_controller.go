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
	r.HTML(200, "tasks", godo.Tasks())
}

func (c *TasksController) Create(req *http.Request, r render.Render) {
	godo.AddTask(req.FormValue("name"))
	r.Redirect("/tasks", 301)
}

func (c *TasksController) Update(params martini.Params, req *http.Request, r render.Render) {
	id, _ := strconv.Atoi(params["id"])
	task, _ := godo.FindTask(id)
	task.Status = req.FormValue("status")
	r.Redirect("/tasks", 301)
}
