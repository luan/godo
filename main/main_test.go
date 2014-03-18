package main_test

import (
	"github.com/codegangsta/martini"
	. "github.com/luan/godo/main"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("App", func() {
	var m *martini.ClassicMartini
	var app *App

	BeforeEach(func() {
		app = NewApp()
		m = martini.Classic()
		SetRoutes(app, m)
	})

	Describe("Tasks", func() {
		Describe("GET /tasks", func() {
			It("sets the get route and renders the template", func() {
				Expect(m.MethodsFor("/tasks")).To(ContainElement("GET"))
				Get(m, "GET", "/tasks")
				Expect(response.Code).To(Equal(200))
			})

			It("returns all the tasks in the list", func() {
				app.AddTask(Task{"do stuff"})
				app.AddTask(Task{"do more stuff"})

				Get(m, "GET", "/tasks")
				Expect(response.Body).To(MatchRegexp("<li>\\s*do stuff\\s*</li>\\s*<li>\\s*do more stuff\\s*</li>"))
			})
		})

		Describe("POST /tasks", func() {
			It("sets the post route", func() {
				Expect(m.MethodsFor("/tasks")).To(ContainElement("POST"))

				Post(m, "POST", "/tasks", map[string]string{"name": "foo"})
				Expect(response.Code).To(Equal(301))
			})

			It("sets the the tasks", func() {
				Post(m, "POST", "/tasks", map[string]string{"name": "foo"})
				Post(m, "POST", "/tasks", map[string]string{"name": "bar"})
				Expect(app.Tasks()).To(Equal([]Task{Task{"foo"}, Task{"bar"}}))
			})
		})
	})
})
