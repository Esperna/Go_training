// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package sexpr

import (
	"bytes"
	"reflect"
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
func TestSexpr(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Color           bool
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
		Complex1        complex128
		Complex2        complex64
		Price           float64
		Anything        interface{}
		Nothing         interface{}
	}

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
		Complex1: complex(1, 2),
		Complex2: complex(3, 4),
		Price:    15.55,
		Anything: []int{1, 2, 3},
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
	if err := Unmarshal(data, &movie); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Unmarshal() = %+v\n", movie)

	// Check equality.
	if !reflect.DeepEqual(movie, strangelove) {
		t.Fatal("not equal")
	}

	//
	dec := NewDecoder(bytes.NewReader(data))
	var movie2 Movie
	if err := dec.Decode(&movie2); err != nil {
		t.Fatalf("Decode failed: %v", err)
	}
	t.Logf("Decode() = %+v\n", movie2)
}

func TestSexprWithStructFieldTag(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int `sexpr:"released"`
		Color           bool
	}
	var tests = []struct {
		want  string
		given Movie
	}{
		{
			`((Title "Dr. Strangelove") (Subtitle "How I Learned to Stop Worrying and Love the Bomb") (released 1964) (Color nil))`,
			Movie{
				Title:    "Dr. Strangelove",
				Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
				Year:     1964,
				Color:    false,
			},
		},
	}
	for _, test := range tests {
		// Encode
		got, err := Marshal(test.given)
		if err != nil {
			t.Fatalf("Marshal failed: %v", err)
		}
		if string(got) != test.want {
			t.Errorf("\ngot : %s\nwant: %s", got, test.want)
		}
		//Decode
		t.Logf("Marshal() = %s\n", got)
		var movie Movie
		if err := Unmarshal(got, &movie); err != nil {
			t.Fatalf("Unmarshal failed: %v", err)
		}
		t.Logf("Unmarshal() = %+v\n", movie)

		// Check equality.
		if !reflect.DeepEqual(movie, test.given) {
			t.Fatal("not equal")
		}
	}
}

func TestSexprWithStructFieldTag2(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int  `sexpr:"released"`
		Color           bool `sexpr:"color,omitempty"`
	}
	var tests = []struct {
		want  string
		given Movie
	}{
		{
			`((Title "Dr. Strangelove") (Subtitle "How I Learned to Stop Worrying and Love the Bomb") (released 1964))`,
			Movie{
				Title:    "Dr. Strangelove",
				Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
				Year:     1964,
				Color:    false,
			},
		},
	}
	for _, test := range tests {
		// Encode
		got, err := Marshal(test.given)
		if err != nil {
			t.Fatalf("Marshal failed: %v", err)
		}
		if string(got) != test.want {
			t.Errorf("\ngot : %s\nwant: %s", got, test.want)
		}

		//Decode
		t.Logf("Marshal() = %s\n", got)
		var movie Movie
		if err := Unmarshal(got, &movie); err != nil {
			t.Fatalf("Unmarshal failed: %v", err)
		}
		t.Logf("Unmarshal() = %+v\n", movie)

		// Check equality.
		if !reflect.DeepEqual(movie, test.given) {
			t.Fatal("not equal")
		}
	}
}
