package graph_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ScarletTanager/algorithms/graph"
)

var _ = Describe("List", func() {
	var (
		g        graph.Graph
		e        error
		vertices []graph.Vertex
	)

	BeforeEach(func() {
		vertices = []graph.Vertex{
			{Attributes: make(graph.Attributes)},
			{Attributes: graph.Attributes{
				"foo": "bar",
			}},
			{Attributes: graph.Attributes{
				"foo": "john",
			}},
			{Attributes: graph.Attributes{
				"foo": "mary",
			}},
		}
	})

	JustBeforeEach(func() {
		g, e = graph.New(vertices)
		Expect(e).NotTo(HaveOccurred())
		Expect(g).NotTo(BeNil())
	})

	Describe("AtIndex", func() {
		When("The requested index is not in the graph", func() {
			It("Returns nil", func() {
				v, _ := g.AtIndex(4)
				Expect(v).To(BeNil())
			})

			It("Returns an error", func() {
				_, e = g.AtIndex(4)
				Expect(e).To(HaveOccurred())
			})
		})

		When("The requested index is within the graph", func() {
			It("Returns the vertex at that position", func() {
				v, _ := g.AtIndex(2)
				Expect(*v).To(Equal(vertices[2]))
			})
		})
	})

	Describe("WithAttribute", func() {
		When("No Vertex has the specified attribute and value", func() {
			It("Returns nil", func() {
				vs := g.WithAttribute("firstName", "frank")
				Expect(vs).To(BeNil())
			})
		})

		When("Only one Vertex has the specified attribute and value", func() {
			It("Returns a list containing only the single matching Vertex", func() {
				vs := g.WithAttribute("foo", "john")
				Expect(vs).To(HaveLen(1))
				expected, _ := g.AtIndex(2)
				Expect(vs[0]).To(Equal(expected))
			})
		})

		When("Multiple vertices have the specified attribute and value", func() {
			BeforeEach(func() {
				vertices[1].Attributes["birthstone"] = "opal"
				vertices[3].Attributes["birthstone"] = "opal"
			})

			It("Returns a list containing the matching vertices", func() {
				vs := g.WithAttribute("birthstone", "opal")
				Expect(vs).To(HaveLen(2))
				e1, _ := g.AtIndex(1)
				e2, _ := g.AtIndex(3)
				Expect(vs).To(ConsistOf(e1, e2))
			})
		})
	})

	Describe("Path", func() {
		var (
			source, target int
		)

		BeforeEach(func() {
			source = 0
			target = 3
		})

		When("No path exists between the source and target", func() {
			It("Returns nil", func() {
				p, err := g.Path(0, 3)
				Expect(p).To(BeNil())
				Expect(err).To(HaveOccurred())
			})
		})

		When("A path exists between the source and target", func() {
			JustBeforeEach(func() {
				for i := source; i < target; i++ {
					g.Link(i, i+1)
				}
			})

			When("But the graph has not yet been searched", func() {
				It("Returns nil", func() {
					p, err := g.Path(0, 3)
					Expect(p).To(BeNil())
					Expect(err).To(HaveOccurred())
				})
			})

			When("A breadth-first search has been performed", func() {
				JustBeforeEach(func() {
					g.SearchBreadthFirst(0)
				})

				It("Returns a slice containing the vertices to traverse on the path, with distances marked correctly", func() {
					p, err := g.Path(source, target)
					Expect(err).NotTo(HaveOccurred())
					Expect(p).NotTo((BeNil()))
					Expect(p).To(HaveLen(target + 1))
					for i := 0; i <= 3; i++ {
						vp, _ := g.AtIndex(i)
						Expect(p[i]).To(Equal(vp))
						Expect(p[i].Get(graph.AttrDistance).(int)).To(Equal(i))
					}
				})

				When("But the search tree does not include the path", func() {
					JustBeforeEach(func() {
						g.SearchBreadthFirst(1)
					})

					It("Returns nil (no path found)", func() {
						p, err := g.Path(0, 3)
						Expect(p).To(BeNil())
						Expect(err).To(HaveOccurred())
					})
				})
			})
		})
	})
})
