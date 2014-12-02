package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"
)

// main contains the only three lines that ever need to go
// in a main function in go. ;-)
func main() {
	if err := Run(os.Stdin, os.Stdout, os.Args); err != nil {
		log.Fatalln(err)
	}
}

// Run runs the command and returns nil if successful, otherwise
// an error.
func Run(in io.Reader, out io.Writer, args []string) error {

	// flags
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	var (
		name = flags.String("name", "", "your name")
	)
	flags.Parse(args[1:])

	// read input, add name, send to output
	s := bufio.NewScanner(in)
	for s.Scan() {
		out.Write(s.Bytes())
		io.WriteString(out, " "+*name)
	}
	if err := s.Err(); err != nil {
		return err
	}

	return nil
}
