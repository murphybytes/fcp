package main

import (
	"github.com/murphybytes/tufer"
	"fmt"
)

type Client struct {
	ctx *Context
}

func NewClient( ctx *Context ) ( *Client, error ) {
	return &Client{ ctx }, nil
}

//
// Entry point for client code
//
func ( c *Client ) Execute( ) {
	c.ctx.LogDebug( "Starting client" )
	commandParser, _ := NewClientCommandParser( c.ctx )
	_, dest, err := commandParser.Parse()

	if err != nil {
		c.ctx.LogFatal( err )
	}
	
	destService := c.GetService( dest )
	var clientConnection *tufer.Client
	clientConnection, err = tufer.NewClient(destService )

	if err != nil {
		c.ctx.LogFatal( err )
	}

	_, err = clientConnection.Connect()

	if err != nil {
		c.ctx.LogFatal( err )
	}

	c.ctx.LogDebug( "Connection from to ", destService, " succeeds" )
}

func (c *Client) GetService( fileInfo *FileInformation )( string ) {
	return fmt.Sprintf( "%s:%d", fileInfo.hostName, *c.ctx.arguments.port )
}


