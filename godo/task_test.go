package godo_test

import (
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

	Describe("FindTask", func() {
		BeforeEach(func() {
			AddTask("task one")
			AddTask("task two")
		})

		It("Finds a task by its Id", func() {
			var (
				task *Task
				err  error
			)

			task, err = FindTask(1)
			Expect(err).To(BeNil())
			Expect(task.Name).To(Equal("task one"))

			task, err = FindTask(2)
			Expect(err).To(BeNil())
			Expect(task.Name).To(Equal("task two"))

			task, err = FindTask(3)
			Expect(err.Error()).To(Equal("Couldn't find task with Id: 3"))
		})
	})
})
