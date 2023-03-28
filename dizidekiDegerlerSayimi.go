package main

import (
    "strconv"
	"fmt"
)

func areAllNumbers(arr []string) bool {
    for key, val := range arr {
        _, err := strconv.Atoi(val)
        if err != nil {
            return false
        }
    }
    return true
}


func main(){


arr := []string{"1", "3", "2", "5"}
if areAllNumbers(arr) {
    fmt.Println("All elements are numbers")
} else {
    fmt.Println("Not all elements are numbers")
}
}