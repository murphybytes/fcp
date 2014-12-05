package main

type ClientCommandParser struct {
	ctx *Context
}


func NewClientCommandParser( ctx *Context )( *ClientCommandParser, error ) {
	return &ClientCommandParser{ ctx }, nil
}

func (p *ClientCommandParser) Parse() ( source *ClientCommand, dest *ClientCommand, error ) {
}
