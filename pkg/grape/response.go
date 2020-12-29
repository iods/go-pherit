// grape is a web scraper in go
/*
Use short declarations for urls, go routines.
a, b, c := "https://ryemiller.io", "https://golang.org", "https://golang.org/doc"
grape.ResponseBody()
go grape.ResponseSize(a)
go grape.ResponseSize(b)
go grape.ResponseSize(c)
time.Sleep(5 * time.Second)
 */
package grape

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	URL  string
	Size int
}

func ResponseSize(url string, channel chan Page) {
	fmt.Println("Getting:", url)

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	channel <- Page{URL: url, Size: len(body)}
}

func ResponseBody() {
	response, err := http.Get("https://ryemiller.io")
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
