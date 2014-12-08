package main

import (
"testing"
"os/user"
"os"
)


func TestNewClientCommand( t *testing.T ) {
	context := new(Context)
	parser, _ := NewClientCommandParser( context )
	if parser == nil {
		t.Error( "Could not create command parser" )
	}
}

func TestParseFileInformation( t *testing.T ) {
	context := new(Context)
	parser, _ := NewClientCommandParser( context )
	commandPart := "foo@bar.com:/home/baz"
	info, e := parser.ParseFileInformation( commandPart )

	t.Log( "ParseFileInformation creation test" )

	if info == nil {
		t.Error( "ParseFileInformation failed to parse", commandPart, "error", e )
	}

	t.Log( "Parse user name" )
	
	if info.userName != "foo" {
		t.Error( "Expected 'foo' got:", info.userName )
	}

	if info.hostName != "bar.com" {
		t.Error( "Expected bar.com, got", info.hostName )
	}

	if info.fileName != "/home/baz" {
		t.Error("Expected filename /home/baz, got", info.fileName)
	}
	
	commandPart = "bar.com:/home/baz"
	
	t.Log( "Test with a missing user name" )

	info, _ = parser.ParseFileInformation( commandPart )
	expected, _ := user.Current()
	
	if info.userName != expected.Username {
		t.Error( "Expected", expected.Username, " got ", info.userName )
	}

	t.Log("Test with missing host name")
	
	commandPart = "/home/baz"
	info, _ = parser.ParseFileInformation( commandPart )
	expectedHost, _ := os.Hostname()

	t.Log( "Expected host:", expectedHost )

	if info.hostName != expectedHost {
		t.Error( "Expected ", expectedHost, " got ", info.hostName )
	}

	if info.fileName != commandPart {
		t.Error( "Expected", commandPart, " got ", info.fileName )
	}

	commandPart = "foo@bar.com" 
	_, err := parser.ParseFileInformation( commandPart )
	if err == nil {
		t.Error( "Missing expected error" )
	} else {
		if err.Error() != "Invalid file specification" {
			t.Error( "Unexpected error", err.Error() )
		}
	}

	
	
}
