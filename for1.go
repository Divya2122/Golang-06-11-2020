package main

import "fmt"

func main() {

    i := 2
    for i <= 3 {
        fmt.Println(i)
        i = i + 3
    }

    for j := 8; j <= 7; j++ {
        fmt.Println(j)
    }

    for {
        fmt.Println("loop")
        break
    }

    for n := 1; n <= 6; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}