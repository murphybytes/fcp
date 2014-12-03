package main

import (
	"time"
)

func server( inChan <-chan bool, outChan chan<- bool, ctx *Context ) {

	for run := true; run == true; {

		select {
		case <-inChan :
			run = false
		default:
			var duration time.Duration			
			duration = 100 * time.Millisecond
			time.Sleep(duration)
		}

	}

	ctx.LogDebug( "Exiting server loop")
	outChan <- true

}
