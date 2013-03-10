# columnswriter
--
    import "github.com/FreekKalter/text/columnswriter"

Package columnswriter imlements a write filter that prints evenly distributed
columns fitted to current terminal window size.

## Usage

#### type Writer

```go
type Writer struct {
}
```

A Writer is a filter that prints evenly distributed columns

#### func  New

```go
func New(output io.Writer, inputSep rune, minWidth, padding int) *Writer
```
Create a writer object with specified values, see Init for explanation of
paramaters.

#### func (*Writer) Flush

```go
func (w *Writer) Flush()
```

#### func (*Writer) Init

```go
func (w *Writer) Init(output io.Writer, inputSep rune, minWidth, padding int) *Writer
```
A Writer must be initialized with a call to Init. The first pararmater (output)
specifies the filter output. The inputSep is the character by wich each field is
seperated in the input later on.

    minWidth: Is the minimum widht of a column
    padding:  The number of spaces between columns

#### func (*Writer) Write

```go
func (w *Writer) Write(buf []byte) (n int, err error)
```
