package main

import (
  "fmt"
  "os"
  "github.com/gaal/go-options/options"
  )

  func main() {
    s := options.NewOptions(`
      fcp - Fast secure file copy
      --
      x,xcite   Exciting option
      `)
    opt := s.Parse(os.Args[1:])
    fmt.Println( opt )
  }
