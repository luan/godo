package main

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
)

func SetRoutes(m *martini.ClassicMartini) {
	m.Use(render.Renderer())
	tc := &TasksController{}

	m.Get("/tasks", tc.List)
	m.Post("/tasks", tc.Create)
	m.Patch("/tasks/:id", tc.Update)
	m.Post("/tasks/:id", tc.Update) // hack cos' HTML has no PATCH
}

func main() {
	m := martini.Classic()
	SetRoutes(m)
	m.Run()
}
