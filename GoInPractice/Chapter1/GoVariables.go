import "fmt"


func main() {
  //Using following syntax one can declare int or any type of data variable

  var i int   //Indicates that "i" is a variable of type int

  var j = 2 //Compiler at compile time will identify that variable "j" will of type int

  k := 2  //Shortest form of declaration.

  fmt.Printf("Printing variables [i = %d], [j = %d], [k = %d]", i, j, k)
}
