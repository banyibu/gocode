package main

import(
//  "fmt"
  "log"
  "net/http"
  "github.com/banyibu/gocode/ch1/lissaj"
)

func main(){
  http.HandleFunc("/", handler)
  log.Fatal(http.ListenAndServe("localhost:8000",nil))
}

func handler(w http.ResponseWriter, r *http.Request){
  //fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
  lissaj.Lissajous(w)
}
