package main

import "fmt"

func GetNames() (string , string){
  return "FirstString", "SecondString"
}

func main() {
  var1, var2 := GetNames()
  fmt.Printf("Returned names: Var1 = %s, Var2 = %s", var1, var2)
}
