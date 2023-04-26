package main

import (
    "fmt"
    "os"
)

func main() {
    name := os.Getenv("INPUT_NAME")
    if name == "" {
        name = "World"
    }

    greeting := fmt.Sprintf("Hello, %s!", name)
    fmt.Println(greeting)
}


