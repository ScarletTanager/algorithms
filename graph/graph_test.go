package graph_test

import (
	"github.com/ScarletTanager/algorithms/graph"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Graph", func() {
	Describe("Vertex", func() {
		var (
			v                  *graph.Vertex
			attrName, expected string
		)

		BeforeEach(func() {
			v = &graph.Vertex{
				Attributes: make(graph.Attributes),
			}

			attrName = "realAttribute"
		})

		Describe("Get", func() {
			Context("When the Vertex has no such attribute", func() {
				BeforeEach(func() {
					attrName = "nonexistentAttribute"
				})

				It("Returns nil", func() {
					val := v.Get(attrName)
					Expect(val).To(BeNil())
				})
			})

			Context("When the Vertex has the given attribute", func() {
				BeforeEach(func() {
					expected = "foobar"
					v.Attributes[attrName] = expected
				})

				It("Returns the correct value", func() {
					Expect(v.Get(attrName).(string)).To(Equal(expected))
				})
			})
		})

		Describe("Set", func() {
			BeforeEach(func() {
				expected = "realValue"
			})
			JustBeforeEach(func() {
				Expect(v.Get(attrName)).To(BeNil())
			})

			It("Sets the attribute value", func() {
				v.Set(attrName, expected)
				val := v.Get(attrName)
				Expect(val).NotTo(BeNil())
				Expect(val.(string)).To(Equal(expected))
			})

			When("The Attributes field has not been initialized", func() {
				BeforeEach(func() {
					v = &graph.Vertex{}
				})

				JustBeforeEach(func() {
					Expect(v.Attributes).To(BeNil())
				})

				It("Initializes Attributes before setting the value", func() {
					v.Set(attrName, expected)
					val := v.Get(attrName)
					Expect(val).NotTo(BeNil())
					Expect(val.(string)).To(Equal(expected))
				})
			})
		})
	})
})
