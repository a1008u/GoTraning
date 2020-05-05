package __3__channel

import (
	"log"
	"time"
)



func push(from,to int, out chan int) {
	for i:=from;i<=to; i++ {
		out <- i
		time.Sleep(time.Microsecond)
	}
}

//THE PURPOSE
//  チャネルの利用
//INTENTION
//
//arg:
//
//return:
//
func main() {
	s1 := 0

	ch := make(chan int, 100)

	go push(1,25, ch)
	go push(26,50,ch)
	go push(51,75,ch)
	go push(76,100, ch)

	for c :=0; c< 100; c++ {
		i := <- ch
		log.Println(i)
		s1 += i
	}

	log.Println(s1)
}
