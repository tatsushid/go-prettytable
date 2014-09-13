package prettytable

import (
	"bytes"
	"strings"
	"testing"
)

func TestTable(t *testing.T) {
	buf := bytes.NewBuffer(make([]byte, 0))
	tbl, err := NewTable(
		Column{Header: "COL1", AlignRight: false},
		Column{Header: "COL2", AlignRight: true},
		Column{Header: "COL3", AlignRight: false},
	)
	if err != nil {
		t.Fatalf("Unable to create table: %s", err)
	}

	tbl.AddRow("foo", "bar", "baz")
	tbl.AddRow("test", "sample")
	tbl.AddRow("あ", "い", "う")

	_, err = tbl.WriteTo(buf)
	if err != nil {
		t.Fatalf("Unable to write to buffer: %s", err)
	}

	expect := `COL1   COL2 COL3
foo     bar baz
test sample
あ       い う
`
	if buf.String() != expect {
		t.Errorf("WriteTo wrote unexpected string.\ngot: %d\n%s\nexpect: %d\n%s",
			len(buf.String()), strings.Replace(buf.String(), " ", "_", -1),
			len(expect), strings.Replace(expect, " ", "_", -1))
	}
}

func TestTableWithVariousArgs(t *testing.T) {
	buf := bytes.NewBuffer(make([]byte, 0))
	tbl, err := NewTable(
		Column{Header: "COL1", AlignRight: false},
		Column{Header: "COL2", AlignRight: true},
		Column{Header: "COL3", AlignRight: false},
	)
	if err != nil {
		t.Fatalf("Unable to create table: %s", err)
	}

	tbl.AddRow(100, "bar", 5.2)
	tbl.AddRow([]byte("test"), "sample")
	tbl.AddRow("あ", true, "う")

	_, err = tbl.WriteTo(buf)
	if err != nil {
		t.Fatalf("Unable to write to buffer: %s", err)
	}

	expect := `COL1   COL2 COL3
100     bar 5.2
test sample
あ     true う
`
	if buf.String() != expect {
		t.Errorf("WriteTo wrote unexpected string.\ngot: %d\n%s\nexpect: %d\n%s",
			len(buf.String()), strings.Replace(buf.String(), " ", "_", -1),
			len(expect), strings.Replace(expect, " ", "_", -1))
	}
}

func TestMinWidth(t *testing.T) {
	buf := bytes.NewBuffer(make([]byte, 0))
	tbl, err := NewTable(
		Column{Header: "COL1", AlignRight: false, MinWidth: 5},
		Column{Header: "COL2", AlignRight: true, MinWidth: 4},
		Column{Header: "COL3", AlignRight: false, MinWidth: 6},
	)
	if err != nil {
		t.Fatalf("Unable to create table: %s", err)
	}

	tbl.AddRow("foo", "bar", "baz")
	tbl.AddRow("test", "sample")
	tbl.AddRow("あ", "い", "う")

	_, err = tbl.WriteTo(buf)
	if err != nil {
		t.Fatalf("Unable to write to buffer: %s", err)
	}

	expect := `COL1    COL2 COL3
foo      bar baz
test  sample
あ        い う
`
	if buf.String() != expect {
		t.Errorf("WriteTo wrote unexpected string.\ngot: %d\n%s\nexpect: %d\n%s",
			len(buf.String()), strings.Replace(buf.String(), " ", "_", -1),
			len(expect), strings.Replace(expect, " ", "_", -1))
	}
}

func TestMaxWidth(t *testing.T) {
	buf := bytes.NewBuffer(make([]byte, 0))
	tbl, err := NewTable(
		Column{Header: "COL1", AlignRight: false, MaxWidth: 3},
		Column{Header: "COL2", AlignRight: true, MaxWidth: 4},
		Column{Header: "COL3", AlignRight: false, MaxWidth: 6},
	)
	if err != nil {
		t.Fatalf("Unable to create table: %s", err)
	}

	tbl.AddRow("foo", "bar", "baz")
	tbl.AddRow("test", "sample")
	tbl.AddRow("あい", "う", "え")

	_, err = tbl.WriteTo(buf)
	if err != nil {
		t.Fatalf("Unable to write to buffer: %s", err)
	}

	expect := `COL COL2 COL3
foo  bar baz
tes samp
あ    う え
`
	if buf.String() != expect {
		t.Errorf("WriteTo wrote unexpected string.\ngot: %d\n%s\nexpect: %d\n%s",
			len(buf.String()), strings.Replace(buf.String(), " ", "_", -1),
			len(expect), strings.Replace(expect, " ", "_", -1))
	}
}
