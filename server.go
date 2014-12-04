package main

import (
	"time"
	"github.com/murphybytes/tufer"
)

func server( inChan <-chan bool, outChan chan<- bool, ctx *Context ) {

	go listen( ctx )

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


func listen( ctx *Context ) {

	ctx.LogDebug( "Started listener routine" )
	listener, err := tufer.NewListener( ctx.arguments.GetServerService() )

	if err != nil {
		ctx.LogFatal( "Error creating listener socket. Error:", err )
	}

	_, err = listener.Accept()

	ctx.LogDebug( "Accept returned:", err )
	
}
