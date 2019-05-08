// Go in action
// @jeffotoni
// 2019-05-06

package main

import "log"
import "net/http"
import "time"
import "io/ioutil"
import "fmt"

func main() {

	req, err := http.NewRequest("GET", "http://localhost:8080/api/estado/extensao", nil)
	if err != nil {
		log.Fatal("error read request", err)
	}

	req.Header.Set("Cache-Control", "no-cache")
	client := &http.Client{Timeout: time.Second * 5}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error read response ", err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error read Body", err)
		return
	}

	fmt.Printf("%s\n", body)
}
