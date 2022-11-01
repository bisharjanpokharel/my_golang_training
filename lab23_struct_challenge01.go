package main

import "fmt"

type Astro struct {
    name string
    age  int
    mission string
    isneeded bool
}

func main() {

    ast1 := Astro{"a1",2,"test",false}

    ast2 := Astro{"a2",3,"test",false}

    ast3 := Astro{"a3",7,"test2",true}

    fmt.Println(ast1)
    fmt.Println(ast2)
    fmt.Println(ast3)
}

