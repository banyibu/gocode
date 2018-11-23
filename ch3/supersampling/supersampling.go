package main

import(
  //"fmt"
  "os"
  "image"
  "image/color"
  "image/png"
)

func main(){
  f1, err := os.Open(
    "/Users/byb/go/src/github.com/banyibu/gocode/ch3/mandelbort/outcolor.png")
  if err != nil{
    panic(err)
  }
  defer f1.Close()
  m1, err := png.Decode(f1)
  if err != nil{
    panic(err)
  }
  // m := image.NewRGBA(image.Rect(0,0,m1.Bounds().Dx() /2, m1.Bounds().Dy() /2))
  // for i := 0; i<m1.Bounds().Dx()/2; i++{
  //   for j := 0; j<m1.Bounds().Dy()/2; j++{
  //     m.Set(2*i,2*j,m1.At(i,j))
  //     //m.Set(2*i,2*j+1,m1.At(i,j))
  //     //m.Set(2*i+1,2*j,m1.At(i,j))
  //     r1, g1, b1, a1 := m1.At(2*i,2*j).RGBA()
  //     r2, g2, b2, a2 := m1.At(2*i, 2*j+1).RGBA()
  //     r3, g3, b3, a3 := m1.At(2*i+1, 2*j).RGBA()
  //     r4, g4, b4, a4 := m1.At(2*i+1, 2*j+1).RGBA()
  //     m.Set(i,j,color.RGBA{uint8(float64(r1+r2+r3+r4)/4), uint8(float64(g1+g2+g3+g4)/4), uint8(float64(b1+b2+b3+b4)/4), uint8(float64(a1+a2+a3+a4)/4)})
  //     //m.Set(i,j,color.RGBA{uint8(r2), uint8(g2), uint8(b2), uint8(a2)})
  //
  //   }
  // }

  m := image.NewRGBA(image.Rect(0,0,m1.Bounds().Dx(), m1.Bounds().Dy()))
  for i := 0; i<m1.Bounds().Dx(); i++{
    for j := 0; j<m1.Bounds().Dy(); j++{
      m.Set(2*i,2*j,m1.At(i,j))
      //m.Set(2*i,2*j+1,m1.At(i,j))
      //m.Set(2*i+1,2*j,m1.At(i,j))
      r1, g1, b1, a1 := m1.At(i,j).RGBA()
      r2, g2, b2, a2 := m1.At(i, j+1).RGBA()
      r3, g3, b3, a3 := m1.At(i+1, j).RGBA()
      r4, g4, b4, a4 := m1.At(i+1, j+1).RGBA()
      m.Set(i,j,color.RGBA{uint8(float64(r1+r2+r3+r4)/4), uint8(float64(g1+g2+g3+g4)/4), uint8(float64(b1+b2+b3+b4)/4), uint8(float64(a1+a2+a3+a4)/4)})
      //m.Set(i,j,color.RGBA{uint8(r2), uint8(g2), uint8(b2), uint8(a2)})

    }
  }
  png.Encode(os.Stdout, m)
  //fmt.Println("success!")
}
