package queue

import "fmt"

type Queue struct {
	Heap []Items
}

type Items struct {
	ID     string
	Weight int
}

func (q *Queue) PushHeap(in Items) {
	q.Heap = append(q.Heap, in)
	q.CompareHeap()
}

func (q *Queue) Len() int {
	return len(q.Heap)
}

func Left(parrent int) int {
	return (2 * parrent) + 1
}

func Right(parrent int) int {
	return (2 * parrent) + 2
}

func Parrent(index int) int {
	return (index - 1) / 2
}

func (q *Queue) Less(x int, y int) bool {
	return q.Heap[x].Weight < q.Heap[y].Weight
}

func (q *Queue) CompareHeap() {
	index := len(q.Heap) - 1

	for {
		if index < 1 || q.Heap[index].Weight > q.Heap[Parrent(index)].Weight {
			break
		}
		if q.Heap[index].Weight <= q.Heap[Parrent(index)].Weight {
			q.Swap(Parrent(index), index)
			index = Parrent(index)
		}
	}
}

func (q *Queue) Swap(x int, y int) {
	q.Heap[x], q.Heap[y] = q.Heap[y], q.Heap[x]
}

func BuildQueue() *Queue {
	return &Queue{}
}

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
