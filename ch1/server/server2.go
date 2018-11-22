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
  http.HandleFunc("/count", counter)
  log.Fatal(http.ListenAndServe("localhost:8000",nil))
}

func hander(w http.ResponseWriter, r *http.Request){
  mu.Lock()
  count++
  mu.Unlock()
  fmt.Fprintf(w, "URL,Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request){
  mu.Lock()
  fmt.Fprintf(w, "Count %d\n", count)
  mu.Unlock()
}

// netstat -nap|grep 8000  look for the process that estable the port
