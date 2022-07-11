package heap_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ScarletTanager/algorithms/heap"
)

var _ = Describe("PriorityQueue", func() {
	var (
		data []int
		h    *heap.Heap[int]
	)

	BeforeEach(func() {
		data = []int{5, 10, 15, 20, 25, 30, 35, 40, 45, 1, 7, 49, 3, 9, 27}
		h = heap.NewMaxHeap(data)
	})

	Describe("Head", func() {
		It("Returns the element at the head of the queue", func() {
			r, e := h.Head()
			Expect(r).To(Equal(49))
			Expect(e).NotTo(HaveOccurred())
		})

		It("Removes the head element from the queue", func() {
			h.Head()
			d := h.Data()
			Expect(d).To(HaveLen(len(data) - 1))
			Expect(d).NotTo(ContainElement(49))
		})
	})

	Describe("Insert", func() {
		var (
			new int
		)

		JustBeforeEach(func() {
			Expect(h.Data()).NotTo(ContainElement(new))
		})

		When("Adding a value that should be near the front of the queue", func() {
			BeforeEach(func() {
				new = 41
			})

			It("Promotes the element to the correct position", func() {
				h.Insert(new)
				Expect(h.Data()[3]).To(Equal(new))
			})
		})

		When("Adding a value that should remain at the back", func() {
			BeforeEach(func() {
				new = 2
			})

			It("Adds the element without incorrectly promoting it", func() {
				h.Insert(new)
				Expect(h.Data()[len(h.Data())-1]).To(Equal(new))
			})
		})
	})

	Describe("SetPriority", func() {
		var (
			new int
		)

		JustBeforeEach(func() {
			Expect(h.Data()).NotTo(ContainElement(new))
		})

		When("Raising the priority", func() {
			BeforeEach(func() {
				new = 37
			})

			It("Promotes the element appropriately", func() {
				// 10 is the position with 1 in it
				h.SetPriority(10, new)
				// Did we replace the old priority?
				Expect(h.Data()).NotTo(ContainElement((1)))
				// Did we promote it?
				Expect(h.Data()[4]).To(Equal(new))
			})
		})

		When("Lowering the priority", func() {
			BeforeEach(func() {
				new = 12
			})

			It("Demotes the element appropriately", func() {
				h.SetPriority(4, new)
				Expect(h.Data()).NotTo(ContainElement(40))
				Expect(h.Data()[8]).To(Equal(new))
			})
		})
	})
})
