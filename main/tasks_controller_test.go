package main_test

import (
	"github.com/luan/godo/godo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"strconv"
)

var _ = Describe("TasksController", func() {
	BeforeEach(func() {
		godo.ResetTasks()
	})

	Describe("GET /tasks", func() {
		It("sets the get route and renders the template", func() {
			Expect(MethodsFor("/tasks")).To(ContainElement("GET"))
			Get("/tasks")
			Expect(response.Code).To(Equal(200))
		})

		It("returns all the tasks in the list", func() {
			godo.AddTask("do stuff")
			godo.AddTask("do more stuff")

			Get("/tasks")
			Expect(response.Body).To(MatchRegexp("<li class=\"pending\">\\s*do stuff\\s*"))
			Expect(response.Body).To(MatchRegexp("<li class=\"pending\">\\s*do more stuff\\s*"))
		})
	})

	Describe("POST /tasks", func() {
		It("sets the post route", func() {
			Expect(MethodsFor("/tasks")).To(ContainElement("POST"))
			Post("/tasks", map[string]string{"name": "foo"})
			Expect(response.Code).To(Equal(301))
		})

		It("sets the the tasks", func() {
			Post("/tasks", map[string]string{"name": "foo"})
			Post("/tasks", map[string]string{"name": "bar"})
			Expect(len(godo.Tasks())).To(Equal(2))
			Expect(godo.Tasks()[0].Name).To(Equal("foo"))
			Expect(godo.Tasks()[1].Name).To(Equal("bar"))
		})
	})

	Describe("PATCH /tasks/:id", func() {
		var (
			anotherTask     godo.Task
			taskToBeUpdated godo.Task
		)
		BeforeEach(func() {
			anotherTask = godo.AddTask("dont update me")
			taskToBeUpdated = godo.AddTask("a task")
		})

		It("sets the patch route", func() {
			Expect(MethodsFor("/tasks/" + strconv.Itoa(taskToBeUpdated.Id))).To(ContainElement("PATCH"))
			Patch("/tasks/"+strconv.Itoa(taskToBeUpdated.Id), map[string]string{"status": "done"})
			Expect(response.Code).To(Equal(301))
		})

		It("Updates the status of the task", func() {
			var task *godo.Task

			Patch("/tasks/"+strconv.Itoa(taskToBeUpdated.Id), map[string]string{"status": "done"})
			Expect(len(godo.Tasks())).To(Equal(2))
			task, _ = godo.FindTask(taskToBeUpdated.Id)
			Expect(task.Status).To(Equal("done"))

			Patch("/tasks/"+strconv.Itoa(taskToBeUpdated.Id), map[string]string{"status": "pending"})
			task, _ = godo.FindTask(taskToBeUpdated.Id)
			Expect(task.Status).To(Equal("pending"))
		})
	})
})
