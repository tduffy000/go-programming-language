package main

import (
  "fmt"
  "os"
  "strings"
  "time"
)



func EchoFor() {
  var s, sep string
  for _, arg := range os.Args[1:] {
    s += sep + arg
    sep = " "
  }
  fmt.Printf(s + "\n")
}

func EchoBase() {
  var s, sep string
  for i := 1; i < len(os.Args); i++ {
    s += sep + os.Args[i]
    sep = " "
  }
  fmt.Printf(s + "\n")
}

func EchoJoin() {
  fmt.Printf(strings.Join(os.Args[1:], " ") + "\n")
}

func main() {
  forStart := time.Now()
  EchoFor()
  forEnd := time.Now()
  baseStart := time.Now()
  EchoBase()
  baseEnd := time.Now()
  joinStart := time.Now()
  EchoJoin()
  joinEnd := time.Now()
  fmt.Printf("for time: %v\nbase time: %v\njoin time: %v\n", forEnd.Sub(forStart), baseEnd.Sub(baseStart), joinEnd.Sub(joinStart))
}
