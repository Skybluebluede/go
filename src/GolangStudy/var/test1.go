/*
    四种变量声明方式
*/

package main

import "fmt"

func main() {
    
    //type1
    var a int
    fmt.Println("a = ", a)

    //type 2
    var b int = 100
    fmt.Println("b = ", b)

    //type 3
    var c = 100
    fmt.Println("c = ", c)
    fmt.Printf("type of c is %T\n", c)

    //type 4 best
    d := 100
    fmt.Println("d = ", d)
    fmt.Printf("type of d is %T\n", d)
}
