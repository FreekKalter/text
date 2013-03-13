package columnswriter

import (
	"fmt"
	"os"
)

// Example of how to call flush after you have written everything
func ExampleWriter_Flush() {
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
}
