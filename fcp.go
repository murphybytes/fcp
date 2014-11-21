package main

import (
  "fmt"
  "flag"
  )

  type Arguments struct {
    deamonize *bool
    server    *bool

  }

  func main() {

    arguments := new(Arguments)
    arguments.deamonize = flag.Bool("d", false, "Run server as a daemon")
    arguments.server    = flag.Bool("s", false, "Run server in foreground")
    
    flag.Parse()

    fmt.Println( *arguments.deamonize )
  }
