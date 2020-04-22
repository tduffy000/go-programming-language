package main

import (
  "fmt"
  "os"
  "strconv"
)

func main() {

  var s, sep string
  for idx, arg := range os.Args[1:] {
    s += sep + "idx: " + strconv.Itoa(idx) + " arg: " + arg
    sep = "\n"
  }
  fmt.Printf(s)
}
