package ___3__sample

import (
	"fmt"
	"sync"
)

type Worker struct {
	in, out chan int
	sbw int
	mtx *sync.Mutex
}

func (w *Worker) readThem() {
	w.sbw++
	go func() {
		partial := 0
		for i := range w.in {
			partial += i
		}
		w.out <- partial

		// 安全にsub workerの処理を終了する
		w.mtx.Lock()
		w.sbw--
		if w.sbw == 0 {
			close(w.out)
		}
		w.mtx.Unlock()
	}()

}

func (w *Worker) gatherResult() int {
	total := 0
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for  i:= range w.out{
			total += i
		}
		wg.Done()
	}()

	wg.Wait()
	return total
}

func main() {

	mtx := &sync.Mutex{}
	in := make(chan int, 100)
	out := make(chan int)
	wrk := Worker{in: in, out:out, mtx:mtx}

	wrNum := 10
	for i:=1; i<=wrNum; i++ {
		wrk.readThem()
	}

	for i:=1;i<=100; i++ {
		in <- i
	}
	close(in)

	res := wrk.gatherResult()
	fmt.Println(res)
}
