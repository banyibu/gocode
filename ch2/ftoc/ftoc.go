package main

import "fmt"

func main(){
  const freezingF, bolingF = 32.0, 212.0
  fmt.Printf("%g degreeF = %g degreeC\n", freezingF, fToC(freezingF))
  fmt.Printf("%g degreeF = %g degreeC\n", bolingF, fToC(bolingF))
}

func fToC(f float64) float64{
  return (f - 32) * 5 / 9
}
