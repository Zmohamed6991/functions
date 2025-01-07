package main 

import (
    "fmt"

)

func callingFunc() {
    fmt.Println("Calling function")
}

func calculatingFunc(a int, b int) int {
    return a + b
}


func main() {

    callingFunc()
    fmt.Println(calculatingFunc(5,4))


}

