package main

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/luan/godo/godo"
	"net/http"
	"strconv"
)


func SetRoutes(m *martini.ClassicMartini) {
	m.Use(render.Renderer())
	m.Get("/tasks", func(r render.Render) {
		r.HTML(200, "tasks", godo.Tasks())
	})

	m.Post("/tasks", func(req *http.Request, r render.Render) {
		godo.AddTask(req.FormValue("name"))
		r.Redirect("/tasks", 301)
	})

	m.Patch("/tasks/:id", func(params martini.Params, req *http.Request, r render.Render) {
		doPatch(params, req, r)
	})
  m.Post("/tasks/:id", func(params martini.Params, req *http.Request, r render.Render) {
		doPatch(params, req, r)
	})

}

func doPatch(params martini.Params, req *http.Request, r render.Render) {
		id, _:= strconv.Atoi(params["id"])
		task, _ := godo.FindTask(id)
		task.Status = req.FormValue("status")
		r.Redirect("/tasks", 301)
}


func main() {
	m := martini.Classic()
	SetRoutes(m)
	m.Run()
}
