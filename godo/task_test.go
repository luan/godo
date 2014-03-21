package godo_test

import (
	"fmt"
	. "github.com/luan/godo/godo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Task", func() {
	BeforeEach(func() {
		ResetTasks()

	})

	Describe("NextStatus", func() {
		It("returns the next possible status of the task", func() {
			task := AddTask("my task")
			task.Status = "pending"
			Expect(task.NextStatus()).To(Equal("done"))
			task.Status = "done"
			Expect(task.NextStatus()).To(Equal("pending"))
		})
	})

	Describe("AddTask", func() {
		It("Creates and returns a Task", func() {
			task := AddTask("my task")
			Expect(task.Name).To(Equal("my task"))
			Expect(task.Status).To(Equal("pending"))
		})

		It("Sets unique identifier to each task", func() {
			task1 := AddTask("my task")
			task2 := AddTask("other task")
			Expect(task1.Id).NotTo(Equal(0))
			Expect(task2.Id).NotTo(Equal(0))
			Expect(task1.Id).NotTo(Equal(task2.Id))
		})

		It("Adds each task to the list", func() {
			AddTask("my task")
			AddTask("another task")

			Expect(len(Tasks())).To(Equal(2))
		})
	})

	Describe("Update Tasks", func() {
		It("updates", func() {
			t := AddTask("my task")
			t.Name = "ANOTHER NAME"
			UpdateTask(t)
			t, _ = FindTask(t.Id)
			Expect(t.Name).To(Equal("ANOTHER NAME"))
		})
	})

	Describe("FindTask", func() {

		var (
			task1 Task
			task2 Task
		)
		BeforeEach(func() {
			task1 = AddTask("task one")
			task2 = Task{Name: "task two", Status: "pending"}
			_ = Dbmap.Insert(&task2)
		})

		It("Finds a task by its Id", func() {
			var (
				task Task
				err  error
			)

			task, err = FindTask(task1.Id)
			Expect(err).To(BeNil())
			Expect(task.Name).To(Equal("task one"))

			task, err = FindTask(task2.Id)
			Expect(err).To(BeNil())
			Expect(task.Name).To(Equal("task two"))

			task, err = FindTask(38938383)
			fmt.Println(err)
			Expect(err.Error()).To(Equal("Couldn't find task with Id: 38938383"))
		})
	})
})
