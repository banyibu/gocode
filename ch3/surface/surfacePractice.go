package main

import (
  "fmt"
  "math"
)

const (
  width, height = 600, 320
  cells = 100
  xyrange = 30.0
  xyscale = width / 2 /xyrange
  zscale = height * 0.4
  angle = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main(){
  fmt.Printf("<svg xmlns = 'http://www.w3.org/2000/svg' "+
    "style='stroke: grey; fill: white; stroke-width: 0.7' "+
    "width='%d' height='%d'>", width,height)
  for i := 0; i<cells; i++{
    for j := 0; j<cells; j++{
      ax, ay, b1 := corner(i+1,j)
      bx, by, b2 := corner(i, j)
      cx, cy, b3 := corner(i, j+1)
      dx, dy, b4 := corner(i+1,j+1)
      if b1||b2||b3||b4{
        fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill = '#0000ff'/>\n", ax,ay,bx,by,cx,cy,dx,dy)
      }else{
        fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill = '#ff0000'/>\n", ax,ay,bx,by,cx,cy,dx,dy)
      }
      //fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n", -ax,-ay,-bx,-by,-cx,-cy,-dx,-dy)
    }
  }
  fmt.Println("</svg>")
}

func corner(i,j int)(float64, float64, bool){
  x := xyrange * (float64(i)/cells - 0.5)
  y := xyrange * (float64(j)/cells - 0.5)
  z := f(x,y)

  sx := width/2 + (x-y)*cos30*xyscale
  sy := height/2 + (x +y)*sin30*xyscale - z*zscale
  if z>1{
    return sx,sy,true
  }else{
    return sx,sy,false
  }
  //return sx,sy
}

//func f(x, y float64) (float64){
 // z :=1- x*x/9 -y*y/4
//  r := math.Sqrt(z*2)
//  return r
//}
func f(x, y float64) float64{
  r := x*x/20-y*y/20
  return r
}
