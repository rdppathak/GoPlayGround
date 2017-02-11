package main

import "fmt"
import "time"

func ParallelFunction() {
  for i := 0; i < 10; i++ {
    fmt.Printf("Running loop for %dth time\n", i)
    time.Sleep(time.Millisecond * 2)
  }
}


func main() {
    go ParallelFunction()
    time.Sleep(time.Millisecond * 3)
    fmt.Printf("Printing from main function\n")
    time.Sleep(time.Millisecond * 40)
}
