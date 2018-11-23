
package main

import(
//  "fmt"
  "log"
  "net/http"
  "github.com/banyibu/gocode/ch3/sf"
)

func main(){
  http.HandleFunc("/", handler)
  log.Fatal(http.ListenAndServe("localhost:8000",nil))
}

func handler(w http.ResponseWriter, r *http.Request){
  //fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
  w.Header().Set("Content-type", "image/svg+xml")
  sf.SfPrint(w)
}

//lsof -i:8000  viewing port occupancy
