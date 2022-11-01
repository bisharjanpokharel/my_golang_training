package main

import "fmt"

type Astro struct {
    name string
    age  int
    mission string
    isneeded bool
}

type nasaMission struct{

    people []Astro
    number int
    message string

}

func main() {

    ast1 := Astro{"a1",2,"test",false}

    ast2 := Astro{"a2",3,"test",false}

    ast3 := Astro{"a3",7,"test2",true}

    fmt.Println(ast1)
    fmt.Println(ast2)
    fmt.Println(ast3)
 
    p := []Astro{ast1, ast2, ast3}

    // print slice

    fmt.Println(p)

    // only display mission


    fmt.Println(p[2].mission)


    
   s := nasaMission{p, 3, "success"}

   // display "s" without fields
    fmt.Println(s)

    // display "s" with fields
    fmt.Printf("%+v", s)

}

