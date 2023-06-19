package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
)

func main() {

  //ipfs kubo RPC API command
  url := "http://127.0.0.1:5001/api/v0/files/read?arg=/NewDir/loremipsum.txt&offset=0&count=10" 
  method := "POST"

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, nil)

  if err != nil {
    fmt.Println(err)
    return
  }
  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))
}