package __waitGroup

import (
	"log"
	"sync"
)

func sum(from, to int, wg *sync.WaitGroup, res *int) {
	defer wg.Done()
	*res = 0
	for i := from; i <= to; i++ {
		*res += i
	}
	return
}

func main() {
	s1 := 0
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go sum(1, 100, wg, &s1)
	wg.Wait()

	log.Println(s1)
}
