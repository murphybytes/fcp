package main

import (
	"flag"
	"os"
	"os/signal"
)



func main() {
	
	ctx := NewContext()
	flag.Parse()

	if  *ctx.arguments.server {
		// create pid if user wants it
		ctx.CreatePidFile()
		signalChannel := make(chan os.Signal, 1)
		signal.Notify( signalChannel, os.Kill, os.Interrupt )

		ctx.LogInfo("Starting in server mode.")
		// fire up in server mode
		inChan    := make(chan bool, 1 )
		replyChan := make(chan bool, 1 )

		go server( inChan, replyChan, ctx )

		sig := <-signalChannel
		ctx.LogDebug("Got signal: ", sig )
		// signal server goroutine to shut down
		inChan<- true
		// block until it stops
		<-replyChan
		// get rid of pid file if one was created
		ctx.RemovePidFile()
		ctx.LogInfo("Server stopped, exiting...")


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
