package main

import (
  "fmt"
  )

func server( inchan <-chan bool, outchan chan<- bool ) {

  for run := true; run == true; {
      fmt.Println( "running")

      select {
        case <-inchan :
          run = false
        default:
      }

  }

  fmt.Println("quitting")
  outchan <- true

}
