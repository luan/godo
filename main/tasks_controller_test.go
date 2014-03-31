package main_test

import (
	"github.com/luan/godo/godo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"strconv"
)

var _ = Describe("TasksController", func() {
	BeforeEach(func() {
		godo.ResetDatabase()
	})

	Describe("GET /tasks", func() {
		It("sets the get route and renders the template", func() {
			Expect(MethodsFor("/tasks")).To(ContainElement("GET"))
			Get("/tasks")
			Expect(response.Code).To(Equal(200))
		})

		It("returns all the tasks in the list", func() {
			t1 := godo.NewTask("do stuff")
			t2 := godo.NewTask("do more stuff")
			godo.NewTaskManager().Add(&t1)
			godo.NewTaskManager().Add(&t2)

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

			tasks, _ := godo.NewTaskManager().FindAll()

			Expect(len(tasks)).To(Equal(2))
			Expect(tasks[0].Name).To(Equal("foo"))
			Expect(tasks[1].Name).To(Equal("bar"))
		})
	})

	Describe("PATCH /tasks/:id", func() {
		var (
			anotherTask     godo.Task
			taskToBeUpdated godo.Task
		)
		BeforeEach(func() {
			anotherTask = godo.NewTask("dont update me")
			taskToBeUpdated = godo.NewTask("a task")
			godo.NewTaskManager().Add(&anotherTask)
			godo.NewTaskManager().Add(&taskToBeUpdated)
		})

		It("sets the patch route", func() {
			Expect(MethodsFor("/tasks/" + strconv.Itoa(taskToBeUpdated.ID))).To(ContainElement("PATCH"))
			Patch("/tasks/"+strconv.Itoa(taskToBeUpdated.ID), map[string]string{"status": "done"})
			Expect(response.Code).To(Equal(301))
		})

		It("Updates the status of the task", func() {
			Patch("/tasks/"+strconv.Itoa(taskToBeUpdated.ID), map[string]string{"status": "done"})
			tasks, _ := godo.NewTaskManager().FindAll()

			Expect(len(tasks)).To(Equal(2))
			task := &godo.Task{}
			godo.NewTaskManager().Find(taskToBeUpdated.ID, task)
			Expect(task.Status).To(Equal("done"))

			task2 := &godo.Task{}
			Patch("/tasks/"+strconv.Itoa(taskToBeUpdated.ID), map[string]string{"status": "pending"})
			godo.NewTaskManager().Find(taskToBeUpdated.ID, task2)
			Expect(task2.Status).To(Equal("pending"))
		})
	})
})
