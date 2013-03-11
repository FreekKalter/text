package columnswriter

import (
	"fmt"
	"os"
)

func ExampleWriter_Flush() {
	toPrint := `This is a test that has a lot of words and treats every word as a column. It should print nicely formatted columns similar to ls directory lisings on nix systems.`
	w := New(os.Stdout, ' ', 0, 2) // Print to stdout, input is separated by spaces, no minimal widht, 2 spaces of padding
	w.nrTerminalColumns = 119
	fmt.Fprint(w, toPrint)
	w.Flush()

	//Output:
	//This      is       a   test    that   has     a          lot      of       words  and  treats     every    word  as   
	//a         column.  It  should  print  nicely  formatted  columns  similar  to     ls   directory  lisings  on    nix  
	//systems.  
}
