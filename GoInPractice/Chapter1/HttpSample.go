package main

import "fmt"
import "io/ioutil"
import "net/http"

func main() {
    resp,  _ := http.Get("http://google.co.in/")
    if resp == nil{
      fmt.Printf("Not able to make connection with server\n")
      return
    }
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Printf("Reading data length: %d", len(string(body)))
    resp.Body.Close()
}
