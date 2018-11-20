
package main
// server2 is a mini respond and count server
import(
  "fmt"
  "log"
  "net/http"
  "sync"
)

var mu sync.Mutex
var count int

func main(){
  http.HandleFunc("/",hander)
//  http.HandleFunc("/count", counter)
  log.Fatal(http.ListenAndServe("127.0.0.1:8000",nil))
}

func hander(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
  for k, v := range r.Header{
    fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
  }
  fmt.Fprintf(w, "Host = %q\n", r.Host)
  fmt.Fprintf(w, "Header[%q] = %q\n", r.RemoteAddr)
  if err := r.ParseForm(); err != nil{
    log.Print(err)
  }
  for k, v := range r.Form{
    fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
  }
}

// GET / HTTP/1.1
// Header["Accept"] = ["text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"]
// Header["Accept-Language"] = ["en-US,en;q=0.5"]
// Header["Accept-Encoding"] = ["gzip, deflate"]
// Header["Connection"] = ["keep-alive"]
// Header["Upgrade-Insecure-Requests"] = ["1"]
// Header["User-Agent"] = ["Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:62.0) Gecko/20100101 Firefox/62.0"]
// Host = "localhost:8000"
// Header["127.0.0.1:57990"] = %!q(MISSING)
