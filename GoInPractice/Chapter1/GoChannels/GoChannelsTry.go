package main

import "fmt"
import "time"

func PrintReceivedNumber(c chan int) {
  num := 0
  for num >=0 {
    num  = <-c
    fmt.Printf("Received item from channel %d\n", num)
  }
}

func main() {
  c := make(chan int)
  a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, -1}
  go PrintReceivedNumber(c)

  for _, v := range a {
    c<- v
  }
  time.Sleep(time.Millisecond * 10)
  fmt.Printf ("Exiting main program\n")
}
