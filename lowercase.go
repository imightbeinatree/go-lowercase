package main

import (
  "fmt"
  "os"
  "strings"
)

func main() {
  lowercase := strings.ToLower(os.Args[1])
  fmt.Println("{\"string\":\""+lowercase+"\"}")
}


