package heap_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestHeap(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Heap Suite")
}
