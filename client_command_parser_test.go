package main

import (
"testing"
)

func TestNewClientCommandParser( t *testing.T ) {
	context := NewContext()
	parser, _ := NewClientCommandParser( context )
	if parser == nil {
		t.Error( "Could not create command parser" )
	}
}
