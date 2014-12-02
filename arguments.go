package main

import (
	"flag"
	"fmt"
)

type Arguments struct {
	server    *bool
	verbose   *bool
	noEncrypt *bool

}

func NewArguments()( *Arguments ) {

	flag.Usage = func() {
		fmt.Println("Usage: fcp [options] [[[user1][@host1]]:sourcefile [[user2][@host2]]:destfile]")
		fmt.Println("Options:")
		flag.PrintDefaults()
	}


	return &Arguments{
		flag.Bool("s", false, "Run in server mode"),
		flag.Bool("v", false, "Verbose log messages"),
		flag.Bool("n", false, "Disable encryption"),
	}
}
