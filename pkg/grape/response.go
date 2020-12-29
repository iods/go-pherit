// grape is a web scraper in go
package grape

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func ResponseSize(url string) {
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
	fmt.Println(len(body))
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
