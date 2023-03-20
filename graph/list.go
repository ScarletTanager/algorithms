package graph

import (
	"errors"
)

func (a AdjacencyList) Link(source, target int) error {
	svp := a[source]
	if svp.edgeIndices == nil {
		svp.edgeIndices = make([]int, 0)
	}

	// It is entirely allowable to have multiple edges linking the same two
	// vertices, but our current model does not provide any good way to distinguish
	// between edges connecting the same vertices, so the lack of a uniqueness test
	// here is a bug, not really a feature.  We'll want to make edges actual "things"
	// later on.
	svp.edgeIndices = append(svp.edgeIndices, target)
	return nil
}

func (a AdjacencyList) LinkBoth(source, target int) {

}

const (
	AttrColor      string = "color"
	AttrColorWhite string = "white"
	AttrColorGray  string = "gray"
	AttrColorBlack string = "black"

	AttrDistance string = "distance"
	AttrParent   string = "parent"
)

func (a AdjacencyList) SearchBreadthFirst(sourceIndex int) (*Vertex, error) {
	var (
		svp *Vertex
		err error
	)

	a.reset()
	// We can either make the channel large enough for the case where every
	// Vertex has an edge to every other Vertex (n vertices * (n-1) others),
	// or we can total up the adjacencies.  For now we're brute-forcing.
	queue := make(chan int, len(a)*(len(a)-1))

	if svp, err = a.AtIndex(sourceIndex); err != nil {
		return nil, err
	}

	svp.Set(AttrColor, AttrColorGray)
	svp.Set(AttrDistance, 0)

	queue <- sourceIndex

	for len(queue) > 0 {
		i := <-queue
		vp, _ := a.AtIndex(i)

		// Visit each of the adjacent vertices
		for _, adj := range vp.edgeIndices {
			avp, _ := a.AtIndex(adj)

			if avp.Get(AttrColor).(string) == AttrColorWhite { // We have not visited this vertex yet
				avp.Set(AttrColor, AttrColorGray) // We've visited here but not looked at its adjacencies yet
				avp.Set(AttrDistance, vp.Get(AttrDistance).(int)+1)
				avp.Set(AttrParent, i)
				queue <- adj
			}
		}

		// Mark this vertex as finished
		vp.Set(AttrColor, AttrColorBlack)
	}

	return svp, nil
}

func (a AdjacencyList) reset() {
	for i := 0; i < len(a); i++ {
		vp, _ := a.AtIndex(i)
		vp.Set(AttrColor, AttrColorWhite)
		vp.Delete(AttrDistance)
		vp.Delete(AttrParent)
	}
}

func (a AdjacencyList) Path(source, target int) ([]*Vertex, error) {
	var (
		svp         *Vertex
		rpath, path []*Vertex
		err         error
	)

	svp, err = a.AtIndex(source)
	if err != nil {
		return nil, err
	}

	rpath = make([]*Vertex, 0)
	if source == target {
		rpath = append(rpath, svp)
		return rpath, nil
	}

	var tvp *Vertex
	ti := target
	for {
		tvp, err = a.AtIndex(ti)
		if err != nil {
			return nil, err
		}

		pi := tvp.Get(AttrParent)
		if pi == nil {
			return nil, errors.New("No path found from source to target")
		}

		// Step up the tree
		ti = pi.(int)

		rpath = append(rpath, tvp)
		if pi == source {
			rpath = append(rpath, svp)
			break
		}
	}

	path = make([]*Vertex, len(rpath))
	for i := len(rpath); i > 0; i-- {
		path[len(rpath)-i] = rpath[i-1]
	}

	return path, nil
}

func (a AdjacencyList) AtIndex(index int) (*Vertex, error) {
	if index > len(a)-1 {
		return nil, errors.New("Graph does not contain index")
	}

	return a[index], nil
}
