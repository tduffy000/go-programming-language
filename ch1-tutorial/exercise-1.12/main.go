package main

import (
  "log"
  "net/http"
  "fmt"
  "strconv"
)

/** TODO: add params for
  * ?cycles={int}
  * ?size={int}
  * ?nframes={int}
  */

func main() {
  // here we can use a function literal (anonymous function) to call our lissajous
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("%s %s %s\n", r.Method, r.URL, r.Proto)
    r.ParseForm()
    fmt.Printf("Form %v\n", r.Form)
    for k, v := range r.Form {
      fmt.Printf("Form[%q] = %q\n", k , v)
    }
    if r.URL.String() != "/favicon.ico" { // this is a hack for Ubuntu
      r.ParseForm()
      inputCycles, _ := strconv.Atoi(r.Form["cycles"][0])
      lissajous(w, inputCycles)
    }
  })
  log.Fatal(http.ListenAndServe("localhost:8001", nil))
}
