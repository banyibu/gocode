package main

import "fmt"

func main(){
  test1()
  fmt.Printf("\n")
  test2()
}

func test1(){
  x := "hello!"
  //两个变量x定义在不同的词法块中，且都是显示块
  for i := 0; i < len(x); i++{
    x := x[i]
    if x != '!'{
      x := x + 'A' -'a'
      fmt.Printf("%c", x)
    }
  }
}

func test2(){
  x := "hello!"
  //三个x，其中两个定义你在显示的词法块中，另一个在for的条件中（隐式块）
  for _, x := range x{
    x := x+'A'-'a'
    fmt.Printf("%c", x)
  }
}
