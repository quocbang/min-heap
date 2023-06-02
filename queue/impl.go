package queue

import "fmt"

// Queue definition.
type Queue struct {
	Heap []Items
}

// Item definition.
type Items struct {
	ID     string
	Weight int
}

// PushHeap is pushing the new product to the slice.
func (q *Queue) PushHeap(in Items) {
	q.Heap = append(q.Heap, in)
	q.CompareHeap()
}

// Len checks how many products are available.
func (q *Queue) Len() int {
	return len(q.Heap)
}

// Left return left child position.
func Left(Parent int) int {
	return (2 * Parent) + 1
}

// Right return right child position.
func Right(Parent int) int {
	return (2 * Parent) + 2
}

// Parent return parent position.
func Parent(index int) int {
	return (index - 1) / 2
}

// Less check whether x is less than y.
func (q *Queue) Less(x int, y int) bool {
	return q.Heap[x].Weight < q.Heap[y].Weight
}

// CompareHeap check and compare slice.
func (q *Queue) CompareHeap() {
	index := len(q.Heap) - 1
	for {
		if index < 1 || q.Heap[index].Weight > q.Heap[Parent(index)].Weight {
			break
		}
		if q.Heap[index].Weight <= q.Heap[Parent(index)].Weight {
			q.Swap(Parent(index), index)
			index = Parent(index)
		}
	}
}

// Swap is swap between the x and y positions.
func (q *Queue) Swap(x int, y int) {
	q.Heap[x], q.Heap[y] = q.Heap[y], q.Heap[x]
}

func BuildQueue() *Queue {
	return &Queue{}
}

// Extract root node, also the lightest weight.
func (q *Queue) Extract() (string, error) {
	if q.Len() < 1 { // only one element
		return "product is epmty", nil
	}
	extractProduct := q.Heap[0].ID
	extractWeight := q.Heap[0].Weight
	lastElement := q.Len() - 1
	q.Swap(0, lastElement)        // swap root with last element.
	q.Heap = q.Heap[:lastElement] // remove last element.

	index := 0
	for Left(index) < lastElement && Right(index) < lastElement {
		if q.Less(Left(index), Right(index)) {
			if q.Heap[Left(index)].Weight < q.Heap[index].Weight {
				q.Swap(Left(index), index)
				index = Left(index)
			} else {
				break
			}
		} else {
			if q.Heap[Right(index)].Weight < q.Heap[index].Weight {
				q.Swap(Right(index), index)
				index = Right(index)
			} else {
				break
			}
		}
		lastElement = q.Len() - 1
	}
	return fmt.Sprintf("product [%v] wgt = [%v] was extracted", extractProduct, extractWeight), nil
}
