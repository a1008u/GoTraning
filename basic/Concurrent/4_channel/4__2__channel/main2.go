package __2_channel

import (
	"fmt"
	"log"
)

func greet(ch chan string) {
	defer close(ch)
	msg := <- ch
	ch <- fmt.Sprintf("Thanks for %s", msg)
	ch <- "Hello David"
	return
}

//THE PURPOSE
//  チャネルの利用（双方向）　
//INTENTION
//
//arg:
//
//return:
//
func main() {
	ch := make(chan string)
	go greet(ch)

	ch <- "Hello John"

	log.Println(<-ch)
	log.Println(<-ch)
}

