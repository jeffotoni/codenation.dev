// Go in action
// @jeffotoni
// 2019-05-06

package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
)

func main() {
  resp, err := http.Get("http://localhost:8080/api/estado/extensao")
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()

  for i, j := range resp.Header {
    fmt.Print(i)
    fmt.Print(" : ")
    fmt.Println(j)
  }

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatal("Error", err)
  }

  fmt.Println(string(body))
}
