package main

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/luan/godo/godo"
	"net/http"
	"strconv"
	"fmt"
)

type TasksController struct{}

func (c *TasksController) List(r render.Render) {
	tasks, _ := godo.NewTaskManager().FindAll()
	fmt.Println(tasks)

	r.HTML(200, "tasks", tasks)
}

func (c *TasksController) Create(req *http.Request, r render.Render) {

	t := godo.NewTask(req.FormValue("name"))
	godo.NewTaskManager().Add(&t)


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
