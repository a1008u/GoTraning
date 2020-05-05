package __mutex

import (
	"fmt"
	"log"
	"sync"
)

func sum(from,to int, wg *sync.WaitGroup, res *string, mtx *sync.Mutex) {
	defer wg.Done()
	for i:=from;i<=to; i++ {
		mtx.Lock()
		*res = fmt.Sprintf("%s|%d|",*res, i)
		mtx.Unlock()
	}
	return
}

//THE PURPOSE
//  一つの変数に値を足す（ゴルーチンを利用した共有変数利用）
//INTENTION
//
//arg:
//
//return:
//
func main() {
	s1 := ""
	mtx := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	wg.Add(4)

	go sum(1,25, wg,&s1, mtx)
	go sum(26,50, wg, &s1, mtx)
	go sum(51,75, wg, &s1, mtx)
	go sum(76,100, wg, &s1, mtx)
	wg.Wait()

	log.Println(s1)

}
