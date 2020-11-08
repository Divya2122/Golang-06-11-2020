package main 
import (
        "fmt"
        "strconv"
)
func main() {
         var i int = 42
         fmt.printf("%v, %t\n", i,i)
         var j string
         j = strconv.Itoa(i)
         fmt.printf("%v, %T\n",j, j)
} 
