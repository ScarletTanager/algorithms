package graph

// Vertex represents a node or vertex in a graph.  An adjacency list is just a slice of unique Vertex instances.
type Vertex struct {
	Attributes  Attributes
	index       int
	edgeIndices []int
}

// Attributes are specific to the graph in use
type Attributes map[string]interface{}

func (v *Vertex) Set(attribute string, value interface{}) {
	if v.Attributes == nil {
		v.Attributes = make(Attributes)
	}
	v.Attributes[attribute] = value
}

// Get returns the attribute value for the named attribute.  If the Vertex has no such attribute, nil
// is returned.
func (v *Vertex) Get(attribute string) interface{} {
	if val, ok := v.Attributes[attribute]; ok {
		return val
	}

	return nil
}

// Delete removes the specified attribute entirely
func (v *Vertex) Delete(attribute string) {
	delete(v.Attributes, attribute)
}

// Index returns the Vertex's index *within the graph* - this is only meaningful
// after the vertex has been used to create the graph
func (v *Vertex) Index() int {
	return v.index
}

// New creates a new graph.
func New(vertices []Vertex) (Graph, error) {
	var l AdjacencyList

	if vertices != nil {
		l = make(AdjacencyList, len(vertices))
		for i, _ := range vertices {
			vToUse := vertices[i]
			vToUse.index = i
			l[i] = &vToUse
		}
	}

	return l, nil
}

type AdjacencyList []*Vertex

type Graph interface {
	// AtIndex returns the Vertex at the specified position in the graph
	AtIndex(int) (*Vertex, error)
	// WithAttribute returns a slice of Vertices with the the specified value for the given attribute.
	// In the case where no vertices are found with the given attribute, nil is returned.
	WithAttribute(string, interface{}) []*Vertex
	// SearchBreadthFirst performs a breadth-first search of the graph starting from the
	// Vertex at the specified index.
	SearchBreadthFirst(int) (*Vertex, error)
	// Link creates a directed edge from source to target.  The source Vertex will own the
	// edge - meaning that the target Vertex will have no knowledge of it (you can only navigate
	// from source to target)
	Link(int, int) error
	// LinkBoth creates a bidirectionally navigable edge linking the source and target vertices
	LinkBoth(int, int)
	// Path returns either an ordered slice of vertices (if a path exists from the source to the target)
	// or nil
	Path(int, int) ([]*Vertex, error)
}
