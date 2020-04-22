// unfortunately the book suggests using alexa.com which was taken over by AWS
// after publishing, so here we would have to use a different idea
package main

import (
  "bufio"
  "fmt"
  "net/http"
  "io/ioutil"
  "os"
  "time"
)
// TODO: write to a file
func main() {

  start := time.Now()
  ch := make(chan string)
  for _, url := range os.Args[1:] {
    go fetch(url, ch) // start a goroutine
  }
  outFile, _ := os.Create("./program.out")
  defer outFile.Close()
  w := bufio.NewWriter(outFile)

  for range os.Args[1:] {
    w.WriteString(<-ch) // receive from our opened channel
    w.Flush()
  }
  fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

}

func fetch(url string, ch chan<- string) {

  start := time.Now()
  resp, err := http.Get(url)
  defer resp.Body.Close()
  if err != nil {
    // send this back to the channel we're reading from in main()
    ch <- fmt.Sprint(err)
    return
  }
  if err != nil {
    ch <- fmt.Sprintf("while reading %s: %v", url, err)
    return
  }
  robots, err := ioutil.ReadAll(resp.Body)
  secs := time.Since(start).Seconds()
  ch <- fmt.Sprintf("%.2fs %s %s", secs, robots, url)
}
