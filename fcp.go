package main

import (
  "fmt"
  "flag"
  "os"
  "os/signal"
  )

  type Arguments struct {
    daemonize *bool
    server    *bool
    verbose   *bool
    noEncrypt *bool

  }

  func main() {

    arguments := new(Arguments)
    arguments.daemonize = flag.Bool("d", false, "Run server as a daemon")
    arguments.server    = flag.Bool("s", false, "Run server in foreground")
    arguments.verbose   = flag.Bool("v", false, "Verbose log messages")
    arguments.noEncrypt = flag.Bool("n", false, "Disable encryption")

   flag.Usage = func() {
     fmt.Println("Usage: fcp [options] [[[user1][@host1]]:sourcefile [[user2][@host2]]:destfile]")
     fmt.Println("Options:")
     flag.PrintDefaults()
   }
    flag.Parse()

    if *arguments.daemonize || *arguments.server {
      signalChannel := make(chan os.Signal, 1)
      signal.Notify( signalChannel, os.Kill, os.Interrupt )

      fmt.Println("Starting in server mode.")
      // fire up in server mode

      sig := <-signalChannel
      fmt.Println("Got signal: ", sig )

    } else {
      // if we are here we are in interactive mode, look for file source and
      // destination as last two arguments
      fileArgs := flag.Args()
      if len(fileArgs) != 2 {
        flag.Usage()

        os.Exit( ERROR_IMPROPER_USAGE )
      }
    }


  }
