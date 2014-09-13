package prettytable_test

import (
	"github.com/tatsushid/go-prettytable"
)

func Example() {
	tbl, err := prettytable.NewTable([]prettytable.Column{
		{Header: "COL1"},
		{Header: "COL2", MinWidth: 6},
		{Header: "COL3", AlignRight: true},
	}...)
	if err != nil {
		panic(err)
	}
	tbl.Separator = " | "
	tbl.AddRow("foo", "bar", "baz")
	tbl.AddRow(1, 2.3, 4)
	tbl.Print()
	// Output:
	// COL1 | COL2   | COL3
	// foo  | bar    |  baz
	// 1    | 2.3    |    4
}

func Example_doublewidthChars() {
	tbl, err := prettytable.NewTable([]prettytable.Column{
		{Header: "名前"},
		{Header: "個数", AlignRight: true},
	}...)
	if err != nil {
		panic(err)
	}
	tbl.Separator = " | "
	tbl.AddRow("りんご", 5)
	tbl.AddRow("みかん", 3)
	tbl.AddRow("柿", 2)
	tbl.Print()
	// Output:
	// 名前   | 個数
	// りんご |    5
	// みかん |    3
	// 柿     |    2
}
