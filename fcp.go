package main

import (
  "fmt"
  "flag"
  "os"
  )

  type Arguments struct {
    daemonize *bool
    server    *bool

  }

  func main() {

    arguments := new(Arguments)
    arguments.daemonize = flag.Bool("d", false, "Run server as a daemon")
    arguments.server    = flag.Bool("s", false, "Run server in foreground")


    flag.Parse()

    if *arguments.daemonize || *arguments.server {
      // fire up in server mode
    } else {
      // if we are here we are in interactive mode, look for file source and
      // destination as last two arguments
      fileArgs := flag.Args()
      if len(fileArgs) != 2 {
        os.Exit( errorImproperUsage )
      }
    }

    fmt.Println( *arguments.daemonize )
  }
