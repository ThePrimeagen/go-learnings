package main

import "fmt"

func add(x int, y int) int {
    return x + y
}

func swap(x, y string) (string, string) {
    return y, x
}

func namedReturned(x, y int) (a, b int) {
    a = x + y;
    b = x - y;
    return
}

func main() {
    fmt.Println("Hello Asher! You are such a fun boy!")
    v := -42 // change me!
    fmt.Printf("v is of type %T\n", v)

}


