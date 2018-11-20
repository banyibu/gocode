package main

import (
  "image"
  "image/color"
  "image/gif"
  "io"
  "math"
  "math/rand"
  "os"
  "time"
  "net/http"
  "log"
  "strconv"
//  "strings"
)

var palette = []color.Color{color.White, color.Black}
const (
  whiteIndex = 0
  blackIndex = 1
)

func main(){
  rand.Seed(time.Now().UTC().UnixNano())
  if len(os.Args)>1 && os.Args[1]=="web"{
    handler := func(w http.ResponseWriter, r *http.Request){
      if q := r.URL.Query();  q.Get("cycles") != ""{
        val, err := strconv.Atoi(q.Get("cycles"))
        if err !=nil{
          lissajous(w,5)
          }else{
            lissajous(w,val)
            }
      }else{
      lissajous(w,5)
      }
    }
    http.HandleFunc("/",handler)
    log.Fatal(http.ListenAndServe("localhost:8000",nil))
    return
  }
  lissajous(os.Stdout, 5)
}

func lissajous(out io.Writer, cy int){
  var cycles float64 = float64 (cy)
  const (
    res = 0.001
    size = 100
    nframes = 64
    delay = 8
  )
  freq := rand.Float64()*3.0
  anim := gif.GIF{LoopCount: nframes}
  phase := 0.0
  for i := 0; i<nframes; i++{
    rect := image.Rect(0,0,2*size+1,2*size+1)
    img := image.NewPaletted(rect,palette)
    for t := 0.0; t < cycles*2*math.Pi; t+=res{
      x := math.Sin(t)
      y := math.Sin(t*freq + phase)
      img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
    }
    phase += 0.1
    anim.Delay = append(anim.Delay, delay)
    anim.Image = append(anim.Image, img)
  }
  gif.EncodeAll(out,&anim)
}
