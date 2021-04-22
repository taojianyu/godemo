package main

import(
 "fmt"
 "os"
)

func main() {
    fmt.Println("****************************************")
    fmt.Println("* Hello, World! This is Golang! >> 你好,Golang!")
    fmt.Println("****************************************")

    for _,arg := range os.Args{
        fmt.Printf("arg=%s",arg)
    }
}