package main

import (
	"github.com/codegangsta/martini-contrib/render"
	"github.com/luan/godo/godo"
	"net/http"
)

type ProjectsController struct{}

func (c *ProjectsController) List(r render.Render) {
	projects, _ := godo.NewProjectManager().FindAllWithTasks()

	r.HTML(200, "projects", projects)
}

func (c *ProjectsController) Create(req *http.Request, r render.Render) {

	p := godo.NewProject(req.FormValue("name"))
	godo.NewProjectManager().Add(&p)

	r.Redirect("/projects", 301)
}
