package greeter

import (
	"fmt"
	"time"
)

func abc(channel chan string) {
	channel <- "one"
	channel <- "three"
	channel <- "five"
}

func def(channel chan string) {
	channel <- "two"
	channel <- "four"
	channel <- "six"
}

func Test() {
	a := make(chan string)
	b := make(chan string)

	go abc(a)
	go def(b)

	fmt.Println(<-a)
	fmt.Println(<-b)
	fmt.Println(<-a)
	fmt.Println(<-b)
	fmt.Println(<-a)
	fmt.Println(<-b)
}

func ReportNap(name string, delay int) {
	for i := 0; i < delay; i++ {
		fmt.Println(name, "sleeping.")
		time.Sleep(4 * time.Second)
	}
	fmt.Println("Wakes up.")
}

func Send(channel chan string) {
	ReportNap("sending go routine", 2)
	fmt.Println("****** sending value *********")
	channel <- "a"
	fmt.Println("*** sending value ***")
	channel <- "b"
}

