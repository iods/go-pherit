package main

import (
	"fmt"
	"github.com/iods/go-pherit/pkg/learngo/grape"
	"github.com/iods/go-pherit/pkg/learngo/greeter"
	"time"
)

func main() {
	greeter.Test()

	channel := make(chan string)
	go greeter.Send(channel)

	greeter.ReportNap("receiving go routine.", 5)
	fmt.Println(<-channel)
	fmt.Println(<-channel)

	pages := make(chan grape.Page)
	urls := []string{
		"https://ryemiller.io",
		"https://golang.org",
		"https://golang.org/doc",
	}

	for _, url := range urls {
		go grape.ResponseSize(url, pages)
	}

	for i := 0; i < len(urls); i++ {
		page := <- pages
		fmt.Printf("%s: %d\n", page.URL, page.Size)
	}

	grape.ResponseBody()

	time.Sleep(5 * time.Second)
}
