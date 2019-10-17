package piscine

package main

import (
    "fmt"
    piscine ".."
)

func PointOne(n *int) {
	*n = 1
}

func main() {
    n := 0
    piscine.PointOne(&n)
    fmt.Println(n)
}
