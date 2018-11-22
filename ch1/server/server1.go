package main

import(
  "fmt"
  "log"
  "net/http"
)

func main(){
  http.HandleFunc("/", handler)
  log.Fatal(http.ListenAndServe("localhost:8000",nil))
}

func handler(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

/*  *ps

    *ps -aux 查看所有进程，每行一个程序

    *top 显示当前运行程序

    *kill 98 （98为PID号，）

    *kill -9 98 （强制杀死98）  */
