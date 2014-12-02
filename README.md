gocmds
======

Template for Go commands.

  * Move everything into `Run` function
  * Call `Run` from `main` (and that's all)
  * Call `Run` from test code

See the `main.go` and `main_test.go` files in this project for a real example, including how flags can work.

```
func main() {
  if err := Run(os.Stdin, os.Stdout, os.Args); err != nil {
    log.Fatalln(err)
  }
}
func Run(in io.Reader, out io.Writer, args []string) error {
  return nil // TODO: work
}
```

Then in test code:

```
func TestRun(t *testing.T) {
  in := strings.NewReader(`Hello`)        // pretend stdin input
  var out bytes.Buffer                    // capture stdout output
  args := []string{"gocmds", "-name=Mat"} // simulate commandline args

  // run the main program
  main.Run(in, &out, args)

  // make assertions on output
}
```