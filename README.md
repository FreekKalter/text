
# columnswriter
---

```go
import "github.com/FreekKalter/text/columnswriter"
```

Package columnswriter imlements a write filter that prints evenly distributed
columns fitted to current terminal window size.


## TYPES

```go
type Writer struct {
    // contains filtered or unexported fields
}

```


#### New
```go
func New(output io.Writer, inputSep rune, minWidth, padding int) *Writer
```

Create a writer object with specified values, see Init for explanation of
paramaters.



 
#### Flush
```go
func (w *Writer) Flush()
```
Does the actual printing, always call this after everthing you want to print has
been printed.




```go
Example:
	// Print to stdout, input separated by spaces, no min widht, 2 spaces padding
	w := New(os.Stdout, ' ', 0, 2)
	w.nrTerminalColumns = 80 //for testing purposes, not necessary in real world
	
	fmt.Fprintln(w, `This is a test that has a lot of words and treats every `)
	fmt.Fprintln(w, `word as a column.`)
	fmt.Fprintln(w, `It should print nicely formatted columns similar to ls `)
	fmt.Fprintln(w, `directory lisings on nix systems.`)
	
	w.Flush()
	
	// Output:
	// This     is     a       test       that     has      a   lot      of
	// words    and    treats  every      word     as       a   column.  It
	// should   print  nicely  formatted  columns  similar  to  ls       directory
	// lisings  on     nix     systems.

```

 
#### Init
```go
func (w *Writer) Init(output io.Writer, inputSep rune, minWidth, padding int) *Writer
```
A Writer must be initialized with a call to Init. The first pararmater (output)
specifies the filter output. The inputSep is the character by wich each field is
seperated in the input later on.

	minWidth: Is the minimum widht of a column
	padding:  The number of spaces between columns




 
#### Write
```go
func (w *Writer) Write(buf []byte) (n int, err error)
```
Implements Writer interface, so you can call it with the whole Fprint family





