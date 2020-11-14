package main

import("fmt"
"os")

func main(){
    fmt.Println(os.Args[1:])
    fmt.Println("你好，世界")
    os.Exit(0)
}
