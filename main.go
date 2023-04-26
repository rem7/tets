package main

import (
    "fmt"
    "os"
)

func main() {
    name := os.Args[1]
    if name == "" {
        name = "World"
    }

    greeting := fmt.Sprintf("Hello, %s!", name)
    fmt.Println(greeting)
}


