package main
import (
  "fmt"
  "net/http"
  "os"
  "strings"
  "io"
)
// input: command args[1:] input urls;
//catch the json content of the url and some other ifomation, such as status and errcode.
func main(){
  for _, url :=range os.Args[1:]{
    if !strings.HasPrefix(url, "http"){
      b := "http://"
      url = b + url
    }
    fmt.Println(url)
    resp, err := http.Get(url)
    if err != nil{
      fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
      os.Exit(1)
    }
    _ ,err = io.Copy(os.Stdout, resp.Body)
    resp.Body.Close()
    if err != nil{
      fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url,err)
      os.Exit(1)
    }
    fmt.Printf("%v",resp.Status)
  }
}
