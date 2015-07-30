package main

import (
	"encoding/asn1"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/gonuts/cbor"
)

type thinger struct {
	String  string
	Integer int
	Float   float64
	Time    time.Time
	Pairs   map[string]string
	List    []string
}

func main() {
	MarshalStruct()
}

func MarshalStruct() {
	t := thinger{
		"Foo",
		123,
		1234.5678,
		time.Now(),
		map[string]string{"hello": "world"},
		[]string{"one", "two", "three"},
	}

	j, _ := json.Marshal(t)
	c, _ := cbor.Marshal(t)
	a, err := asn1.Marshal(t)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}

	w := tabwriter.NewWriter(os.Stdout, 20, 4, 1, ' ', 0)
	fmt.Fprintln(w, "Encoding\tSize\tGraph")
	tpl := "%s\t%d\t%s\n"

	fmt.Fprintf(w, tpl, "JSON", len(j), strings.Repeat("#", len(j)))
	fmt.Fprintf(w, tpl, "ASN.1", len(a), strings.Repeat("#", len(a)))
	fmt.Fprintf(w, tpl, "CBOR", len(c), strings.Repeat("#", len(c)))
	w.Flush()
}
