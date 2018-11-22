package main

import "fmt"

const boilingF = 212.0

func main(){
  var f =  boilingF
  var c = (f-32) * 5 / 9
  fmt.Printf("boling point = %g gegree F or %g degree C\n", f, c)
  // output : boiling point = 212degreeF or 100dereeFC
}
