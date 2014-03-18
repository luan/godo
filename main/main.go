package main

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/martini-contrib/binding"
)


type App struct {
	tasks []Task
}

type Task struct {
	Name string `form:"name" binding:"required"`
}

func NewApp() *App {
	return &App{[]Task{}}
}

func SetRoutes(app *App, m *martini.ClassicMartini) {
	m.Use(render.Renderer())
	m.Get("/tasks", func(r render.Render) {
		r.HTML(200, "tasks", app.tasks)
	})

	m.Post("/tasks", binding.Bind(Task{}), binding.ErrorHandler, app.AddTask, func(r render.Render) {
		r.Redirect("/tasks", 301)
	})
}

func (app *App) AddTask(task Task) {
	app.tasks = append(app.tasks, task)
}

func (app *App) Tasks() []Task {
	return app.tasks
}

func main() {
	app := NewApp()
	m := martini.Classic()
	SetRoutes(app, m)
	m.Run()
}
