package heap

import "errors"

func (h *Heap[O]) Head() (O, error) {
	var retVal O
	if len(h.data) < 1 {
		return retVal, errors.New("Underflow error")
	}

	retVal = h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.truncate()

	h.Maxify(1)

	return retVal, nil
}

func (h *Heap[O]) Insert(elem O) {
	h.data = append(h.data, elem)
	h.raise(len(h.data))
}

func (h *Heap[O]) SetPriority(pos int, pri O) error {
	if !h.isValidPosition(pos) {
		return errors.New("Invalid queue position")
	}

	current := h.data[pos-1]
	h.data[pos-1] = pri
	if pri > current {
		h.raise(pos)
	}

	if pri < current {
		h.Maxify(pos)
	}

	// If pri == current, do nothing
	return nil
}

func (h *Heap[O]) raise(pos int) {
	pri := h.data[pos-1]
	p := parent(pos)

	for h.isValidPosition(p) && h.data[p-1] < pri {
		h.swap(p, pos)
		pos = p
		p = parent(pos)
	}
}

func parent(pos int) int {
	return pos / 2
}
