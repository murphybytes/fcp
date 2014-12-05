package main

import (
	"github.com/murphybytes/tufer"
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
}


