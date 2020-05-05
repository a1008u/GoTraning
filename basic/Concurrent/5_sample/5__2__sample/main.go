package ___2__sample

import (
	"log"
)

func readThem(in <-chan string, out chan<- string) {
	for i := range in {
		log.Println(i)
	}
	out <- "done"
}

// main関数がGoroutineからすべてのメッセージを受信してprintしたことがわかります。
// main関数は、すべての着信メッセージが送信されたことが通知されたときにのみ終了します。
// WaitGroupを必要とせずにチャネルを介してメッセージを渡すことにより、
// 作業が完了したことをGoroutineに別のGoroutineに通知させる。
func main() {
	log.SetFlags(0)
	in, out  := make(chan string), make(chan string)
	go readThem(in, out)

	strs := []string{"a","b", "c", "d", "e", "f"}
	for _, s := range strs {
		in <- s
	}
	close(in)
	<-out
}
