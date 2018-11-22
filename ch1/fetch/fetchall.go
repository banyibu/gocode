package main
// fetchall catch several urls and report the time and size of the urls
import(
  "fmt"
  "io"
  "io/ioutil"
  "net/http"
  "os"
  "time"
)
// input and output example
/**go run fetchall.go www.baidu.com www.hao123.com    *
  *Get www.hao123.com: unsupported protocol scheme "" *
  *Get www.baidu.com: unsupported protocol scheme ""  *
  *%!f(func() float64=0x5eca10)s elapsed              */


func main(){
  start := time.Now()
  ch := make(chan string)
  for _, url := range os.Args[1:]{
    go fetch(url,ch)
  }
  for range os.Args[1:]{
    fmt.Println(<-ch)
  }
  fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds)
}

func fetch(url string, ch chan<-string){
  start := time.Now()
  resp, err := http.Get(url)
  if err != nil{
    ch <- fmt.Sprint(err)
    return
  }
  //_, _ = io.Copy(os.Stdout, resp.Body)
  nbytes, err := io.Copy(ioutil.Discard, resp.Body)
  resp.Body.Close()
  if err != nil{
    ch <- fmt.Sprintf("while reading %s: %v", url, err)
    return
  }
  secs := time.Since(start).Seconds()
  ch <- fmt.Sprintf("%.2fs %7d %s", secs,nbytes, url)
}
