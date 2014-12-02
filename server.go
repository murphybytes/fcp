package main

import (
  "fmt"
)

func server( inChan <-chan bool, outChan chan<- bool, ctxt *Context ) {

  for run := true; run == true; {

      select {
        case <-inChan :
          run = false
        default:
      }

  }

  fmt.Println("quitting")
  outChan <- true

}
