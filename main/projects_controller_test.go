package main_test

import (
	"github.com/luan/godo/godo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	_ "strconv"
)

var _ = Describe("ProjectsController", func() {
	BeforeEach(func() {
		godo.ResetDatabase()
	})

	Describe("GET+POST /projects", func() {
		It("sets the get route and renders the template", func() {
			Post("/projects", map[string]string{"name": "p1"})
			Post("/projects", map[string]string{"name": "p2"})
			Expect(response.Code).To(Equal(301))

			Get("/projects")
			Expect(response.Body).To(MatchRegexp("<li>p1"))
			Expect(response.Body).To(MatchRegexp("<li>p2"))
		})
	})
})
