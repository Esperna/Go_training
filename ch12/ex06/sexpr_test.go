// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package sexpr

import (
	"encoding/json"
	"fmt"
	"testing"
)

// Test verifies that encoding and decoding a complex data value
// produces an equal result.
//
// The test does not make direct assertions about the encoded output
// because the output depends on map iteration order, which is
// nondeterministic.  The output of the t.Log statements can be
// inspected by running the test with the -v flag:
//
// 	$ go test -v gopl.io/ch12/sexpr
//
type SyntaxError struct {
	*json.SyntaxError
	input []byte
}

func (e SyntaxError) Error() string {
	return fmt.Sprintf("syntax error near: `%s`", string(e.input[e.Offset-1:]))
}

func Test(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Color           bool
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
		//Complex1        complex128 //Not Supported by JSON
		//Complex2        complex64 //Not Supported by JSON
		Price float64
		//Func            func(int) int //Not Supported by JSON
		Anything interface{}
		Nothing  interface{}
	}
	// f := func(x int) int {
	// 	return x * x
	// }
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
		//Complex1: complex(1, 2),
		//Complex2: complex(3, 4),
		Price: 15.55,
		//Func:     f,
		Anything: []string{"a", "b", "c"},
		Nothing:  nil,
	}

	// Encode it
	data, err := Marshal(strangelove)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)

	// Decode it
	var movie Movie
	err = json.Unmarshal(data, &movie)

	if e, ok := err.(*json.SyntaxError); ok {
		t.Fatalf("Unmarshal failed: %v", SyntaxError{e, data})
	}
	//t.Logf("Unmarshal() = %+v\n", movie)
	data2, err := Marshal(movie)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data2)

	// Check equality.
	// if !reflect.DeepEqual(movie, strangelove) {
	// 	t.Fatal("not equal")
	// }

	// Pretty-print it:
	// data, err = MarshalIndent(strangelove)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// t.Logf("MarshalIdent() = %s\n", data)
}
