// using min-heap to get products that have the lightest weight.
package main

import (
	"fmt"
	"log"
	"quocbang/min-heap/queue"

	fake "github.com/brianvoe/gofakeit/v6"
)

func main() {
	q := queue.BuildQueue()
	// push new product to slice.
	for i := 0; i < 100; i++ {
		q.PushHeap(queue.Items{
			ID:     fake.Company(),
			Weight: int(fake.Uint32()),
		})
	}

	// remove root(top) this is lightest weight product.
	productNeedsToExtract := q.Len()
	for i := 0; i < productNeedsToExtract; i++ {
		extractElement, err := q.Extract()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(extractElement)
	}
}
