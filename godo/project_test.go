package godo_test

import (
	. "github.com/luan/godo/godo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Project", func() {
	BeforeEach(func() {
		ResetDatabase()
	})


	Describe("Create a projects", func() {
		It("Creates and returns a project", func() {
			project := NewProject("my Project")
			Expect(project.Name).To(Equal("my Project"))
		})

	})

	Describe("Getting Projects with Tasks", func() {
		It("Returns ProjectTaskViews", func() {
			project_manager := NewProjectManager()
			proj1 := NewProject("p1")
			proj2 := NewProject("p2")
			project_manager.Add(&proj1)
			project_manager.Add(&proj2)

			task1 := NewTaskWithProjectID("t1", proj1.ID)
			task2 := NewTaskWithProjectID("t2", proj1.ID)
			task3 := NewTaskWithProjectID("t3", proj2.ID)
			NewTaskManager().Add(&task1)
			NewTaskManager().Add(&task2)
			NewTaskManager().Add(&task3)

			projects, _ := project_manager.FindAllWithTasks()

			Expect(projects[0].Tasks[0].Name).To(Equal("t1"))
			Expect(projects[0].Tasks[1].Name).To(Equal("t2"))
			Expect(projects[0].Name).To(Equal("p1"))
			Expect(projects[1].Tasks[0].Name).To(Equal("t3"))
			Expect(projects[1].Name).To(Equal("p2"))
		})
		It("Returns ProjectTaskViews even without any tasks", func() {
			project_manager := NewProjectManager()
			proj1 := NewProject("p1")
			project_manager.Add(&proj1)

			projects, _ := project_manager.FindAllWithTasks()
			Expect(projects[0].Name).To(Equal("p1"))
			Expect(len(projects[0].Tasks)).To(Equal(0))
		})
	})

})
