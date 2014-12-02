package main_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/matryer/gocmds"
)

func TestMain(t *testing.T) {

	in := strings.NewReader(`Hello`)        // pretend stdin input
	var out bytes.Buffer                    // capture stdout output
	args := []string{"gocmds", "-name=Mat"} // simulate commandline args

	// run the main program
	main.Run(in, &out, args)

	// make assertions on output
	if out.String() != "Hello Mat" {
		t.Error("wrong output:", out.String())
	}
}
