package godo_test

import (
	. "github.com/luan/godo/godo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Task", func() {
	var (
		tm *TaskManager
	)

	BeforeEach(func() {
		tm = NewTaskManager()
		ResetDatabase()
	})

	Describe("NextStatus", func() {
		It("returns the next possible status of the task", func() {
			task := NewTask("my task")
			tm.Add(&task)
			task.Status = "pending"
			Expect(task.NextStatus()).To(Equal("done"))
			task.Status = "done"
			Expect(task.NextStatus()).To(Equal("pending"))
		})
	})

	Describe("Adding Tasks", func() {
		It("Creates and returns a Task", func() {
			task := NewTask("my task")
			Expect(task.Name).To(Equal("my task"))
			Expect(task.Status).To(Equal("pending"))
		})

		It("Sets unique identifier to each task", func() {
			task1 := NewTask("my task")
			task2 := NewTask("other task")
			tm.Add(&task1)
			tm.Add(&task2)

			Expect(task1.ID).NotTo(Equal(0))
			Expect(task2.ID).NotTo(Equal(0))
			Expect(task1.ID).NotTo(Equal(task2.ID))
		})

		It("Adds each task to the list", func() {
			task1 := NewTask("my task")
			task2 := NewTask("other task")
			tm.Add(&task1)
			tm.Add(&task2)


			tasks, _ := tm.FindAll()

			Expect(len(tasks)).To(Equal(2))
		})
	})

	Describe("Update Tasks", func() {
		It("updates", func() {
			t := NewTask("my task")
			tm.Add(&t)
			t.Name = "ANOTHER NAME"
			tm.Update(&t)

			foundTask := Task{}
			tm.Find(t.ID, &foundTask)
			Expect(foundTask.Name).To(Equal("ANOTHER NAME"))
		})
	})

	Describe("FindTask", func() {

		var (
			task1 Task
			task2 Task
		)
		BeforeEach(func() {
			task1 = NewTask("task one")
			tm.Add(&task1)
			task2 = Task{Name: "task two", Status: "pending"}
			_ = Dbmap.Insert(&task2)
			tm.Add(&task2)
		})

		It("Finds a task by its ID", func() {
			var (
				task Task
				err  error
			)

			err = tm.Find(task1.ID, &task)
			Expect(err).To(BeNil())
			Expect(task.Name).To(Equal("task one"))

			err = tm.Find(task2.ID, &task)
			Expect(err).To(BeNil())
			Expect(task.Name).To(Equal("task two"))

			err = tm.Find(38938383, &task)
			Expect(err.Error()).To(Equal("sql: no rows in result set"))
		})
	})
})
